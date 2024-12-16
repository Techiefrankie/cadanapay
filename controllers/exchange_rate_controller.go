package controllers

import (
	"cadanapay/api"
	"cadanapay/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetExchangeRate(context *gin.Context) {
	persons, er := service.GetPersonsInstance()
	if er != nil {
		fmt.Println("Error while getting person's instance: ", er)
	}

	fmt.Println("Original Persons: ", persons)
	fmt.Println("Sorted Persons Asc: ", persons.Sort(true))
	fmt.Println("Sorted Persons Desc: ", persons.Sort(false))
	fmt.Println("Filter Persons by USD: ", persons.FilterByCurrency("USD"))
	fmt.Println("Filter Persons by EUR: ", persons.FilterByCurrency("EUR"))
	fmt.Println("Filter Persons by NGN: ", persons.FilterByCurrency("NGN"))
	fmt.Println("Group Persons by Currencies: ", persons.GroupBySalaryCurrency())

	var rateRequest api.RateRequest
	err := context.BindJSON(&rateRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.GetResponse("Invalid request", http.StatusBadRequest))
		return
	}

	currencyPair := rateRequest.CurrencyPair
	pair := strings.Split(currencyPair, "-")

	if len(pair) != 2 {
		msg := fmt.Sprintf("Invalid currency pair: %s", currencyPair)
		context.JSON(http.StatusBadRequest, api.GetResponse(msg, http.StatusBadRequest))
		return
	}

	baseCurrency := pair[0]
	quoteCurrency := pair[1]

	rate, err := service.GetFirstAvailableRate(baseCurrency, quoteCurrency)

	if err != nil {
		context.JSON(http.StatusInternalServerError, api.GetResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	fmt.Println("Exchange rate for ", currencyPair, " is ", rate)
	result := make(map[string]float64)
	result[rateRequest.CurrencyPair] = rate

	context.JSON(http.StatusOK, result)
}
