package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type NinoxCall struct {
	baseUrl    string
	apiKey     string
	teamID     string
	databaseID string
	tableID    string
}

type CallConfig func(*NinoxCall)


func NewNinoxCall(opts ...CallConfig) *NinoxCall {
	ninox := NinoxCall{
		baseUrl: "https://api.ninoxdb.de/v1",
	}
	for _, opt := range opts {
		opt(&ninox)
	}
	return &ninox
}

func WithApiKey(apiKey string) CallConfig {
	return func(nx *NinoxCall) {
		nx.apiKey = apiKey
	}
}

func WithTeamAndDatabase(teamID, databaseID string) CallConfig {
	return func(nx *NinoxCall) {
		nx.teamID = teamID
		nx.databaseID = databaseID
	}
}

func (nx *NinoxCall) Records(tableID string, data interface{}) {
	nx.tableID = tableID

	url, err := nx.buildUrl()
	authHeader, err := nx.getAuthHeader()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
}

func (nx *NinoxCall) Create(tableID string, data interface{}) {
	nx.tableID = tableID

	url, err := nx.buildUrl()
	authHeader, err := nx.getAuthHeader()
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respbody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(respbody))

}

func (nx *NinoxCall) buildUrl() (string, error) {
	if nx.tableID == "" {
		return "", errors.New("No table ID given!")
	}
	return fmt.Sprintf("%s/teams/%s/databases/%s/tables/%s/records", nx.baseUrl, nx.teamID, nx.databaseID, nx.tableID), nil
}

func (nx *NinoxCall) getAuthHeader() (string, error) {
	if nx.apiKey == "" {
		return "", errors.New("No API key given!")
	}
	return fmt.Sprintf("Bearer %s", nx.apiKey), nil
}
