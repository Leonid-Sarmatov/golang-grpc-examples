package server

import (
	"log"
	"net"
	"context"
	proto "currency_converter/server/internal/grpc/grpc_package"
	"google.golang.org/grpc"
)

/* Структура сервера, отвечающего за конвертацию валют */
type Server struct {
	convertor CurrencyConverterService
	proto.CurrencyConverterServer
	grpcServer *grpc.Server
	listener   net.Listener
}

/* Интерфейс для преобразователя валют */
type CurrencyConverterService interface {
	Convert(input, output string, value float64) (float64, error)
}

/* Конструктор сервера */
func NewServer(conv CurrencyConverterService) (*Server, error) {
	// Создание слушателя для порта
	lis, err := net.Listen("tcp", ":40001")
	if err != nil {
		log.Printf("Can not open tcp port %v", err)
		return nil, err
	}

	// Инициализация полей
	server := &Server{
		convertor: conv,
		listener:  lis,
		grpcServer: grpc.NewServer(),
	}

	// Регистрация сервера и передача необходимых структур
	proto.RegisterCurrencyConverterServer(server.grpcServer, server)
	return server, nil
}

/* Запуск сервера */
func (s *Server) Start() error {
	log.Println("Starting gRPC server on :40001")
	return s.grpcServer.Serve(s.listener)
}

/* Остановка сервера */
func (s *Server) Stop() error {
	log.Println("Stopping gRPC server...")
	s.grpcServer.GracefulStop()
	return nil
}

/* 
Метод сервера для преобразования 
запроса grpc в вызов метода сервисного слоя
*/
func (s *Server) ConvertCurrency(ctx context.Context, request *proto.CurrencyRequest) (*proto.CurrencyResponse, error) {
	res, err := s.convertor.Convert(request.FromCurrency, request.ToCurrency, request.Amount)
	if err != nil {
		log.Printf("Can not convert currency %v", err)
		return &proto.CurrencyResponse{
			ConvertedAmount: 0,
		}, err
	}

	return &proto.CurrencyResponse{
		ConvertedAmount: res,
	}, nil
}