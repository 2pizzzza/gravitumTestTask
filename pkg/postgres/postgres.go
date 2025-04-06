//nolint:all
package postgres

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"testTaskGravitum/internal/config"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cnf *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cnf.DB.Username, cnf.DB.Password, cnf.DB.Host, cnf.DB.Port, cnf.DB.Database,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse db config: %w", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Minute * 5

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return pool, nil
}

func RunMigrations(cfg *config.Config) error {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)

	m, err := migrate.New(
		"file://database/migration",
		url,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func (p *Postgres) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return p.Pool.Query(ctx, query, args...)
}

func (p *Postgres) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.Pool.QueryRow(ctx, query, args...)
}
