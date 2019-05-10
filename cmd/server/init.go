package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/phuwn/fakeFB/pkg/config/database"
	"github.com/phuwn/fakeFB/pkg/config/logger"
	"github.com/phuwn/fakeFB/pkg/server"
	"github.com/phuwn/fakeFB/pkg/store"
	"github.com/phuwn/fakeFB/pkg/store/user"
)

// Init server
func Init() (*server.Server, func()) {
	pgdb, cleanup := database.New(os.Getenv("PG_DATASOURCE"))
	store := newStore()
	logger := logger.NewLogger()

	return server.NewServer(pgdb, store, logger), cleanup
}

func newStore() *store.Store {
	return &store.Store{
		User: user.NewPGService(),
	}
}
