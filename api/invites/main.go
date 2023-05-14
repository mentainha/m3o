package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/invites/handler"
	pb "m3o.dev/api/invites/proto"
)

func main() {
	srv := service.New(
		service.Name("invites"),
	)

	// Connect to the database
	db, err := sql.NewDB("invites")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Invite{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterInvitesHandler(srv.Server(), &handler.Invites{DB: db.DB})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
