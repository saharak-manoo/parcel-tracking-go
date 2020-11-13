package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ETracking struct {}

type Payload struct {
	Courier           string     `json:"courier"`
	TrackingNumber    string     `json:"trackingNumber"`
}

type Response struct {
	Meta              Meta       `json:"meta"`
	Data              Data       `json:"data"`
}

type Meta struct {
	Code              int        `json:"code"`
	Message           string     `json:"message"`
}

type Data struct {
	Courier           string     `json:"courier"`
	CurrentStatus     string     `json:"currentStatus"`
	Detail            Detail     `json:"detail"`
	Timelines         []Timeline `json:"timelines"`
}

type Detail struct {
	Sender            string     `json:"sender"`
	Recipient         string     `json:"recipient"`
}

type Timeline struct {
	Date     		 string          `json:"date"`
	Details      []Event         `json:"details"`
}

type Event struct {
	DateTime     time.Time       `json:"dateTime"`
	Date     		 string          `json:"date"`
	Time     		 string          `json:"time"`
	Description  string          `json:"description"`
}

type Events []Event

func (e *ETracking) Hostname() string {
	return "https://api.etrackings.com"
}

func (e *ETracking) Track(courier string, trackingNumber string) (data Data, err error) {
	url := fmt.Sprintf("%s/v2/tracks/find", e.Hostname())
	log.Println("Call api url -> ", url)
  method := "POST"

  payload := Payload{
		Courier: courier,
		TrackingNumber: trackingNumber,
	}

	payloadByte, _ := json.Marshal(payload)

  client := &http.Client {}
  req, err := http.NewRequest(method, url, bytes.NewReader(payloadByte))

  if err != nil {
    log.Println(err)
    return data, err
	}

  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("etracking-api-key", os.Getenv("ETRACKINGS_API_KEY"))
  req.Header.Add("etracking-key-secret", os.Getenv("ETRACKINGS_KEY_SECRET"))

  res, err := client.Do(req)
  if err != nil {
    log.Println(err)
    return data, err
	}

  defer res.Body.Close()

  bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	response := Response{}
	json.Unmarshal([]byte(bodyString), &response)

	if response.Meta.Code != 200 {
		return data, errors.New(response.Meta.Message)
	}

	return response.Data, nil
}