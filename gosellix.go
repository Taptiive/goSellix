package goSellix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SellixClient struct {
	Client  http.Client
	Request http.Request
	AuthKey string
	Method  string
	Url     string
}

func NewClient(authKey string) SellixClient {
	return SellixClient{
		Client:  *http.DefaultClient,
		AuthKey: authKey,
	}
}

func (r SellixClient) BlacklistByID(blacklistID string) Blacklist {
	r.Method = "GET"
	r.Url = "https://dev.sellix.io/v1/blacklist/" + blacklistID
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	r.Request = *req
	res, err := r.Client.Do(&r.Request)
	if err != nil {
		return Blacklist{}
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Blacklist{}
	}
	var SellixBlacklist Blacklist
	err = json.Unmarshal(buf, &SellixBlacklist)
	if err != nil {
		return Blacklist{}
	}
	return SellixBlacklist
}

func (r SellixClient) OrderByID(orderUniqID string) Order {
	r.Method = "GET"
	r.Url = "https://dev.sellix.io/v1/orders/" + orderUniqID
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	r.Request = *req
	res, err := r.Client.Do(&r.Request)
	if err != nil {
		return Order{}
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Order{}
	}
	var SellixOrder Order
	err = json.Unmarshal(buf, &SellixOrder)
	if err != nil {
		return Order{}
	}
	return SellixOrder
}
