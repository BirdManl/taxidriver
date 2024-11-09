// taxi.pb.go - сгенерированная структура для сообщений, определенных в taxi.proto
package gen

import (
	"github.com/golang/protobuf/proto"
)

// Taxi - структура сообщения, описывающая такси
type Taxi struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

// Reset сбрасывает структуру Taxi
func (x *Taxi) Reset() {
	*x = Taxi{}
}

// String возвращает строковое представление Taxi
func (x *Taxi) String() string {
	return proto.CompactTextString(x)
}

// ProtoMessage необходимо для реализации интерфейса protobuf
func (*Taxi) ProtoMessage() {}
