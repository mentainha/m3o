# API Ingress

Ingress for M3O services

## Overview

The M3O API Ingress is a single entrypoint for apps, functions and more hosted on the M3O platform.

Supported endpoints below:

- [api](#api)
- [apps](#apps)
- [functions](#functions)
- [url shortener](#url-shortener)
- [user verification](#user-verification)

## API

We provide a proxy to the /v1/* API endpoints. Which can be called via the following endpoint:

```
https://m3o.one/api/[service]/[endpoint]
```

Which will route to:

```
https://api.m3o.com/v1/[service]/[endpoint]
```

Where `Authorization: Bearer XXX` is present it will be passed through to the client. 

Otherwise we provide GET/POST access to the /v1/* API via `api_key=xxx` http param.

### Example URL

```
https://m3o.one/api/helloworld/call?api_key=xxxx
```

Routes to `https://api.m3o.com/v1/helloworld/call` with `Authorization: Bearer XXX` set.

### Request Params

The `/api` endpoint supports both HTTP params and JSON body where `Content-Type: application/json` is specified.

## Apps

Apps are given a unique id and subdomain using m3o.app e.g [helloworld.m3o.app](https://helloworld.m3o.app) resolves to the app id helloworld.

The `/v1/app/Resolve` endpoint is called with the ID specified to resolve the backend URL for the app

## Functions

Functions are given a unique id and subdomain much like apps using m3o.sh e.g [helloworld.m3o.sh](https://helloworld.m3o.sh/) resolves to the 
function helloworld. 

The `/v1/function/Proxy` endpoint is called with the ID specified to resolve the backend URL for the function

## URL Shortener

The [URL](https://github.com/micro/services/tree/master/url) service provides link shortening and sharing. The URL Proxy fronts those urls 
as a single entrypoint at https://m3o.one. 

- Assumes url is of format `https://m3o.one/u/AArfeZE`
- Will call `https://api.m3o.com/url/resolve?shortURL=https://m3o.one/u/AArfeZE`
- URL service should return `destinationURL=https://foobar.com/example`
- Proxy will issue a 302 redirect

URLs can also be served via `/url/[id]` e.g `https://m3o.one/url/xyz`

## User Verification

User email verification is performed via the `/v1/user/VerifyEmail` endpoint. The ingress serves `user.m3o.com` for this purpose.

This is also available via the `/user/[email]` endpoint

## Usage

Specify `M3O_API_TOKEN` to the binary and run the ingress

### Micro Services

Set the config values 

```
micro config set micro.app.domain m3o.app
micro config set micro.app.domain m3o.sh
micro config set micro.url.host_prefix https://m3o.one/u/
micro config set micro.user.verify_email_url https://user.m3o.com
```

Deploy the url, function, app, user services

```
micro run github.com/micro/services/app
micro run github.com/micro/services/function
micro run github.com/micro/services/url
micro run github.com/micro/services/user
```
