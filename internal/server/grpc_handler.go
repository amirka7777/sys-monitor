package server

import (
	"context"
	"log"

	"github.com/amir/sys-monitor/internal/storage"
	"github.com/amir/sys-monitor/proto"
)

// gRPC-сервер который принимает данные
type MetricsServer struct {
	proto.UnimplementedMetricsServiceServer
	store *storage.Storage
}

func NewMetricsServer(store *storage.Storage) *MetricsServer {
	return &MetricsServer{store: store}
}

func (s *MetricsServer) SendMetrics(ctx context.Context, req *proto.MetricRequest) (*proto.MetricResponse, error) {

	log.Printf("[gRPC-сервер] Принят пакет от %s. Записываю метрики в базу данных!\n", req.ServerId)

	// log.Printf("[gRPC-сервер] Получены метрики от сервера: %s", req.ServerId)
	// log.Printf("-> Cpu использовано: %.2f%%", req.CpuUsage)
	// log.Printf("-> Свободное пространство на диске: %d байт", req.FreeDiskSpace)
	// log.Printf("-> Timestamp: %d", req.Timestamp)

	return &proto.MetricResponse{
		Success: true,
		Message: "Метриик успешно приняты сервером!",
	}, nil

}
