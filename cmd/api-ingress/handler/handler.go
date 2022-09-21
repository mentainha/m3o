package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"m3o.dev/api/client"
)

var (
	// API Key
	APIKey = os.Getenv("M3O_API_TOKEN")

	// API Url
	APIHost = "https://api.m3o.com"

	// dot com host
	ComHost = "m3o.com"

	// host to proxy for URLs
	URLHost = "m3o.one"

	// host to proxy for Apps
	AppHost = "m3o.app"

	// host to proxy for Functions
	FunctionHost = "m3o.sh"

	// host for user auth
	UserHost = "user.m3o.com"
)

type App struct {
	Name string
	Url  string
}

type Handler struct {
	mtx    sync.RWMutex
	appMap map[string]*backend

	ftx     sync.RWMutex
	funcMap map[string]*backend

	client      *client.Client
	lastUpdated time.Time
	appList     []*App
}

type backend struct {
	url     *url.URL
	created time.Time
}

func (h *Handler) listApps() ([]*App, error) {
	var rsp map[string]interface{}
	req := map[string]interface{}{}

	if err := h.client.Call("app", "list", req, &rsp); err != nil {
		return nil, err
	}

	appList := []*App{}
	apps, _ := rsp["services"].([]interface{})

	for _, v := range apps {
		app := v.(map[string]interface{})
		name := app["name"].(string)
		url := app["url"].(string)
		appList = append(appList, &App{
			Name: name,
			Url:  url,
		})
	}

	return appList, nil
}

