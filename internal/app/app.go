package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testTaskGravitum/internal/config"
	"testTaskGravitum/pkg/httpserver"
	"testTaskGravitum/pkg/logger"
)

func New(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	application := httpserver.New(log, cfg.App.Host, cfg.App.Port)

	go application.MustRun()

	stop := make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	log.Info("Server is dead")

}
