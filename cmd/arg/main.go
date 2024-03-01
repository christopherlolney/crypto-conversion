package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/christopherlolney/crypto-conversion/internal/clients"
	"github.com/christopherlolney/crypto-conversion/internal/handlers"
)

func main() {

	ctx := context.Background()

	logger := log.Default()
	args := os.Args
	if len(args) < 4 {
		logger.Fatalln("Function main requires 3 arguments: (holdings validCurrency validCurrency)")
	}
	holdings, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		logger.Fatalln(err)
	}

	currencyType70 := args[2]

	currencyType30 := args[3]

	coinbaseClient := clients.New()

	err = handlers.HandleConversion(ctx, &coinbaseClient, holdings, currencyType70, currencyType30)
	if err != nil {
		logger.Fatalln(err)
	}
}
