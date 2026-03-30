# Go NLP Text Pipeline

**High-performance concurrent text processing service** written in Go.

Demonstrates production-ready backend skills: clean architecture, concurrency, dual API design and containerization.

### ✨ Features

- ⚡ **Fast concurrent batch processing** using goroutines and worker pools
- 🔌 **Dual API**: REST (Gin) + gRPC with Protocol Buffers
- 📦 **Single & batch** text processing support
- 🐳 **Fully Dockerized** with multi-stage builds and docker-compose
- 📊 **Clean Architecture** with clear layer separation
- 🔍 Supports key NLP tasks: classification, sentiment analysis, summarization, NER
- 📝 Structured logging with zerolog

### 🚀 Quick Start

```bash
# Local run
make run

Using Docker
make docker-up


### API Examples

**Single text processing**
```bash
curl -X POST http://localhost:8080/api/v1/process \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Go is excellent for building scalable NLP services",
    "task": "classify"
  }'

Batch processing

curl -X POST http://localhost:8080/api/v1/batch \
  -H "Content-Type: application/json" \
  -d '{
    "texts": ,
    "task": "sentiment"
  }'

🛠 Tech Stack

Go 1.23 — Goroutines, Context, Clean Architecture
Gin + gRPC + Protocol Buffers
Docker & Docker Compose
Zerolog — Structured logging

📋 Make Commands

make run — start locally
make docker-up — start with Docker
make build — build binary
make test — run tests
