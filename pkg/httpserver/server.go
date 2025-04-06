package httpserver

import (
	"fmt"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"testTaskGravitum/pkg/logger"
)

type Server struct {
	App *http.Server
	log *slog.Logger
	Mux *http.ServeMux
}

func New(log *slog.Logger, host string, port string) *Server {
	mux := http.NewServeMux()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(mux)
	loggedMux := logger.LoggingMiddleware(log)(handler)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: loggedMux,
	}

	return &Server{
		log: log,
		App: httpServer,
		Mux: mux,
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
