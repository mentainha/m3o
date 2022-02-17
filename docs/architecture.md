# Architecture

Read below to learn more about how M3O works

## Infrastructure

M3O is built on existing public cloud infrastructure using managed kubernetes along with our own [platform automation](https://github.com/m3o/platform). 
We host the open source project [Micro](https://github.com/micro/micro) as our base cloud OS and use it to power all the [Micro Services](https://github.com/micro/services).

## Control Plane

We host our own [Dev UX](#dev-ux) on top of the infrastructure stack and a [backend](https://github.com/m3o/backend) 
which acts as the management control plane. This productizes the entire offering and allows for publishing of services with configurable pricing.

## Micro Services

Developers build and contribute to [Micro Services](https://github.com/micro/services), which act as an abstraction layer for existing third party 
public APIs. We then automate the building and publishing of those services and generate client libraries for them all. 

## Development

This project is a combination of open source projects and a platform managed by the M3O team. Our goal is to enable any developer to 
contribute to the open source while benefiting from the platform as a shared resource.

## Source

The core services, automation and the backend are all open source:

- [services](https://github.com/micro/services) - offered on the M3O platform
- [platform](https://github.com/m3o/platform) - infrastructure automation
- [backend](https://github.com/m3o/backend) - powering the M3O platform

## Dev UX

We provide the following dev UX for the consumption of Micro services:

- [m3o-web](https://github.com/m3o/m3o-web) - Next.js based Web dashboard which can be self hosted
- [m3o-js](https://github.com/m3o/m3o-js) - JS client library with statically typed interfaces and examples
- [m3o-go](https://github.com/m3o/m3o-go) - Go client library with code generated functions and examples
- [m3o-cli](https://github.com/m3o/m3o-cli) - Command line interface for terminal access to services
