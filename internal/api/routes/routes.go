package routes

import (
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
	"gorm.io/gorm"
)

type Routes struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRoutes(logger *logger.Logger, db *gorm.DB) *Routes {
	return &Routes{
		db:     db,
		logger: logger,
	}
}
