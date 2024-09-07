package zarinpal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (z *ZarinPalService) Unverified() (*UnverifiedResponseData, error) {
	req := struct {
		MerchantID string `json:"merchant_id"`
	}{
		MerchantID: z.API,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(UnverifiedURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var unverifiedResp UnverifiedResponse
	if err := json.NewDecoder(resp.Body).Decode(&unverifiedResp); err != nil {
		return nil, err
	}

	if unverifiedResp.Errors.Code != 0 {
		return nil, GetErrorMessage(unverifiedResp.Errors.Code, unverifiedResp.Errors.Validations)
	} else if unverifiedResp.Data.Authorities == nil {
		return nil, errors.New("error occurred")
	}

	return &unverifiedResp.Data, nil
}
