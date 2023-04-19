package seen

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api/db"
	"m3o.dev/apps/chat/api/seen/handler"
	pb "m3o.dev/apps/chat/api/seen/proto"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("seen")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.SeenInstance{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterSeenHandler(srv.Server(), &handler.Seen{DB: db.Debug()})
}
