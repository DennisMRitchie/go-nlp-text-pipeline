# go-nlp-text-pipeline

High-performance concurrent **Go** text processing pipeline with **NLP** focus.

### Features
- ⚡ Concurrent batch processing with goroutines
- 🔌 REST + gRPC support (proto ready)
- 📊 OpenTelemetry ready
- 🛡️ Rate limiting & graceful shutdown
- 📈 Production-grade structured logging

### Quick Start
```bash
make run
# or
docker-compose up --build


Example Request

curl -X POST http://localhost:8080/api/v1/process \
  -H "Content-Type: application/json" \
  -d '{"text": "Go is awesome for building fast NLP services", "task": "classify"}'

Tech Stack
Go 1.23 • Gin • gRPC • Protobuf • Docker • Kubernetes-ready
