package handler

import (
	"bytes"
	"encoding/json"
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

type Handler struct {
	mtx    sync.RWMutex
	appMap map[string]*backend

	ftx     sync.RWMutex
	funcMap map[string]*backend
}

type backend struct {
	url     *url.URL
	created time.Time
}

// v1Proxy manages requests to the /v1 api by converting an api key param to authorization header
func (h *Handler) v1Proxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("api_key")
	// strip the api key
	r.Form.Del("api_key")

	// try authorization header
	if len(key) == 0 {
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, "Bearer ") {
			key = strings.TrimPrefix(auth, "Bearer ")
		}
	}

	if len(key) == 0 {
		http.Error(w, "Missing api key", 401)
		return
	}

	c := client.NewClient(nil)
	c.SetToken(key)

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		http.Error(w, "Bad url format", 400)
		return
	}

	service := parts[2]
	endpoint := parts[3]

	var req interface{}

	if ct := r.Header.Get("Content-Type"); ct == "application/json" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if len(b) == 0 {
			b = []byte(`{}`)
		}
		req = json.RawMessage(b)
	} else {
		vals := map[string]string{}
		for k := range r.Form {
			vals[k] = r.Form.Get(k)
		}
		req = vals
	}

	var rsp json.RawMessage

	err := c.Call(service, endpoint, req, &rsp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// write the resposne
	w.Write(rsp)
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
	// no subdomain
	if r.Host == AppHost {
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
	var id string

	// if it's /api then use the v1 proxy
	if strings.HasPrefix(r.URL.Path, "/api/") {
		h.v1Proxy(w, r)
		return
	}

	// check url prefix
	if strings.HasPrefix(r.URL.Path, "/url/") {
		id = strings.TrimPrefix(r.URL.Path, "/url/")
	}

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

	var params string

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
	// m3o.one
	if strings.HasSuffix(r.Host, URLHost) {
		h.urlProxy(w, r)
		return
	}

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

	if r.Host == UserHost {
		h.userProxy(w, r)
		return
	}

	// if it's /api then use the v1 proxy
	if strings.HasPrefix(r.URL.Path, "/api/") {
		h.v1Proxy(w, r)
		return
	}

	// if it's /url then resolve by id
	if strings.HasPrefix(r.URL.Path, "/url/") {
		h.urlProxy(w, r)
	}
}

func New() *Handler {
	h := &Handler{
		appMap:  make(map[string]*backend),
		funcMap: make(map[string]*backend),
	}
	return h
}
