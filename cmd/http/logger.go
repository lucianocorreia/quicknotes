package main

import (
	"io"
	"log/slog"
	"time"
)

func newLogger(out io.Writer, lavel slog.Level) *slog.Logger {
	log := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
		AddSource:   true,
		Level:       lavel,
		ReplaceAttr: replaceTimeFormat,
	}))

	return log
}

func replaceTimeFormat(group []string, a slog.Attr) slog.Attr {
	if a.Key == "time" {
		// value := time.Now().Format(time.RFC3339)
		value := time.Now().Format("2006-01-02 15:04:05")
		return slog.Attr{Key: a.Key, Value: slog.StringValue(value)}
	}

	return slog.Attr{Key: a.Key, Value: a.Value}
}
