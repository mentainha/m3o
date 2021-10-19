# M3O API

The M3O API consists of a set of many separate public apis all consumed through a single gateway.

## Overview

M3O APIs are a standard set of http/json APIs which act as programmable building blocks 
for rapid development of any product or services. Pick up and use one or more APIs as 
easily as importing a library and making a function call from a library.

## Examples

Here's a simple helloworld

Curl

```
curl -H "Authorization: Bearer $M3O_API_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name": "John"}' \
     https://api.m3o.com/v1/helloworld/call
```

Go

```go
token := os.Getenv("M3O_API_TOKEN")

query := map[string]interface{}{
	"name": "John",
}
b, _ := json.Marshal(query)

req, _ := http.NewRequest("POST", "https://api.m3o.com/v1/helloworld/call", bytes.NewReader(b))
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+token)

rsp, err := http.DefaultClient.Do(req)
```

See the [examples](../examples) for more use cases.

## API Endpoint

The canonical API endpoint is

```
https://api.m3o.com/v1/
```

All service endpoints are append like so. 

```
# Assuming helloworld with Call endpoint

https://api.m3o.com/v1/helloworld/Call
```

## Authorization

All API calls require a valid API token. New tokens can be generated on the [API settings](https://m3o.com/settings/keys) page.

An API token should be passed in the `Authorization: Bearer` header

```
Authorization: Bearer $MICRO_API_TOKEN
```

## Data Format

All request/responses are in JSON format and require a `Content-Type: application/json` header to be passed on each request.

## Public APIs

A list of public APIs can be found on [m3o.com/explore](https://m3o.com/explore)
