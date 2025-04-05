package app

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testTaskGravitum/internal/config"
	"testTaskGravitum/internal/service/user"
	internalPostgres "testTaskGravitum/internal/storage/postgres"
	"testTaskGravitum/pkg/httpserver"
	"testTaskGravitum/pkg/logger"
	pkgPostgres "testTaskGravitum/pkg/postgres"
)

func New(cfg *config.Config) {
	ctx := context.Background()
	l := logger.New(cfg.Log.Level)

	conn, err := pkgPostgres.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	l.Info("Server is live")

	_ = internalPostgres.New(conn)

	repo := internalPostgres.NewUsersRepository(conn)

	_ = user.New(repo)

	application := httpserver.New(l, cfg.App.Host, cfg.App.Port)

	go application.MustRun()

	stop := make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	l.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	l.Info("Server is dead")

}
