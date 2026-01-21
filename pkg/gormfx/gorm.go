package gormfx

import (
	"context"
	"cotacao-fretes/pkg/probes"

	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Config struct {
	DSN string `env:"GORM_DSN" envDefault:""`
}

func New(d gorm.Dialector) (*gorm.DB, error) {
	return gormtrace.Open(d, &gorm.Config{})
}

func Check(db *gorm.DB) probes.Check {
	return func(ctx context.Context) error {
		d, err := db.DB()
		if err != nil {
			return err
		}

		return d.PingContext(ctx)
	}
}

func Migrate(models ...any) fx.Option {
	return fx.Options(
		fx.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(models...)
		}),
	)
}

func CreateInitData(structs any, data any) fx.Option {
	return fx.Options(
		fx.Invoke(func(db *gorm.DB) error {
			var count int64
			db.Model(&structs).Count(&count)
			if count == 0 {
				return db.Create(data).Error
			}
			return nil
		}),
	)
}
