package service

import (
	"encoding/xml"
	"errors"
	"strings"
)

type Currency struct {
	XMLName  xml.Name `xml:"Valuta"`
	Currency []struct {
		Name       string `xml:"Name"`
		EngName    string `xml:"EngName"`
		Nominal    string `xml:"Nominal"`
		ParentCode string `xml:"ParentCode"`
	} `xml:"Item"`
}

type Response struct {
	XMLName  xml.Name `xml:"ValCurs"`
	Response []struct {
		Nominal string `xml:"Nominal"`
		Value   string `xml:"Value"`
	} `xml:"Record"`
}

func (r *Currency) UnmarshalCurrency(data string, code string) (string, error) {
	currency := new(Currency)
	err := xml.Unmarshal([]byte(data), currency)
	if err != nil {
		return "", err
	}

	for _, curNode := range currency.Currency {
		if curNode.EngName == code {
			return strings.TrimSpace(curNode.ParentCode), nil
		}

	}
	return "", errors.New("no such currency")
}

func (r *Response) UnmarshalResponse(data string) (string, error) {
	response := new(Response)
	err := xml.Unmarshal([]byte(data), response)
	if err != nil {
		return "", err
	}

	for _, respNode := range response.Response {
		return strings.TrimSpace(respNode.Value), nil
	}

	return "", errors.New("error in UnmarshalResponse")
}
