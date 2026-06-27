package server

import (
	"context"
	"log"

	"github.com/amir/sys-monitor/proto"
)

// gRPC-сервер который принимает данные
type MetricsServer struct {
	proto.UnimplementedMetricsServiceServer
}

func NewMetricsServer() *MetricsServer {
	return &MetricsServer{}
}

func (s *MetricsServer) SendMetrics(ctx context.Context, req *proto.MetricRequest) (*proto.MetricResponse, error) {

	log.Printf("[gRPC-сервер] Получены метрики от сервера: %s", req.ServerId)
	log.Printf("-> Cpu использовано: %.2f%%", req.CpuUsage)
	log.Printf("-> Свободное пространство на диске: %d байт", req.FreeDiskSpace)
	log.Printf("-> Timestamp: %d", req.Timestamp)

	return &proto.MetricResponse{
		Success: true,
		Message: "Метриик успешно приняты сервером!",
	}, nil

}
