package service

import (
	"cadanapay/api"
	"cadanapay/util"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func GetFirstAvailableRate(baseCurrency, quoteCurrency string) (float64, error) {
	log.Printf("Getting first available rate for the %s-%s pair", baseCurrency, quoteCurrency)

	// Create a channel to receive exchange rates
	rateChan := make(chan float64, 2)
	errChan := make(chan error, 2)

	// Wait group to ensure both goroutines complete
	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch rate from the first API
	go func() {
		defer wg.Done()
		rateResponse, err := GetExchangeRateApiResponse(baseCurrency, quoteCurrency)

		if err != nil {
			errChan <- err
			return
		}

		rateChan <- rateResponse.ConversionRate
	}()

	// Fetch rate from the second API
	go func() {
		defer wg.Done()
		rateResponses, err := GetOpenExchangeRatesApiResponse(baseCurrency, quoteCurrency)

		if err != nil {
			errChan <- err
			return
		}

		rateChan <- rateResponses.Rates[quoteCurrency]
	}()

	// Wait for one successful response or both errors
	go func() {
		wg.Wait()
		close(rateChan)
		close(errChan)
	}()

	// Return the first available rate or error
	select {
	case rate := <-rateChan:
		return rate, nil
	case <-errChan:
		return 0, errors.New(fmt.Sprintf("Both services failed to fetch the exchange rate for %s-%s", baseCurrency, quoteCurrency))
	}
}

func GetExchangeRateApiResponse(base, target string) (api.ExchangeRateApiResponse, error) {
	apiKey, err := util.GetApiKey("ExchangeRateApiKey")

	// Build the URL for the API call
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", apiKey, base, target)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return api.ExchangeRateApiResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error occurred")
		}
	}(resp.Body)

	// Read the response body
	var response api.ExchangeRateApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return api.ExchangeRateApiResponse{}, err
	}

	return response, nil
}

func GetOpenExchangeRatesApiResponse(base, target string) (api.OpenExchangeRatesApiResponse, error) {
	apiKey, err := util.GetApiKey("OpenExchangeRatesApiKey")

	// Build the URL for the API call
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", apiKey)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return api.OpenExchangeRatesApiResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error occurred")
		}
	}(resp.Body)

	// Read the response body
	var response api.OpenExchangeRatesApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return api.OpenExchangeRatesApiResponse{}, err
	}

	return response, nil
}
