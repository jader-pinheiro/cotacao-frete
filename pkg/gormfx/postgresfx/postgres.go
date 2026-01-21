package postgresfx

import (
	"cotacao-fretes/pkg/gormfx"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg gormfx.Config) (gorm.Dialector, error) {
	return postgres.Open(cfg.DSN), nil
}
