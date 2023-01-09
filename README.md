# M3O [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Universal Micro services gateway

## Overview

M3O transforms public APIs into smaller easier to use Micro services which you can explore, discover and consume 
as serverless building blocks. The platform generates unified api docs, client libraries and examples on top 
of [Micro](https://github.com/micro/micro) services using protobuf to OpenAPI conversion and some custom tooling. 

## Features

Here are the main features:

- **1️⃣ ONE Platform** - Discover, explore and consume public APIs all in one place. 
- **☝️ ONE Account** - Manage your API usage with one account and one token.
- **⚡ ONE Framework** - Learn, develop and integrate using one set of docs and libraries.

## Services

Here are some of the APIs:

- [**Apps**](https://m3o.com/app) - Serverless app deployment
- [**Database**](https://m3o.com/db) - Serverless postgres database
- [**Functions**](https://m3o.com/function) - Serverless lambda functions
- [**Users**](https://m3o.com/user) - User management and authentication
- [**Email**](https://m3o.com/email) - Send emails in a flash
- [**SMS**](https://m3o.com/sms) - Send an SMS message

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
  -d '{"name": "helloworld", "repo": "github.com/m3o/helloworld"}' \
  https://api.m3o.com/v1/app/Run
```

Find all the curl examples in [m3o-sh](https://github.com/m3o/m3o-sh)

### Go

Import packages from [go.m3o.com](https://pkg.go.dev/go.m3o.com)

```go
import (
        "go.m3o.com"
        "go.m3o.com/app"
)
```

Create a new client with your API token and call it

```go
client := m3o.New(os.Getenv("M3O_API_TOKEN"))
rsp, err := client.App.Run(&app.RunRequest{
	Name: "helloworld",
	Repo: "github.com/m3o/helloworld"
})
fmt.Println(rsp, err)
```

Find all the Go examples in [m3o-go](https://github.com/m3o/m3o-go)

### JS

Install the [m3o](https://www.npmjs.com/package/m3o) package

```
npm install m3o
```

Call app run like so

```javascript
const m3o = require("m3o").default(process.env.M3O_API_TOKEN);

// Call returns a personalised "Hello $name" response
async function main() {
  let rsp = await m3o.app.run({
    name: "helloworld",
    repo: "github.com/m3o/helloworld"
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
m3o app run --name=helloworld --repo=github.com/m3o/helloworld
```

See the [m3o-cli](https://github.com/m3o/m3o-cli/tree/main/examples) for examples

## Self Hosting

We're still working on configuration for self hosting
