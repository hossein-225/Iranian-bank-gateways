package zarinpal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	zarinpalerror "github.com/hossein-225/Iranian-bank-gateways/internal/errors"
)

func (z *ZarinPalService) Request(ctx context.Context, request *PaymentRequestDto) (string, error) {
	paymentRequest := PaymentRequest{
		MerchantID:  z.API,
		CallbackURL: request.CallbackURL,
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
		return "", fmt.Errorf("failed to marshal payment request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Zarinpal.RequestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			z.Logger.Error("Error closing response body", zap.Error(closeErr))
		}
	}()

	var paymentResp map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		return "", fmt.Errorf("failed to decode payment response: %w", err)
	}

	if errorsField, ok := paymentResp["errors"].(map[string]any); ok {
		if code, ok := errorsField["code"].(float64); ok && code != 0 {
			return "", fmt.Errorf("payment error: %w", zarinpalerror.GetZarinpalError(int(code), errorsField["validations"]))
		}
	}

	if dataField, ok := paymentResp["data"].(map[string]any); ok {
		if authority, ok := dataField["authority"].(string); ok {
			return config.AppConfig.Zarinpal.PayURL + authority, nil
		}
	}

	return "", errors.New("unexpected response structure")
}
