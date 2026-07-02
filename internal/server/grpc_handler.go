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

	err := s.store.SaveMetrics(req.ServerId, req.CpuUsage, req.FreeDiskSpace, req.Timestamp)
	if err != nil {
		log.Printf("[База данных] не удалось сохранить метрики: %v\n", err)
		return &proto.MetricResponse{
			Success: false,
			Message: "Ошибка при сохранении на сервере",
		}, nil
	}

	return &proto.MetricResponse{
		Success: true,
		Message: "Метриик успешно приняты сервером!",
	}, nil

}
