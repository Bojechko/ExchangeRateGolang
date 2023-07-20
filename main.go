package main

import (
	"ExchangeRateGolang/service"
	"flag"
	"log"
)

func main() {
	code := flag.String("code", "USD", "Код валюты")
	date := flag.String("date", "20/07/2023", "Дата")
	flag.Parse()

	apiRequester := service.NewApiRequester()

	currencyCode, err := apiRequester.GetCurrencyCode("https://www.cbr.ru/scripts/XML_val.asp?d=0", *code)
	if err != nil {
		log.Fatalln(err)
	}

	path := "https://www.cbr.ru/scripts/XML_dynamic.asp?date_req1=" + *date + "&date_req2=" + *date + "&VAL_NM_RQ=" + currencyCode

	exchangeRate, err := apiRequester.GetExchangeRate(path)
	if err != nil {
		log.Fatalln(err)
	}

	answer := *code + ":" + exchangeRate
	println(answer)
}
