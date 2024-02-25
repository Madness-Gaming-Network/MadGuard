package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	// TODO: How do you do "idiomatic" config / globals in Go?
	var TIMEOUT_IN_SECONDS time.Duration = 5
	// os.Getenv("TIMEOUT_IN_SECONDS") ???

	// https://gin-gonic.com/docs/quickstart/
	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO: Set to $PORT env var
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error during listen and serve: %s\n", err)
		}
	}()

	// TODO: Extract all this interrupt stuff into a helper?

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill with no args sends syscanll.SIGTERM (default)
	// kill with -2 sends:syscall.SIGINT
	// kill with -9 is syscall.SIGKILL but cant be caught, so we don't' add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Received quit signal\n")

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_IN_SECONDS*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %s\n", err)
	}

	<-ctx.Done()
	log.Printf("Caught ctx.Done(), timeout of %d seconds, exiting...\n", TIMEOUT_IN_SECONDS)

}
