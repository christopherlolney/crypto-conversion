package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CoinBaseService struct {
	host       string
	HttpClient *http.Client
}

type CoinbaseExchangeResponse struct {
	Data struct {
		Currency string            `json:"currency"`
		Rates    map[string]string `json:"rates"`
	} `json:"data"`
}

func New() CoinBaseService {
	return CoinBaseService{
		host: "https://api.coinbase.com/v2/exchange-rates",
		HttpClient: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// Gets the exchange rates as a map for a given currency type
// Notes http.Get would handle this, but I wanted to experiment from the client.do as that would be required for patch requests
func (c *CoinBaseService) GetExchangeRates(ctx context.Context, holding_type string) (*CoinbaseExchangeResponse, error) {

	u, err := url.Parse(c.host)
	if err != nil {
		return nil, err
	}

	u.Query().Add("currency", holding_type)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	defer resp.Body.Close()

	//TODO check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed to get successful response from coinbase server please check network connections and try again \n Status: %d", resp.StatusCode))
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var exchangeResponse CoinbaseExchangeResponse
	err = json.Unmarshal(bytes, &exchangeResponse)
	if err != nil {
		return nil, err
	}

	return &exchangeResponse, nil
}
