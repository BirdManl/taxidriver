// taxi_grpc.pb.go - сгенерированные функции для gRPC
package gen

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// TaxiServiceClient - интерфейс для клиента gRPC
type TaxiServiceClient interface {
	RequestTaxi(ctx context.Context, in *TaxiRequest, opts ...grpc.CallOption) (*TaxiResponse, error)
}

type taxiServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *taxiServiceClient) RequestTaxi(ctx context.Context, in *TaxiRequest, opts ...grpc.CallOption) (*TaxiResponse, error) {
	// Здесь реализуется вызов метода
	return nil, fmt.Errorf("RequestTaxi not implemented")
}

// TaxiServiceServer - интерфейс для сервера gRPC
type TaxiServiceServer interface {
	RequestTaxi(context.Context, *TaxiRequest) (*TaxiResponse, error)
}

// TaxiRequest - структура запроса для метода RequestTaxi
type TaxiRequest struct {
	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

// TaxiResponse - структура ответа для метода RequestTaxi
type TaxiResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

// RegisterTaxiServiceServer - регистрация сервиса на сервере
func RegisterTaxiServiceServer(s *grpc.Server, srv TaxiServiceServer) {
	// здесь регистрируется сервис
}
