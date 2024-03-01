package handlers_test

import (
	"context"
	"testing"

	"github.com/christopherlolney/crypto-conversion/internal/clients/mocks"
	"github.com/christopherlolney/crypto-conversion/internal/handlers"
)

type testCase struct {
	Input  input
	output string
}

type input struct {
	description    string
	holdings       float64
	currencyType70 string
	currencyType30 string
}

func TestHandleConversion(T *testing.T) {
	ctx := context.Background()
	coinbaseClient := mocks.CoinBaseServiceMock{}

	T.Log("Testing Handle Conversion")
	for _, testCase := range testHandleConversionTestCases() {
		T.Log(testCase.Input.description)
		errString := ""
		err := handlers.HandleConversion(ctx, &coinbaseClient, testCase.Input.holdings, testCase.Input.currencyType70, testCase.Input.currencyType30)
		if err != nil {
			errString = err.Error()
		}
		if errString != testCase.output {
			T.Errorf(`Test Failed to succeed
			case: %s
			holdings: %.2f
			currencyType70: %s
			currencyType30: %s
			expectedOutput: %s
			actualOutput: %s
			`,
				testCase.Input.description,
				testCase.Input.holdings,
				testCase.Input.currencyType70,
				testCase.Input.currencyType30,
				testCase.output,
				errString,
			)
		}

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
			output: "",
		}, {
			Input: input{
				description:    "Testing negative holdings",
				holdings:       -100.00,
				currencyType70: "BTC",
				currencyType30: "ETH",
			},
			output: "holdings needs to be a positive value",
		}, {
			Input: input{
				description:    "Testing invalid currency type",
				holdings:       100.00,
				currencyType70: "DERF",
				currencyType30: "ETH",
			},
			output: "Invalid currency type: DERF",
		}, {
			Input: input{
				description:    "Testing 0 holdings",
				holdings:       0,
				currencyType70: "BTC",
				currencyType30: "ETH",
			},
			output: "holdings needs to be a positive value",
		},
	}
}
