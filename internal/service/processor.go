package service

import (
    "context"
    "sync"
    "time"

    "github.com/DennisMRitchie/go-nlp-text-pipeline/internal/model"
    "github.com/DennisMRitchie/go-nlp-text-pipeline/proto/nlp"
)

type Processor struct{}

func NewProcessor() *Processor {
    return &Processor{}
}

func (p *Processor) Process(ctx context.Context, req *model.TextRequest) (*nlp.ProcessResponse, error) {
    start := time.Now()

    // Simulate heavy NLP work with concurrency-friendly design
    // In real project you would call Python ML service via gRPC or HTTP
    result := simulateNLP(req.Text, req.Task)

    return &nlp.ProcessResponse{
        Result:     result,
        Confidence: 0.87,
        Metadata: map[string]string{
            "processed_in": time.Since(start).String(),
            "language":     "en",
        },
    }, nil
}

func (p *Processor) BatchProcess(ctx context.Context, req *model.BatchRequest) (*nlp.BatchResponse, error) {
    var wg sync.WaitGroup
    results := make([]*nlp.ProcessResponse, len(req.Texts))
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

    return &nlp.BatchResponse{Results: results}, nil
}

// Simulate different NLP tasks
func simulateNLP(text string, task string) string {
    switch task {
    case "classify":
        return "Technology"
    case "sentiment":
        return "positive"
    case "summarize":
        return "This text discusses advancements in natural language processing using Go for high-throughput services."
    case "ner":
        return "Entities: [Go, NLP, Kubernetes]"
    default:
        return "processed"
    }
}
