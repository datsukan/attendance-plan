package infrastructure

import (
	"log/slog"
	"os"

	"github.com/oklog/ulid/v2"
)

func NewLogger() *slog.Logger {
	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &handlerOptions))

	requestID := ulid.Make().String()
	logger.With("request_id", requestID)

	return logger
}
