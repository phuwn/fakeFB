package server

import (
	"github.com/jinzhu/gorm"

	"github.com/phuwn/fakeFB/pkg/config/logger"
	"github.com/phuwn/fakeFB/pkg/store"
)

// Server of app
type Server struct {
	db     *gorm.DB
	store  *store.Store
	logger *logger.LoggingService
}

// srv store a global config of server
var srv Server

// NewServer init server of app
func NewServer(db *gorm.DB, store *store.Store, logger *logger.LoggingService) *Server {
	srv.db = db
	srv.store = store
	srv.logger = logger
	return &srv
}

// InitLogger init logger to server of testing
func InitLogger(logger *logger.LoggingService) *Server {
	srv.logger = logger
	return &srv
}

// GetSrvConfig get config of server
func GetSrvConfig() *Server {
	return &srv
}

// DB get database of server
func (s *Server) DB() *gorm.DB {
	return s.db
}

// Store access to stores of server
func (s *Server) Store() *store.Store {
	return s.store
}

// Logger log all activity for server
func (s *Server) Logger() *logger.LoggingService {
	return s.logger
}
