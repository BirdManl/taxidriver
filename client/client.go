package main

import (
	"context"
	"fmt"
	"log"
	"my-taxi-app/gen"

	"google.golang.org/grpc"
)

func main() {
	// Подключение к серверу
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := gen.NewTaxiServiceClient(conn)

	// Отправка запроса
	req := &gen.RideRequest{
		PickupLocation: "ул. Пушкина, 1",
		Destination:    "ул. Лермонтова, 10",
	}

	resp, err := client.RequestRide(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при запросе поездки: %v", err)
	}

	// Вывод ответа
	fmt.Printf("Ваш водитель: %s, автомобиль: %s\n", resp.DriverName, resp.VehicleInfo)
}
