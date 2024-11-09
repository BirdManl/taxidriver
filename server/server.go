package main

import (
	"context"
	"fmt"
	"log"
	"my-taxi-app/gen"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	gen.UnimplementedTaxiServiceServer
}

func (s *server) RequestRide(ctx context.Context, req *gen.RideRequest) (*gen.RideResponse, error) {
	// Логика для обработки запроса
	return &gen.RideResponse{
		DriverName:  "Иван Иванов",
		VehicleInfo: "Toyota Prius, Белый",
	}, nil
}

func main() {
	// Настройка gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось слушать на порту 50051: %v", err)
	}

	s := grpc.NewServer()
	gen.RegisterTaxiServiceServer(s, &server{})

	// Запуск сервера
	fmt.Println("Сервер слушает на порту 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
