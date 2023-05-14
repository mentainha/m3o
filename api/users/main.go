package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/users/handler"
	pb "m3o.dev/api/users/proto"
)

func main() {
	srv := service.New(
		service.Name("users"),
	)

	// Connect to the database
	db, err := sql.NewDB("users")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.User{}, &handler.Token{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterUsersHandler(srv.Server(), &handler.Users{DB: db, Time: time.Now})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
