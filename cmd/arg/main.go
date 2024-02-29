package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// ctx := context.Background()

	logger := log.Default()
	args := os.Args
	if len(args) < 4 {
		logger.Fatalln("Function main requires 3 arguments a numeric value for holdings and 2 valid crypto currency types for conversion")
	}
	holdings, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		logger.Fatalln("Could not convert holdings to type float")
	}

	currencyType70 := args[2]

	currencyType30 := args[3]

	logger.Println(fmt.Sprintf("Holdings : %.2f, Currency to convert 70 percent of holdings: %s, Currency to Convert 30 percent of holdings: %s", holdings, currencyType70, currencyType30))

}
