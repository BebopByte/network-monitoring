package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DeviceRequest struct {
	IP           string `json:"ip"`
	Hostname     string `json:"hostname"`
	Online       bool   `json:"online"`
	ResponseTime int64  `json:"responseTimeMs"`
	ScannedAt    string `json:"scannedAt"`
}

func sendDevice(device Device, apiUrl string) error {

	payload := DeviceRequest{
		IP:           device.IP,
		Hostname:     device.Hostname,
		Online:       device.Online,
		ResponseTime: device.ResponseTime.Milliseconds(),
		ScannedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	jsonBytes, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	response, err := http.Post(
		apiUrl,
		"application/json",
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Printf(
		"Sent %s to monitor API (status=%d)\n",
		device.IP,
		response.StatusCode,
	)

	return nil
}

func sendDevices(devices []DeviceRequest, apiUrl string) error {

	jsonBytes, err := json.Marshal(devices)

	if err != nil {
		return err
	}

	response, err := http.Post(
		apiUrl,
		"application/json",
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Printf(
		"Sent %d devices to monitor API (status=%d)\n",
		len(devices),
		response.StatusCode,
	)

	return nil

}
