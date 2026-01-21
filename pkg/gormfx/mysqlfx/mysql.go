package mysqlfx

import (
	"cotacao-fretes/pkg/gormfx"
	"os"

	mysqldriver "github.com/go-sql-driver/mysql"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg gormfx.Config) (gorm.Dialector, error) {
	sqltrace.Register("mysql", &mysqldriver.MySQLDriver{}, sqltrace.WithServiceName(os.Getenv("DD_SERVICE")))
	db, err := sqltrace.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}
	return mysql.New(mysql.Config{
		Conn: db,
	}), nil
}
