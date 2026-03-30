# Go NLP Text Pipeline

**High-performance concurrent text processing pipeline in Go** with built-in support for NLP tasks.

Demonstrates production-ready Go backend skills: concurrency, clean architecture, REST + gRPC, Docker, and CI/CD.

### ✨ Features
- ⚡ Fast concurrent batch processing with goroutines
- 🔌 REST API + ready for gRPC/Protobuf
- 📦 Batch & single text processing
- 🐳 Fully Dockerized with docker-compose
- 📊 Clean Architecture + structured logging
- 🔍 NLP-focused tasks: classification, sentiment, summarization, NER

### 🚀 Quick Start

```bash
# Local run
make run

# Docker
make docker-up

API Examples
Process single text:

curl -X POST http://localhost:8080/api/v1/process \
  -H "Content-Type: application/json" \
  -d '{"text": "Go is excellent for building scalable NLP services", "task": "classify"}'

  Batch processing:

  curl -X POST http://localhost:8080/api/v1/batch \
  -H "Content-Type: application/json" \
  -d '{"texts": ["Great product!", "Terrible experience"], "task": "sentiment"}'

  Tech Stack

Go 1.23 (Goroutines, Context, Generics)
Gin + gRPC
Docker & Docker Compose
GitHub Actions CI

Commands

make run — start locally
make build — build binary
make docker-up — run with Docker