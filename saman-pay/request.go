package samanpay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (r *PaymentRequest) SendRequest(username, password string) (*http.Response, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", RequestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	client := &http.Client{}
	return client.Do(req)
}
