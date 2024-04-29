package support

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func ShutdownHandler(ctx context.Context, ctxCancel context.CancelFunc, srv *http.Server) {

	// Listen for OS signals to gracefully shut down the server
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	defer ctxCancel()

	// Shutdown the server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %s\n", err)
	}
	log.Println("Server gracefully stopped")

}
