package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/seen/handler"
	pb "m3o.dev/api/seen/proto"
)

func main() {
	srv := service.New(
		service.Name("seen"),
	)

	// Connect to the database
	db, err := sql.New("seen")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.SeenInstance{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterSeenHandler(srv.Server(), &handler.Seen{DB: db.Debug()})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
