package db

import (
	"os"
	"path/filepath"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/util/user"
)

func New(name string) (*gorm.DB, error) {
	dbAddress := "sqlite://" + name + ".db"

	// Connect to the database
	cfg, err := config.Get(name + ".database")
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	addr := cfg.String(dbAddress)

	var db *gorm.DB

	if strings.HasPrefix(addr, "postgres") {
		db, err = gorm.Open(postgres.Open(addr), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: name + "_",
			},
		})
	} else {
		path := filepath.Join(user.Dir, "service", name)
		os.MkdirAll(path, 0755)
		file := filepath.Join(path, "db.sqlite")
		db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	}

	return db, err
}
