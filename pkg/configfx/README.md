# configfx

https://github.com/caarlos0/env

## Usage

```go
package main

import (
	"fmt"
	"cotacao-fretes/pkg/configfx"
	"go.uber.org/fx"
)

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
}

func main() {
	app := fx.New(
		configfx.Module[Config](),
		fx.Invoke(func(cfg Config) {
			fmt.Println(cfg)
		}),
	)

	app.Run()
}
```
