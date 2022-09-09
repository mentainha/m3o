# Micro Platform

A build of the [Micro](https://github.com/micro/micro) platform including our v1 plugin

## Overview

This build of the micro platform includes the v1 plugin which short circuits the request flow to 
bypass proxying and directly pass through requests from the micro api to the v1 api. Removing 
the extra network hop plus internal routing should speedup requests.

## Usage

Use [M3O Cloud](https://github.com/m3o/cloud) pulumi automation to boot the platform

Alternatively start the server with `MICRO_PROFILE=platform` if the [platform](https://github.com/micro/micro/blob/master/profile/platform/platform.go) 
profile dependencies are available. 

Set the following for v1 routing

```
micro config set micro.v1.address [address of v1 service]
```

