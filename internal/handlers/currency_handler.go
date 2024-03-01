package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/christopherlolney/crypto-conversion/internal/clients"
)

// handler function for converting holdings into 2 groups 70% of holdings and 30% of holdings
// after splitting it converts those groups from their current value in USD to 2 selected crypto values.
func HandleConversion(ctx context.Context, service clients.CoinBaseService, holdings float64, currencyType70 string, currencyType30 string) {
	logger := log.Default()

	//TODO research floating point math accuracy
	holdings70 := holdings * .7

	holdings30 := holdings - holdings70

	coinBaseResp, err := service.GetExchangeRates(ctx, "USD")
	if err != nil {
		logger.Fatalln(err)
	}

	// Validate Currency types
	rate70 := coinBaseResp.Data.Rates[currencyType70]
	if rate70 == "" {
		logger.Fatalf("Invalid currency type: %s", currencyType70)
	}
	rate30 := coinBaseResp.Data.Rates[currencyType30]
	if rate70 == "" {
		logger.Fatalf("Invalid currency type: %s", currencyType30)
	}

	converted70, err := convertCurrency(rate70, holdings70)
	if err != nil {
		logger.Fatalln(err)
	}

	converted30, err := convertCurrency(rate30, holdings30)
	if err != nil {
		logger.Fatalln(err)
	}

	//Output
	logger.Printf("$%.2f => %.4f %s\n", holdings70, converted70, currencyType70)
	logger.Printf("$%.2f => %.4f %s\n", holdings30, converted30, currencyType30)

}

// Split conversion into it's own method for readability
func convertCurrency(rate string, holdings float64) (float64, error) {
	rateFloat, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		return 0, err
	}

	converted := holdings * rateFloat

	return converted, nil
}
