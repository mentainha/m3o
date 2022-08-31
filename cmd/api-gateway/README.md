# API Gateway

Ingress for apps and functions

## Overview

The M3O API Gateway is a single entrypoint for apps, functions and more. It manages the usage of custom domains.

## Apps

Apps are given a unique id and subdomain using m3o.app e.g [helloworld.m3o.app](https://helloworld.m3o.app) resolves to the app id helloworld.

## Functions

Functions are given a unique id and subdomain much like apps using m3o.sh e.g [helloworld.m3o.sh](https://helloworld.m3o.sh/) resolves to the 
function helloworld. 

## URL

The [url](https://github.com/micro/services/tree/master/url) service provides link shortening and sharing. The URL Proxy fronts those urls 
as a single entrypoint at https://m3o.one. We don't serve directly because of time wasted on ssl certs, etc.

- Assumes url is of format `https://m3o.one/u/AArfeZE`
- Will call `https://api.m3o.com/url/resolve?shortURL=https://m3o.one/u/AArfeZE`
- URL service should return `destinationURL=https://foobar.com/example`
- Proxy will issue a 301 redirect

## Usage

Set the host prefix

```
micro config set micro.url.host_prefix https://m3o.one/u/
micro config set micro.app.domain m3o.app
```

Deploy the url, function, app services

```
micro run github.com/micro/services/app
micro run github.com/micro/services/function
micro run github.com/micro/services/url
```
