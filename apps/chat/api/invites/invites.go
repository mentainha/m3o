package invites

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api/db"
	"m3o.dev/apps/chat/api/invites/handler"
	pb "m3o.dev/apps/chat/api/invites/proto"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("invites")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Invite{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterInvitesHandler(srv.Server(), &handler.Invites{DB: db})
}
