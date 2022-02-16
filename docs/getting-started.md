# Getting Started

A quick guide to getting started with M3O

## Sign Up

[Sign up](https://m3o.com/register) or [Log in](https://m3o.com/login) to M3O so you can get a token. You can use password based log in or oauth.

Note: For security purposes, if you log in with oauth, the password is no longer valid unless you reset it.

## Try out Hello world!

Here's a simple example to call the helloworld API:

Hello World example

```shell
curl "https://api.m3o.com/v1/helloworld/Call" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "name": "John"
}'
```

## Discover APIs

Browse all of our public APIs on the [Explore](https://m3o.com/explore) page.

## Generate an API token

To create a new access token, visit [API keys](https://m3o.com/account/keys) page. You can specify which APIs the token can call or leave scopes blank to create a token that has access to all the APIs.

## Using an API

M3O APIs can be accessed via the url 

```
https://api.m3o.com/v1/[service]/[endpoint]
```

All API calls require an API token provided via the `Authorization: Bearer` header. 

Our APIs provide a JSON based request/response format and also require a `Content-Type: application/json` header for each request.

After familiarising yourself with an API by reading the API reference (e.g https://m3o.com/db/api#Read) you probably want to start calling the API. M3O APIs can be called via HTTP in any language or using our node.js and Golang specific clients.

## Download the CLI

M3O includes a command line interface from which you can explore and query services. 

Go to https://github.com/m3o/m3o/releases/latest to download the latest release binary.

## Account Billing
The majority of our APIs are free, with a 1 million request per month quota. Third party APIs or calls that perform specific functions such as SMS or email are charged for and billed per request. Each account has a credit based balance. Any request made to a paid API will be debited from the balance in real time. When your balance is zero requests to paid APIs will return a blocked status.

Additional API calls beyond the free quota are charged at $0.000001 per request (aka $1 per 1 million requests). 

To check your balance and top-up your account head to the Billing page.

## Fair Usage Policy

Our free tier is rate limited to 100 requests/second. This can be unlocked to 1000 requests/second by upgrading to the Pro subscription. We also invoke fair usage limits on compute and storage. During the beta period no hard limits are enforced, we will reach out if there are any issues.

## Feedback or Support

Join the community and ask questions on Discord. If you're subscribed to the Pro tier email support@m3o.com for help. Otherwise you can provide feedback at contact@m3o.com.
