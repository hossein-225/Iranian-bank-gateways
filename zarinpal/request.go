package zarinpal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (z *ZarinPalService) Request(request PaymentRequestDto) (string, error) {
	paymentRequest := PaymentRequest{
		MerchantID:  z.API,
		CallbackURL: z.Redirect,
		Amount:      request.Amount,
		Description: request.Description,
		Currency:    request.Currency,
		Metadata: map[string]string{
			"email":    request.Email,
			"mobile":   request.Mobile,
			"order_id": request.OrderID,
		},
	}

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(PaymentURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var paymentResp PaymentResponse
	json.NewDecoder(resp.Body).Decode(&paymentResp)

	if paymentResp.Errors.Code != 0 {
		return "", GetErrorMessage(paymentResp.Errors.Code, paymentResp.Errors.Validations)
	} else if paymentResp.Data.Code != 100 {
		validations := map[string][]string{"validations": {"خطایی رخ داده است"}}
		return "", GetErrorMessage(paymentResp.Data.Code, validations)
	}

	return PayURL + paymentResp.Data.Authority, nil
}
