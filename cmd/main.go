package main

import (
	"context"
	"errors"
	"net/http"

	"clean-artchit/internal/apater/repository/pgx"
	api_v1 "clean-artchit/internal/api/v1"
	"clean-artchit/internal/config"
	"clean-artchit/internal/domain/service"
	"clean-artchit/internal/domain/usecase/product"
	"clean-artchit/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	log := logger.InitLogger()
	defer log.Sync()

	ctx := context.Background()
	cfg := config.GetConfig()

	pool, err := pgxpool.New(ctx, cfg.BuildPostgresDns())
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	defer pool.Close()

	mux := http.NewServeMux()
	productRep := pgx.NewProductRepository(pool, log)
	productService := service.NewProductService(productRep, log)
	productUsecase := product.NewProductUsecase(productService, log)
	productHandler := api_v1.NewProductHandler(productUsecase, log)
	productHandler.RegisterRoutes(mux)

	addr := cfg.BuildServerAddress()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Info("Server starting", zap.String("address", addr))
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server failed", zap.Error(err))
	}
}
