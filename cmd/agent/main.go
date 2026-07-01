package main

import (
	"context"
	"log"
	"time"

	"github.com/amir/sys-monitor/internal/agent"
	"github.com/amir/sys-monitor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	log.Println("[Агент] запуск мониторинга")

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC серверу: %v", err)
	}
	defer conn.Close()

	client := proto.NewMetricsServiceClient(conn)
	serverID := "wsl-ubuntu-amir"

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	log.Println("[Агент] Успешно запущен! начинаю сбор информации раз в 5 секунд")
	for range ticker.C {
		req, err := agent.CollectMetrics(serverID)
		if err != nil {
			log.Printf("Ошибка сбора метрик: %v", err)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		res, err := client.SendMetrics(ctx, req)
		cancel()
		if err != nil {
			log.Printf("Ошибка при отправке метриков на сервер по gRPC: %v", err)
		} else {
			log.Printf("Метрики успешно отправлены! Ответ сервера: %s (Succsec: %t)", res.Message, res.Success)
		}
	}
}
