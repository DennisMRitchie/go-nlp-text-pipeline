package grpc

import (
	"context"

	"github.com/NatalieDaw92055/go-nlp-text-pipeline/internal/service"
	pb "github.com/NatalieDaw92055/go-nlp-text-pipeline/proto/nlp"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTextProcessorServer
	processor *service.Processor
}

func NewGRPCServer(processor *service.Processor) *Server {
	return &Server{processor: processor}
}

func (s *Server) ProcessText(ctx context.Context, req *pb.ProcessRequest) (*pb.ProcessResponse, error) {
	// Преобразуем в внутреннюю модель
	return s.processor.Process(ctx, &model.TextRequest{
		Text: req.Text,
		Task: req.Task,
	})
}

func (s *Server) BatchProcess(ctx context.Context, req *pb.BatchRequest) (*pb.BatchResponse, error) {
	return s.processor.BatchProcess(ctx, &model.BatchRequest{
		Texts: req.Texts,
		Task:  req.Task,
	})
}

// Register регистрирует gRPC-сервер
func Register(grpcServer *grpc.Server, processor *service.Processor) {
	pb.RegisterTextProcessorServer(grpcServer, NewGRPCServer(processor))
}
