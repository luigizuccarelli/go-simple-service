package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/microlib/simple"
)

type Response struct {
	Status     string `json:"status"`
	StatusCode string `json:"statuscode"`
	Message    string `json:"message"`
	//Payload []Data
	Result string `json:"result"`
}

type CompositeResponse struct {
	Payload string `json:"payload"`
}

type Data struct {
	Type      string      `json:"type"`
	Model     string      `json:"model"`
	Mac       string      `json:"mac"`
	DeviceMac string      `json:"devicemac"`
	Rssi      json.Number `json:"rssi"`
	Diastolic json.Number `json:"diastolic"`
	Systolic  json.Number `json:"systolic"`
	Oxygen    json.Number `json:"oxygen"`
	Static    json.Number `json:"staticheartrate"`
	HeartRate json.Number `json:"heartrate"`
	Battery   json.Number `json:"battery"`
	Count     json.Number `json:"count"`
}

func SimpleHandler(w http.ResponseWriter, r *http.Request, logger *simple.Logger) {
	var response *Response
	// data, _ := RedisClient.Get("timeseriesdata").Result()
	body, err := ioutil.ReadAll(r.Body)
	logger.Debug(fmt.Sprintf("Input data %s", string(body)))
	if err != nil {
		logger.Error(fmt.Sprintf("could not read body data %v", err))
		response = &Response{StatusCode: "500", Status: "KO", Message: fmt.Sprintf("Could not read body data %v", err), Result: ""}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		//json.Unmarshal([]byte(body), &result)
		//ilogger.Debug(fmt.Sprintf("Payload data %s", result.Payload))
		response = &Response{Status: "OK", StatusCode: "200", Message: "Request Successfulli", Result: string(body)}
		w.WriteHeader(http.StatusOK)
	}
	b, _ := json.MarshalIndent(response, "", "	")
	fmt.Fprintf(w, string(b))
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok version 1.0")
}
