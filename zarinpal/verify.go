package zarinpal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (z *ZarinPalService) Verify(authority string, amount int, orderID string) (*VerifyResponseData, error) {
	verifyReq := VerifyRequest{
		MerchantID: z.API,
		Authority:  authority,
		Amount:     amount,
		OrderID:    orderID,
	}

	jsonData, err := json.Marshal(verifyReq)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(VerifyURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var verifyResp VerifyResponse
	json.NewDecoder(resp.Body).Decode(&verifyResp)

	if verifyResp.Errors.Code != 0 {
		return nil, GetErrorMessage(verifyResp.Errors.Code, verifyResp.Errors.Validations)
	} else if verifyResp.Data.Code != 100 {
		return nil, GetErrorMessage(verifyResp.Data.Code, "خطایی رخ داده است")
	}

	return &verifyResp.Data, nil
}
