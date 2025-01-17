package services

import (
	"fmt"
	"log"
)

type Currency struct {
	CurrencyName string
	OneCurrensyToOneUSD float64
}


type CurrencyConverter struct {
	CurrencyMap  map[string]*Currency
}


func (cc *CurrencyConverter) Add(c *Currency) {
	cc.CurrencyMap[c.CurrencyName] = c
}


func (cc *CurrencyConverter) Convert(input, output string, value float64) (float64, error) {
	k1, ok := cc.CurrencyMap[input]
	if !ok {
		return 0, fmt.Errorf("This currency not found: %v", input)
	}

	res := value * k1.OneCurrensyToOneUSD

	k2, ok := cc.CurrencyMap[output]
	if !ok {
		return 0, fmt.Errorf("This currency not found: %v", input)
	}

	res = res / k2.OneCurrensyToOneUSD

	log.Printf("%v %v = %v %v", value, input, res, output)
	
	return res, nil
}