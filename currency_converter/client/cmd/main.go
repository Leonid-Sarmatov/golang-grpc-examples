package main

import (
	"context"
	proto "currency_converter/client/internal/grpc/grpc_package"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial("localhost:40001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Printf("Connection - OK\n")

	client := proto.NewCurrencyConverterClient(conn)

	// Отправляем запрос к серверу
	response, err := client.ConvertCurrency(context.Background(), &proto.CurrencyRequest{
		Amount:       100.0,
		FromCurrency: "EUR",
		ToCurrency:   "USD",
	})
	if err != nil {
		log.Printf("Error converting currency: %v", err)
	} else {
		fmt.Printf("Converted amount: %.2f %s\n", response.ConvertedAmount, "USD")
	}
}
