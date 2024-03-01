package mocks

import (
	"context"
	"encoding/json"

	_ "embed"

	"github.com/christopherlolney/crypto-conversion/internal/clients"
)

type CoinBaseServiceMock struct {
}

//go:embed responses/example_response.json
var mockCoinbaseResponse []byte

func (c *CoinBaseServiceMock) GetExchangeRates(ctx context.Context, holding_type string) (*clients.CoinbaseExchangeResponse, error) {

	var input clients.CoinbaseExchangeResponse

	err := json.Unmarshal(mockCoinbaseResponse, &input)

	return &input, err
}
