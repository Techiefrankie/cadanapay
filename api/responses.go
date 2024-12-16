package api

type Response struct {
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
}

type ExchangeRateApiResponse struct {
	Result             string  `json:"result"`
	Documentation      string  `json:"documentation"`
	TermsOfUse         string  `json:"terms_of_use"`
	TimeLastUpdateUnix int64   `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string  `json:"time_last_update_utc"`
	TimeNextUpdateUtc  string  `json:"time_next_update_utc"`
	TimeNextUpdateUnix int64   `json:"time_next_update_unix"`
	BaseCode           string  `json:"base_code"`
	TargetCode         string  `json:"target_code"`
	ConversionRate     float64 `json:"conversion_rate"`
}

type OpenExchangeRatesApiResponse struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func GetResponse(description string, code int) Response {
	return Response{
		Description: description,
		StatusCode:  code,
	}
}
