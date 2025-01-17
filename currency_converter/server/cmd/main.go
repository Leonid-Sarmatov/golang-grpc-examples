package main

import (
	server "currency_converter/server/internal/server"
	services "currency_converter/server/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	usd := &services.Currency{
		CurrencyName: "USD",
		OneCurrensyToOneUSD: 1.0,
	}

	eur := &services.Currency{
		CurrencyName: "EUR",
		OneCurrensyToOneUSD: 0.92,
	}

	gbp := &services.Currency{
		CurrencyName: "GBP",
		OneCurrensyToOneUSD: 0.27,
	}

	a := &services.CurrencyConverter{
		CurrencyMap: make(map[string]*services.Currency),
	}

	a.Add(eur)
	a.Add(gbp)
	a.Add(usd)

	s, err := server.NewServer(a)
	if err != nil {
		log.Fatalf("Can not create new server: %v\n", err)
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