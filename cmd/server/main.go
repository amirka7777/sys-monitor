package main

import (
	"log"
	"net"

	"github.com/amir/sys-monitor/internal/server"
	"github.com/amir/sys-monitor/proto"
	"google.golang.org/grpc"
)

func main() {

	log.Println("[Сервер] запущено хранилище метрик")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось открыть сервер на порту 50051: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	metricsHandler := server.NewMetricsServer()

	proto.RegisterMetricsServiceServer(grpcServer, metricsHandler)

	log.Println("[Сервер] Успешно запущен и слушает порт 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при работе gRPC сервера: %v", err)
	}

}
