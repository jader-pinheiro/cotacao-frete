package mysqlfx_test

import (
	"context"
	"cotacao-fretes/pkg/gormfx"
	"cotacao-fretes/pkg/gormfx/mysqlfx"
	"fmt"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlContainer, err := mysql.RunContainer(
		ctx,
		testcontainers.WithImage("mysql:8.0.36"),
		mysql.WithDatabase("foo"),
		mysql.WithPassword("root"),
	)
	if err != nil {
		panic(err)
	}
	defer mysqlContainer.Terminate(ctx)
	dsn, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("dsn:", dsn)
	os.Setenv("GORM_DSN", dsn)

	res := m.Run()
	os.Exit(res)
}

type Person struct {
	gorm.Model
	Name string
}

func TestFx(t *testing.T) {
	app := fxtest.New(
		t,
		gormfx.Module(),
		mysqlfx.Module(),
		fx.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&Person{})
		}),
	)

	app.RequireStart()

	app.RequireStop()
}
