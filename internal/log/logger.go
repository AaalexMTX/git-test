package log

import (
	"log/slog"
	"os"
	"time"
)

var Logger *slog.Logger

func LoggerInit() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().Format(time.DateTime))
			}
			return a
		},
	}))
	slog.SetDefault(Logger)
}
