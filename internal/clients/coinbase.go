package clients

import (
	"context"
	"net/http"
)

type CoinBaseService struct {
	host       string
	HttpClient http.Client
}

func (c *CoinBaseService) GetExchangeRates(ctx context.Context, holding_type string) {

	//TODO add http call to https://api.coinbase.com/v2/exchange-rates?currency=holding_type

}
