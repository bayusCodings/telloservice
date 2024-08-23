package utils

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"NGN": true,
}

func IsSupportedCurrency(currency string) bool {
	_, ok := supportedCurrencies[currency]
	return ok
}
