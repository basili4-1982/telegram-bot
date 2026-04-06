package rate

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"resty.dev/v3"
)

type Rate struct {
	client *resty.Client
}

func NewRate(client *resty.Client) *Rate {
	if client == nil {
		log.Fatalln("client is  nil")
	}
	return &Rate{client: client}
}

func (r Rate) Get(symbol string) (float64, error) {
	uri := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := r.client.R().Get(uri)
	if err != nil {
		return 0, fmt.Errorf("error while get rate: %w", err)
	}

	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("error while decode rate: %w", err)
	}

	rateVal, err := strconv.ParseFloat(result.Price, 64)

	return rateVal, nil
}
