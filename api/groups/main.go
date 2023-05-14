package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/groups/handler"
	pb "m3o.dev/api/groups/proto"
)

func main() {
	srv := service.New(
		service.Name("groups"),
	)

	// Connect to the database
	db, err := sql.NewDB("groups")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Group{}, &handler.Membership{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterGroupsHandler(srv.Server(), &handler.Groups{DB: db.Debug()})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
