# Gorm Module

https://github.com/go-gorm/gorm

## Usage

```go
package main

import (
	"fmt"
	"cotacao-fretes/pkg/gormfx"
	"cotacao-fretes/pkg/gormfx/gormfx_mysql"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type AModel struct {
	gorm.Model
	Name string 
}

func main() {
	app := fx.New(
		gormfx.Module(),
		gormfx_mysql.Module(), // using MySQL
		gormfx.Migrate(&AModel{}),
		fx.Invoke(func(db *gorm.DB) {
			fmt.Println(db)
		}),
	)

	app.Run()
}
```

## Configuration

Check ./gorm.go#Config
