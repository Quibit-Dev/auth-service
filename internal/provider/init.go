package provider

import (
	"auth-service/internal/config"
	"auth-service/internal/database"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func InitGlobal(cfg *config.Config) {
	db = database.InitPostgres(cfg)
}
