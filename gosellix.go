package goSellix

import (
	"bytes"
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

func (r SellixClient) getBody() []byte {
	res, err := r.Client.Do(&r.Request)
	if err != nil {
		return []byte("error making request")
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("error making request")
	}
	return buf
}
func NewClient(authKey string) SellixClient {
	return SellixClient{
		Client:  *http.DefaultClient,
		AuthKey: authKey,
	}
}

func (r SellixClient) DeleteBlacklist(blacklistUniqid string) BlacklistCreation {
	r.Method = "DELETE"
	r.Url = "https://dev.sellix.io/v1/blacklists/" + blacklistUniqid
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	r.Request = *req
	buf := r.getBody()
	var SellixBlacklist BlacklistCreation
	err := json.Unmarshal(buf, &SellixBlacklist)
	if err != nil {
		return BlacklistCreation{}
	}
	return BlacklistCreation{}
}

func (r SellixClient) EditBlacklist(blacklistUniqid, blacklistType, blacklistData, blacklistMessage string) BlacklistCreation {
	r.Method = "PUT"
	r.Url = "https://dev.sellix.io/v1/blacklists/" + blacklistUniqid
	payload := map[string]interface{}{"type": blacklistType, "data": blacklistData, "note": blacklistMessage}
	byts, err := json.Marshal(payload)
	if err != nil {
		return BlacklistCreation{}
	}
	req, _ := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(byts))
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	req.Header.Set("Content-Type", "application/json")
	r.Request = *req
	buf := r.getBody()
	var SellixBlacklist BlacklistCreation
	err = json.Unmarshal(buf, &SellixBlacklist)
	if err != nil {
		return BlacklistCreation{}
	}
	return BlacklistCreation{}
}

func (r SellixClient) CreateBlacklist(blacklistType, blacklistData, blacklistMessage string) BlacklistCreation {
	r.Method = "POST"
	r.Url = "https://dev.sellix.io/v1/blacklists"
	payload := map[string]interface{}{"type": blacklistType, "data": blacklistData, "note": blacklistMessage}
	byts, err := json.Marshal(payload)
	if err != nil {
		return BlacklistCreation{}
	}
	req, _ := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(byts))
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	req.Header.Set("Content-Type", "application/json")
	r.Request = *req
	buf := r.getBody()
	var SellixBlacklist BlacklistCreation
	err = json.Unmarshal(buf, &SellixBlacklist)
	if err != nil {
		return BlacklistCreation{}
	}
	return BlacklistCreation{}
}

func (r SellixClient) BlacklistByID(blacklistID string) BlacklistSpec {
	r.Method = "GET"
	r.Url = "https://dev.sellix.io/v1/blacklists/" + blacklistID
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	r.Request = *req
	buf := r.getBody()
	var SellixBlacklist BlacklistSpec
	err := json.Unmarshal(buf, &SellixBlacklist)
	if err != nil {
		return BlacklistSpec{}
	}
	return SellixBlacklist
}

func (r SellixClient) AllBlacklist() Blacklist {
	r.Method = "GET"
	r.Url = "https://dev.sellix.io/v1/blacklists"
	req, _ := http.NewRequest(r.Method, r.Url, nil)
	req.Header.Add("Authorization", "Bearer "+r.AuthKey)
	r.Request = *req
	buf := r.getBody()
	var SellixBlacklist Blacklist
	err := json.Unmarshal(buf, &SellixBlacklist)
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
