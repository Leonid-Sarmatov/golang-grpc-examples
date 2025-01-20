package server

import (
	"time"
	proto "weather_service/server/internal/grpc_package"

	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.WeatherServiceServer
	weatherGetter GetWeatherService
	grpcServer *grpc.Server
	listener   net.Listener
}

/* Интерфейс для системы определения погоды */
type GetWeatherService interface {
	GetWeatherInfo(sity string) (Weather, error)
}

/* Интерфейс для самой погоды */
type Weather interface {
	GetSityName() string
	GetTemperature() float32
	GetPercipitation() float32
	GetWindSpeed() float32
	GetCondition() string
}

/* Конструктор */
func NewServer(gw GetWeatherService) (*Server, error) {
	// Создание слушателя для порта
	lis, err := net.Listen("tcp", ":40001")
	if err != nil {
		log.Printf("Can not open tcp port %v", err)
		return nil, err
	}

	// Инициализация полей
	server := &Server{
		weatherGetter: gw,
		listener: lis,
		grpcServer: grpc.NewServer(),
	}

	proto.RegisterWeatherServiceServer(server.grpcServer, server)
	return server, nil
}

/* Запуск сервера */
func (s *Server) Start() error {
	log.Println("Starting gRPC server on :40001")
	return s.grpcServer.Serve(s.listener)
}

/* остановка сервера */
func (s *Server) Stop() error {
	log.Println("Stopping gRPC server...")
	s.grpcServer.GracefulStop()
	return nil
}

func (s *Server) GetWeatherUpdates(request *proto.Request, stream grpc.ServerStreamingServer[proto.Response]) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			weather, err := s.weatherGetter.GetWeatherInfo(request.CityName)
			if err != nil {
				log.Printf("Can not get weather info: %v\n", err)
				continue
			}

			err = stream.Send(&proto.Response{
				CityName: weather.GetSityName(),
				Temperature: weather.GetTemperature(),
				Precipitation: weather.GetPercipitation(),
				WindSpeed: weather.GetWindSpeed(),
				Condition: weather.GetCondition(),
			})

			if err != nil {
				log.Printf("Send message was failed: %v\n", err)
				return err
			}

			time.Sleep(1 * time.Second)
		}
	}
}
