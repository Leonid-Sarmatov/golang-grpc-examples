package main

import (
	"log"
	"context"
	"fmt"
	proto "currency_converter/client/internal/grpc/grpc_package"
	"google.golang.org/grpc"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial("localhost:40001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewCurrencyConverterClient(conn)

	// Отправляем запрос к серверу
	response, err := client.ConvertCurrency(context.Background(), &proto.CurrencyRequest{
		Amount:       100.0,
		FromCurrency: "EUR",
		ToCurrency:   "USD",
	})
	if err != nil {
		log.Printf("Error converting currency: %v", err)
	}
	fmt.Printf("Converted amount: %.2f %s\n", response.ConvertedAmount, "USD")
}