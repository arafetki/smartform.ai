package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		router: router,
	}
}

func (server *Server) Start(address string) error {

	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
		defer cancel()

		shutdownErrChan <- server.router.Shutdown(ctx)

	}()

	err := server.router.Start(address)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	err = <-shutdownErrChan
	if err != nil {
		return err
	}
	server.router.Logger.Info("server has stopped")
	return nil
}
