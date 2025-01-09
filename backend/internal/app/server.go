package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (app *application) serveHTTP() error {

	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), app.cfg.Server.ShutdownPeriod)
		defer cancel()

		shutdownErrChan <- app.echo.Shutdown(ctx)

	}()

	app.logger.Info("🚀 server started", "env", app.cfg.App.Env, "address", app.cfg.Server.Addr)
	if err := app.echo.Start(app.cfg.Server.Addr); err != nil && err != http.ErrServerClosed {
		return err
	}

	err := <-shutdownErrChan
	if err != nil {
		return err
	}
	app.wg.Wait()
	app.logger.Warn("server stopped gracefully")
	return nil
}
