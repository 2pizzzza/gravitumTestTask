package httpserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"testTaskGravitum/pkg/logger"
	"time"
)

const (
	_defaultAddr            = ":80"
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	App *http.Server
	log *slog.Logger

	address         string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

func New(log *slog.Logger, host string, port string) *Server {
	_ = http.NewServeMux()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: nil,
	}

	return &Server{
		log: log,
		App: httpServer,
	}
}

func (a *Server) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *Server) Run() error {
	if err := a.App.ListenAndServe(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *Server) Stop() {
	if err := a.App.Close(); err != nil {
		a.log.Error("error while closing HTTP httpserver", logger.Err(err))
	}
}
