package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeviceRequest struct {
	IP           string `json:"ip"`
	Hostname     string `json:"hostname"`
	Online       bool   `json:"online"`
	ResponseTime int64  `json:"responseTimeMs"`
}

func sendDevice(device Device, apiUrl string) error {

	payload := DeviceRequest{
		IP:           device.IP,
		Hostname:     device.Hostname,
		Online:       device.Online,
		ResponseTime: device.ResponseTime.Milliseconds(),
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
