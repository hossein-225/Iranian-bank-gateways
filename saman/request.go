package saman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (r *PaymentService) SendRequest(amount int, resNum, cellNumber, redirectUrl string) (string, error) {

	paymentRequest := &PaymentRequest{
		Action:      "token",
		TerminalId:  r.TerminalId,
		Amount:      amount,
		ResNum:      resNum,
		RedirectUrl: redirectUrl,
		CellNumber:  cellNumber,
	}

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", RequestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("خطایی رخ داده است : %s", GetErrorMessage(resp.StatusCode))
	}

	var responseMap map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", fmt.Errorf("خطایی رخ داده است : %s", GetErrorMessage(resp.StatusCode))
	}

	token, ok := responseMap["token"].(string)
	if !ok {
		return "", fmt.Errorf(GetErrorMessage(10))
	}

	paymentURL := fmt.Sprintf("%s?token=%s", PayURL, token)

	return paymentURL, nil
}
