package service

import (
	"context"
	"sync"
	"time"

	"github.com/DennisMRitchie/go-nlp-text-pipeline/internal/model"
	pb "github.com/DennisMRitchie/go-nlp-text-pipeline/proto/nlp"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

// Process обрабатывает один текст (симуляция NLP-задачи)
func (p *Processor) Process(ctx context.Context, req *model.TextRequest) (*pb.ProcessResponse, error) {
	start := time.Now()

	// Здесь в реальном проекте был бы вызов внешнего ML-сервиса (Python + HuggingFace / vLLM и т.д.)
	result := simulateNLP(req.Text, req.Task)

	return &pb.ProcessResponse{
		Result:     result,
		Confidence: 0.89,
		Metadata: map[string]string{
			"processed_in":    time.Since(start).String(),
			"language":        "en",
			"task":            req.Task,
			"model_simulated": "go-nlp-pipeline-v1",
		},
	}, nil
}

// BatchProcess обрабатывает несколько текстов параллельно
func (p *Processor) BatchProcess(ctx context.Context, req *model.BatchRequest) (*pb.BatchResponse, error) {
	var wg sync.WaitGroup
	results := make([]*pb.ProcessResponse, len(req.Texts))
	mu := sync.Mutex{}

	for i, text := range req.Texts {
		wg.Add(1)
		go func(idx int, t string) {
			defer wg.Done()

			resp, _ := p.Process(ctx, &model.TextRequest{Text: t, Task: req.Task})

			mu.Lock()
			results[idx] = resp
			mu.Unlock()
		}(i, text)
	}

	wg.Wait()

	return &pb.BatchResponse{Results: results}, nil
}

// simulateNLP — симуляция разных NLP задач
func simulateNLP(text string, task string) string {
	switch task {
	case "classify":
		return "Technology & AI"
	case "sentiment":
		return "positive"
	case "summarize":
		return "Go — отличный язык для создания высокопроизводительных NLP-сервисов с отличной поддержкой concurrency."
	case "ner":
		return "ORG: xAI, TECH: Go, NLP, Kubernetes"
	default:
		return "text_processed"
	}
}
