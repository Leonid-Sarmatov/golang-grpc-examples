package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	server "weather_service/server/internal/server"
	service "weather_service/server/internal/service"
)

func main() {
	gw, err := service.NewGetWeather()
	if err != nil {
		log.Fatalf("Can not create get weather info service: %v\n", err)
	}

	s, err := server.NewServer(gw)
	if err != nil {
		log.Fatalf("Can not create grpc server: %v\n", err)
	}

	go func() {
		err = s.Start()
		if err != nil {
			log.Fatalf("Start: %v\n", err)
		}
	}()

	// Создаем канал для приема сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала
	<-stop
	log.Println("Shutting down server...")

	// Корректное завершение работы сервера
	if err := s.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
	log.Println("Server stopped gracefully")
}