func (h *Handler) functionProxy(w http.ResponseWriter, r *http.Request) {
	// no subdomain
	if r.Host == FunctionHost {
		return
	}

	// lookup the app map
	h.ftx.RLock()
	bk, ok := h.funcMap[r.Host]
	h.ftx.RUnlock()

	// check the url map
	if ok && time.Since(bk.created) < time.Minute {
		r.Host = bk.url.Host
		r.Header.Set("Host", r.Host)
		httputil.NewSingleHostReverseProxy(bk.url).ServeHTTP(w, r)
		return
	}

	subdomain := strings.TrimSuffix(r.Host, "."+FunctionHost)

	// only process one part for now
	parts := strings.Split(subdomain, ".")
	if len(parts) > 1 {
		log.Print("[function/proxy] more parts than expected", parts)
		return
	}

	// currently service id is the subdomain
	id := subdomain

	log.Printf("[function/proxy] resolving host %s to id %s\n", r.Host, id)

	apiURL := APIHost + "/function/proxy"

	// use /v1/
	if len(APIKey) > 0 {
		apiURL = APIHost + "/v1/function/proxy"
	}

	// make new request
	log.Printf("[function/proxy] Calling: %v", apiURL+"?id="+id)
	req, err := http.NewRequest("GET", apiURL+"?id="+id, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	// set the api key after we're given the header
	if len(APIKey) > 0 {
		req.Header.Set("Authorization", "Bearer "+APIKey)
	}

	// call the backend for the url
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if rsp.StatusCode != 200 {
		log.Printf("[function/proxy] Error calling api: status: %v %v", rsp.StatusCode, string(b))
		http.Error(w, "unexpected error", 500)
		return
	}

	result := map[string]interface{}{}

	if err := json.Unmarshal(b, &result); err != nil {
		log.Print("[function/proxy] failed to unmarshal response")
		http.Error(w, err.Error(), 500)
		return
	}

	// get the destination url
	u, _ := result["url"].(string)
	if len(u) == 0 {
		log.Print("[function/proxy] no response url")
		return
	}

	uri, err := url.Parse(u)
	if err != nil {
		log.Print("[function/proxy] failed to parse url", err.Error())
		return
	}

	h.ftx.Lock()
	h.funcMap[r.Host] = &backend{
		url:     uri,
		created: time.Now(),
	}
	h.ftx.Unlock()

	r.Host = uri.Host
	r.Header.Set("Host", r.Host)

	httputil.NewSingleHostReverseProxy(uri).ServeHTTP(w, r)
}

func (h *Handler) appProxy(w http.ResponseWriter, r *http.Request) {
	// load the landing page
	if r.Host == AppHost {
		h.mtx.RLock()
		lastUpdated := h.lastUpdated
		appList := h.appList
		h.mtx.RUnlock()

		if lastUpdated.IsZero() || time.Since(lastUpdated) > time.Minute {
			newList, err := h.listApps()
			if err == nil {
				h.mtx.Lock()
				h.appList = newList
				h.lastUpdated = time.Now()
				appList = newList
				h.mtx.Unlock()
			}
		}

		// if home app exists load it
		var home *App
		for _, app := range appList {
			if app.Name == "home" {
				home = app
				break
			}
		}

		// we have a home app that will load instead of the index
		if home != nil {
			u, err := url.Parse(home.Url)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			rx := httputil.NewSingleHostReverseProxy(u)
			r.URL.Host = u.Host
			r.URL.Scheme = u.Scheme
			r.Header.Set("X-Forwarded-Host", r.Header.Get("host"))
			r.Host = u.Host
			rx.ServeHTTP(w, r)
			return

		}

		// load the default index

		// parse the web template
		t, err := template.New("index").Parse(WebTemplate)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// load the index template with app data
		if err := t.ExecuteTemplate(w, "index", appList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	// lookup the app map
	h.mtx.RLock()
	bk, ok := h.appMap[r.Host]
	h.mtx.RUnlock()

	// check the url map
	if ok && time.Since(bk.created) < time.Minute {
		r.Host = bk.url.Host
		r.Header.Set("Host", r.Host)
		httputil.NewSingleHostReverseProxy(bk.url).ServeHTTP(w, r)
		return
	}

	subdomain := strings.TrimSuffix(r.Host, "."+AppHost)
	subdomain = strings.TrimSuffix(subdomain, "."+ComHost)

	// only process one part for now
	parts := strings.Split(subdomain, ".")
	if len(parts) > 1 {
		log.Print("[app/proxy] more parts than expected", parts)
		return
	}

	// currently service id is the subdomain
	id := subdomain

	log.Printf("[app/proxy] resolving host %s to id %s\n", r.Host, id)

	apiURL := APIHost + "/app/resolve"

	// use /v1/
	if len(APIKey) > 0 {
		apiURL = APIHost + "/v1/app/resolve"
	}

	// make new request
	log.Printf("[app/proxy] Calling: %v", apiURL+"?id="+id)
	req, err := http.NewRequest("GET", apiURL+"?id="+id, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	// set the api key after we're given the header
	if len(APIKey) > 0 {
		req.Header.Set("Authorization", "Bearer "+APIKey)
	}

	// call the backend for the url
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if rsp.StatusCode != 200 {
		log.Printf("[app/proxy] Error calling api: status: %v %v", rsp.StatusCode, string(b))
		http.Error(w, "unexpected error", 500)
		return
	}

	result := map[string]interface{}{}

	if err := json.Unmarshal(b, &result); err != nil {
		log.Print("[app/proxy] failed to unmarshal response")
		http.Error(w, err.Error(), 500)
		return
	}

	// get the destination url
	u, _ := result["url"].(string)
	if len(u) == 0 {
		log.Print("[app/proxy] no response url")
		return
	}

	uri, err := url.Parse(u)
	if err != nil {
		log.Print("[app/proxy] failed to parse url", err.Error())
		return
	}

	h.mtx.Lock()
	h.appMap[r.Host] = &backend{
		url:     uri,
		created: time.Now(),
	}
	h.mtx.Unlock()

	r.Host = uri.Host
	r.Header.Set("Host", r.Host)

	httputil.NewSingleHostReverseProxy(uri).ServeHTTP(w, r)
}

func (h *Handler) urlProxy(w http.ResponseWriter, r *http.Request) {
	uri := url.URL{
		Scheme: r.URL.Scheme,
		Host:   r.URL.Host,
		Path:   r.URL.Path,
	}

	// if the host is blank we have to set it
	if len(uri.Host) == 0 {
		if len(r.Host) > 0 {
			log.Printf("[url/proxy] Host is set from r.Host %v", r.Host)
			uri.Host = r.Host
		} else {
			log.Printf("[url/proxy] Host is nil, defaulting to: %v", URLHost)
			uri.Host = URLHost
			uri.Scheme = "https"
		}
	}

	if len(uri.Scheme) == 0 {
		uri.Scheme = "https"
	}

	apiURL := APIHost + "/url/resolve"

	// use /v1/
	if len(APIKey) > 0 {
		apiURL = APIHost + "/v1/url/resolve"
	}

	var id, params string

	// check url prefix
	if strings.HasPrefix(r.URL.Path, "/url/") {
		id = strings.TrimPrefix(r.URL.Path, "/url/")
	}

	if len(id) > 0 {
		params = "?id=" + id
	} else {
		params = "?shortURL=" + uri.String()
	}

	// make new request
	log.Printf("[url/proxy] Calling: %v", apiURL+params)
	req, err := http.NewRequest("GET", apiURL+params, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	// set the api key after we're given the header
	if len(APIKey) > 0 {
		req.Header.Set("Authorization", "Bearer "+APIKey)
	}

	// call the backend for the url
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if rsp.StatusCode != 200 {
		log.Printf("[url/proxy] Error calling api: status: %v %v", rsp.StatusCode, string(b))
		http.Error(w, "unexpected error", 500)
		return
	}

	result := map[string]interface{}{}

	if err := json.Unmarshal(b, &result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get the destination url
	url, _ := result["destinationURL"].(string)
	if len(url) == 0 {
		return
	}

	// return the redirect url to caller
	http.Redirect(w, r, url, 302)
}

func (h *Handler) userProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// can't operate without the api key
	if len(APIKey) == 0 {
		return
	}

	token := r.Form.Get("token")
	if len(token) == 0 {
		log.Print("Missing token")
		return
	}

	redirectUrl := r.Form.Get("redirectUrl")
	if len(redirectUrl) == 0 {
		log.Print("Missing redirect url")
		return
	}

	// check access and redirect
	uri := APIHost + "/v1/user/VerifyEmail"

	b, _ := json.Marshal(map[string]interface{}{
		"token": token,
	})

	req, err := http.NewRequest("POST", uri, bytes.NewReader(b))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	req.Header.Set("Authorization", "Bearer "+APIKey)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rsp.Body.Close()

	// discard the body
	io.Copy(ioutil.Discard, rsp.Body)

	if rsp.StatusCode == 200 {
		http.Redirect(w, r, redirectUrl, 302)
		return
	}

	// non 200 status code

	// redirect to failure url
	failureUrl := r.Form.Get("failureRedirectUrl")
	if len(failureUrl) == 0 {
		return
	}

	// redirect
	http.Redirect(w, r, failureUrl, 302)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// m3o.app
	if strings.HasSuffix(r.Host, AppHost) {
		h.appProxy(w, r)
		return
	}

	// m3o.sh
	if strings.HasSuffix(r.Host, FunctionHost) {
		h.functionProxy(w, r)
		return
	}

	// user.m3o.com
	if r.Host == UserHost {
		h.userProxy(w, r)
		return
	}

	// if it's /url then resolve by id
	if strings.HasPrefix(r.URL.Path, "/url/") {
		h.urlProxy(w, r)
	}

	// if it's /user then use user proxy
	if strings.HasPrefix(r.URL.Path, "/user/") {
		h.userProxy(w, r)
	}

	// if it's /url then resolve by id
	if strings.HasPrefix(r.URL.Path, "/u/") {
		h.urlProxy(w, r)
	}

	// 404
	http.NotFound(w, r)
}

func New() *Handler {
	h := &Handler{
		client:  client.NewClient(&client.Options{Token: APIKey}),
		appMap:  make(map[string]*backend),
		funcMap: make(map[string]*backend),
	}
	return h
}
