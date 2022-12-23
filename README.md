# M3O [![Website](https://img.shields.io/:signup-free-orange.svg)](https://m3o.com) [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Discord](https://img.shields.io/:join-discord-purple.svg)](https://discord.gg/TBR9bRjd6Z) [![Mastodon](https://img.shields.io/:mastodon-server-black.svg)](https://M3O.org) [![Twitter](https://img.shields.io/:twitter-follow-blue.svg)](https://twitter.com/m3oservices)

Universal Micro services gateway

## Overview

M3O transforms public APIs into smaller easier to use Micro services which you can explore, discover and consume 
as serverless building blocks. The platform generates unified api docs, client libraries and examples on top 
of [Micro](https://github.com/micro/micro) services using protobuf to OpenAPI conversion and some custom tooling. 
The API gateway manages access using a pay per request billing model by debiting from a virtual wallet which 
you can topup via a standard credit or debit card.

## Features

Here are the main features:

- **1️⃣ ONE Platform** - Discover, explore and consume public APIs all in one place. 
- **☝️ ONE Account** - Manage your API usage with one account and one token.
- **⚡ ONE Framework** - Learn, develop and integrate using one set of docs and libraries.

## Services

So far there are over 60+ services. Here are some of the highlights:

### ☁️ Cloud

- [**Apps**](https://m3o.com/app) - Serverless app deployment
- [**Database**](https://m3o.com/db) - Serverless postgres database
- [**Functions**](https://m3o.com/function) - Serverless lambda functions
- [**Events**](https://m3o.com/event) - Publish and subscribe to messages
- [**Users**](https://m3o.com/user) - User management and authentication
- [**Space**](https://m3o.com/space) - Infinite cloud storage
- [**Search**](https://m3o.com/search) - Indexing and full text search

### 🗺️ Logistics

- [**Address**](https://m3o.com/address) - Address lookup by postcode
- [**Geocoding**](https://m3o.com/geocoding) - Geocode an address to gps location and the reverse.
- [**Location**](https://m3o.com/location) - Real time GPS location tracking and search
- [**Places**](https://m3o.com/place) - Search for geographic points of interest
- [**Routes**](https://m3o.com/routing) - Etas, routes and turn by turn directions
- [**IP2Geo**](https://m3o.com/ip) - IP to geolocation lookup

### ⚙️ Miscellaneous

- [**Email**](https://m3o.com/email) - Send emails in a flash
- [**Image**](https://m3o.com/image) - Quickly upload, resize, and convert images
- [**OTP**](https://m3o.com/otp) - One time password generation
- [**QR Codes**](https://m3o.com/qr) - QR code generator
- [**SMS**](https://m3o.com/sms) - Send an SMS message
- [**Weather**](https://m3o.com/weather) - Real time weather forecast

See the full list at [m3o.com/explore](https://m3o.com/explore).

## Getting Started

- Head to [m3o.com](https://m3o.com) and signup for a free account.
- Browse services on the [Explore](https://m3o.com/explore) page.
- Call any service using your token in the `Authorization: Bearer [Token]` header
- All services are available through one API endpoint: `https://api.m3o.com/v1/*`.
- Use [m3o-cli](https://github.com/m3o/m3o-cli), [m3o-js](https://github.com/m3o/m3o-js) and [m3o-go](https://github.com/m3o/m3o-go) clients for development

## Quick Start

Grab your API token from the dashboard and export it

```
export M3O_API_TOKEN=xxxxxxx
```

### Curl

```bash
curl \
  -H "Authorization: Bearer $M3O_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "John"}' \
  https://api.m3o.com/v1/helloworld/call
```

Find all the curl examples in [m3o-sh](https://github.com/m3o/m3o-sh)

### Go

Import packages from [go.m3o.com](https://pkg.go.dev/go.m3o.com)

```go
import (
        "go.m3o.com"
        "go.m3o.com/helloworld"
)
```

Create a new client with your API token and call it

```go
client := m3o.New(os.Getenv("M3O_API_TOKEN"))
rsp, err := client.Helloworld.Call(&helloworld.CallRequest{
	Name: "John",
})
fmt.Println(rsp, err)
```

Find all the Go examples in [m3o-go](https://github.com/m3o/m3o-go)

### JS

Install the [m3o](https://www.npmjs.com/package/m3o) package

```
npm install m3o
```

Call helloworld like so

```javascript
const m3o = require("m3o").default(process.env.M3O_API_TOKEN);

// Call returns a personalised "Hello $name" response
async function main() {
  let rsp = await m3o.helloworld.call({
    name: "John",
  });
  console.log(rsp);
}

main();
```

Find more JS examples in [m3o-js](https://github.com/m3o/m3o-js)

### CLI

Install the cli

```
curl -fssl https://install.m3o.com/cli | /bin/bash
```

Example call

```
m3o helloworld call --name=Alice
```

See the [m3o-cli](https://github.com/m3o/m3o-cli/tree/main/examples) for examples

## Learn More

- Signup on the [Website](https://m3o.com)
- Follow us on [Twitter](https://twitter.com/m3oservices)
- Ask questions on [Discord](https://discord.gg/TBR9bRjd6Z)
- Join the community on [Mastodon](https://m3o.org)

## Development

[1 click deploy](https://marketplace.digitalocean.com/apps/micro) a Micro Dev environment on a DigitalOcean Droplet

Use our [refcode](https://marketplace.digitalocean.com/apps/micro?refcode=1eb1b2aca272&action=deploy) so we get $25 credit too!
