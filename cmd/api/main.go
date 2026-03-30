package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DennisMRitchie/go-nlp-text-pipeline/internal/handler"
	"github.com/DennisMRitchie/go-nlp-text-pipeline/internal/service"
	"github.com/DennisMRitchie/go-nlp-text-pipeline/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	log := logger.Get()

	log.Info().Msg("🚀 Starting Go NLP Text Pipeline v1.0")

	processor := service.NewProcessor()
	h := handler.NewHandler(processor)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/process", h.ProcessText)
		v1.POST("/batch", h.BatchProcess)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Info().Str("port", "8080").Msg("HTTP server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Forced shutdown")
	}

	log.Info().Msg("✅ Server stopped gracefully")
}
