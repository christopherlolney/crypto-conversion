package handlers_test

import (
	"context"
	"testing"

	"github.com/christopherlolney/crypto-conversion/internal/clients"
	"github.com/christopherlolney/crypto-conversion/internal/handlers"
)

type testCase struct {
	Input  input
	output struct{}
}

type input struct {
	description    string
	holdings       float64
	currencyType70 string
	currencyType30 string
}

func TestHandleConversion(T *testing.T) {
	ctx := context.Background()
	//TODO Mock
	coinbaseClient := clients.New()

	for _, testCase := range testHandleConversionTestCases() {
		T.Log(testCase.Input.description)
		handlers.HandleConversion(ctx, coinbaseClient, testCase.Input.holdings, testCase.Input.currencyType70, testCase.Input.currencyType30)
	}

}

func testHandleConversionTestCases() []testCase {
	return []testCase{
		{
			Input: input{
				description:    "Testing valid conversion",
				holdings:       100.00,
				currencyType70: "BTC",
				currencyType30: "ETH",
			},
			output: struct{}{},
		},
	}
}
