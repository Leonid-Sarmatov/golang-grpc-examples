package main

import (
	proto "weather_service/client/internal/grpc_package"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"context"
	"io"
	"fmt"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial("localhost:40001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewWeatherServiceClient(conn)

	// Отправляем запрос и получаем поток
	stream, err := client.GetWeatherUpdates(context.Background(), &proto.Request{
		CityName: "New York",
	})
	if err != nil {
		log.Fatalf("Error while calling GetWeatherUpdates: %v", err)
	}

	log.Println("Connected to server, waiting for weather updates...")

	// Читаем поток данных
	for {
		// Получаем данные от сервера
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server has closed the connection.")
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving data: %v", err)
		}

		// Выводим данные на экран
		fmt.Printf("City: %s, Temperature: %.2f°C, Precipitation: %.2f%%, Wind Speed: %.2f m/s, Condition: %s\n",
			response.CityName,
			response.Temperature,
			response.Precipitation,
			response.WindSpeed,
			response.Condition,
		)
	}
}