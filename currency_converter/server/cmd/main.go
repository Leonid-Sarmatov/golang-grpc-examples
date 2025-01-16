package main

import (
	server "currency_converter/server/internal/server"
	services "currency_converter/server/internal/services"
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

	usd.Add(eur)
	eur.Add(gbp)
	gbp.Add(usd)

	a := &services.CurrencyConverter{
		StartCurrency: usd,
	}

	server.NewServer(a)

	for {
		
	}
}