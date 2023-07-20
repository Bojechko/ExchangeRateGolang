package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ApiRequester struct {
}

func NewApiRequester() *ApiRequester {
	return &ApiRequester{}
}

func (r *ApiRequester) GetCurrencyCode(path string, code string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	cur := Currency{}
	currencyCode, err := cur.UnmarshalCurrency(string(body), code)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return currencyCode, nil
}

func (r *ApiRequester) GetData(path string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	response := Response{}
	responceExchangeRate, err := response.UnmarshalResponse(string(body))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return responceExchangeRate, nil
}
