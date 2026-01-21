package slogfx

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

type Config struct {
	Level       string   `env:"LOG_LEVEL" envDefault:"info"`
	Channel     string   `env:"LOG_CHANNEL" envDefault:"default"`
	Application string   `env:"LOG_APPLICATION" envDefault:"local"`
	Env         string   `env:"LOG_ENV" envDefault:"dev"`
	Tags        []string `env:"LOG_TAGS" envDefault:"dev,local"`
}

func SetLevel(lvl *slog.LevelVar, logLevel string) error {
	return lvl.UnmarshalText([]byte(logLevel))
}

func New(cfg Config) (*slog.Logger, *slog.LevelVar, error) {
	lvl := new(slog.LevelVar)

	if err := lvl.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, nil, err
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "time":
				t := a.Value.Time().UTC().Format(time.RFC3339)
				return slog.Attr{Key: "timestamp", Value: slog.StringValue(t)}
			case "msg":
				return slog.Attr{Key: "message", Value: a.Value}
			case "level":
				return slog.Attr{Key: "level", Value: slog.StringValue(strings.ToLower(a.Value.String()))}
			default:
				return a
			}
		},
	}))

	l = l.With(
		"channel", cfg.Channel,
		"application", cfg.Application,
		"environment", cfg.Env,
		"tags", cfg.Tags,
	)

	slog.SetDefault(l)

	return l, lvl, nil
}
