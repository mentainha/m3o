# M3O Chat

Private and group messaging platform

## Overview

A proof of concept video messaging app built on top of Micro. It was primarily geared towards distributed teams,
as it was written during the pandemic based on our own experiences. Messaging is the main component with video as small bubbles above
the chat window. It's group based with rooms and private DMs built in.

## Install

### Micro Server

Download the [latest release](https://github.com/micro/micro/releases/latest) of Micro

```
micro server
```

Login with default username/password: admin/micro

```
micro login
```

### Setup API

Start the api

```bash
for api in chats codes groups invites seen streams threads users; do
  ## generate an auth rule to open the api
  micro auth create rule --resource=service:$api:* --access=granted --priority=1 $api

  ## run the service
  micro run github.com/micro/chat/api/$api
done
```

### Optional Configuration

Set the environment variables by creating `.env.local` file in the repo root with the following:

Optionally configure sendgrid for invite and password reset emails

```
SENDGRID_API_KEY=xxxxxxxx
```

Optionally configure twilio for audio/video calls

```
TWILIO_API_KEY=xxxxxxx
TWILIO_API_SECRET=xxxxx
TWILIO_ACCOUNT_SID=xxxxxx
```

Optionally configure a remotely authenticated Micro API

```
MICRO_API_ENDPOINT=xxxxxxx
MICRO_API_KEY=xxxxxx$(micro user token)
MICRO_API_NAMESPACE=xxxxxx
```

### Start the app

```
npm run dev
```

The application is accessible on http://localhost:3000
