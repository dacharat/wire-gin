package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dacharat/wire-gin/di"
	"github.com/dacharat/wire-gin/routers"
	"github.com/dacharat/wire-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func gracefullyShutdown(router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(errors.Wrap(err, "Cannot open socket"))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := utils.GetContext()
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(errors.Wrap(err, "Forced shutdown"))
	}
}

func main() {
	api := di.New()

	gracefullyShutdown(routers.New(api))
}
