package groups

import (
	"m3o.dev/apps/chat/api/db"
	"m3o.dev/apps/chat/api/groups/handler"
	pb "m3o.dev/apps/chat/api/groups/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("groups")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Group{}, &handler.Membership{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterGroupsHandler(srv.Server(), &handler.Groups{DB: db.Debug()})
}
