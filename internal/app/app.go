package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testTaskGravitum/internal/config"
	"testTaskGravitum/internal/http/handler"
	"testTaskGravitum/internal/service/user"
	internalPostgres "testTaskGravitum/internal/storage/postgres"
	"testTaskGravitum/pkg/httpserver"
	"testTaskGravitum/pkg/logger"
	pkgPostgres "testTaskGravitum/pkg/postgres"
)

func New(cfg *config.Config) {
	ctx := context.Background()
	l := logger.New(cfg.Log.Level)

	application := httpserver.New(l, cfg.App.Host, cfg.App.Port)

	conn, err := pkgPostgres.New(ctx, cfg)
	if err != nil {
		l.Error("Failed connect to database")
		application.Stop()
	}
	l.Info("Server is live")

	err = pkgPostgres.RunMigrations(cfg)
	if err != nil {
		l.Error("Failed apply migrations")
		fmt.Println(err)
		application.Stop()
	}
	_ = internalPostgres.New(conn)

	repo := internalPostgres.NewUsersRepository(conn)

	service := user.New(repo)

	handlers := handler.New(service)

	handlers.RegisterRouter(application.Mux)

	go application.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	l.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	l.Info("Server is dead")

}
