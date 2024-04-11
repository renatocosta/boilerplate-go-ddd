package support

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func ShutdownApp(ctx context.Context, cleanup func(), errorApp *string) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		slog.Info("signal.Notify", v)
	case <-ctx.Done():
		cleanup()
		if errorApp != nil && *errorApp != "" {
			slog.Info(*errorApp)
		}
		//slog.Info("ctx.Done: ", done)
	}

}
