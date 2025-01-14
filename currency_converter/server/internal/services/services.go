package services

/*
Для описания валюты создана структура
с именем и с коэфициентов перевода в доллары США

Было бы эффективнее хранить экземпляры в
хэш таблице или же в массиве, но в данном решении
применяется что то похожее на итератор или компоновщик,
где валюты связанны
*/

type Currency struct {
	CurrencyName string
	OneCurrensyToOneUSD float64
	NextCurrency *Currency
}

func (c *Currency) Next() *Currency {
	return c.NextCurrency
}

func (c *Currency) Add(next *Currency) {
	c.NextCurrency = next
}


type CurrencyConverter struct {
	StartCurrency *Currency
}


func (cc *CurrencyConverter) Convert(input, output string, value float64) (float64, error) {
	// Инициализация
	now := cc.StartCurrency

	// Перебираем валюты, пока не встретим первую
	for now := cc.StartCurrency; now.CurrencyName != input; now = now.Next() {}

	// Переводим исходную валюту в доллары
	res := value * now.OneCurrensyToOneUSD

	// Перебираем валюты, пока не встретим вторую
	for now := cc.StartCurrency; now.CurrencyName != output; now = now.Next() {}

	// Переводим доллары в целевую валюту
	return res / now.OneCurrensyToOneUSD, nil
}