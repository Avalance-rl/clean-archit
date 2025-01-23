package pgx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Pgx struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPgx(ctx context.Context, dsn string, maxConns int32, logger *zap.Logger) *Pgx {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	config.MaxConns = maxConns

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	if err = pool.Ping(ctx); err != nil {
		logger.Fatal("failed to ping database", zap.Error(err))
	}

	return &Pgx{
		pool:   pool,
		logger: logger,
	}
}

func (p *Pgx) Close(ctx context.Context) error {
	done := make(chan struct{}, 1)
	go func() {
		p.pool.Close()
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func buildPostgresDns(host, port, sslMode, user, password, name string) string {
	return fmt.Sprintf(
		"host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		host, port, sslMode, user, password, name,
	)
}
