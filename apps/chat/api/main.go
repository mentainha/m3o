package main

import (
	"log"

	"github.com/micro/micro/v3/service"

	"m3o.dev/apps/chat/api/chats"
	"m3o.dev/apps/chat/api/groups"
	"m3o.dev/apps/chat/api/invites"
	"m3o.dev/apps/chat/api/seen"
	"m3o.dev/apps/chat/api/streams"
	"m3o.dev/apps/chat/api/threads"
	"m3o.dev/apps/chat/api/users"
)

func main() {
	srv := service.New(
		service.Name("chat"),
	)

	// register chats handler
	chats.Register(srv)
	// register groups handler
	groups.Register(srv)
	// register invites handler
	invites.Register(srv)
	// register seen handler
	seen.Register(srv)
	// register streams handler
	streams.Register(srv)
	// register threads handler
	threads.Register(srv)
	// register users handler
	users.Register(srv)

	// run the service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
