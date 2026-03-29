package main

import (
    "context"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/DennisMRitchie/go-nlp-text-pipeline/internal/handler"
    "github.com/DennisMRitchie/go-nlp-text-pipeline/internal/service"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    pb "github.com/DennisMRitchie/go-nlp-text-pipeline/proto/nlp"
)

func main() {
    processor := service.NewProcessor()
    h := handler.NewHandler(processor)

    // REST API
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

    // Graceful shutdown
    go func() {
        log.Println("🚀 Go NLP Text Pipeline started on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe: %s\n", err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("✅ Server exited gracefully")
}
