package users

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api/db"
	"m3o.dev/apps/chat/api/users/handler"
	pb "m3o.dev/apps/chat/api/users/proto"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("users")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.User{}, &handler.Token{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterUsersHandler(srv.Server(), &handler.Users{DB: db, Time: time.Now})
}
