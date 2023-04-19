package streams

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api/db"
	"m3o.dev/apps/chat/api/streams/handler"
	pb "m3o.dev/apps/chat/api/streams/proto"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("streams")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Token{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterStreamsHandler(srv.Server(), &handler.Streams{
		DB:     db,
		Events: events.DefaultStream,
		Time:   time.Now,
	})
}
