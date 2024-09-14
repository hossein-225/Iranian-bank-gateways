package saderat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	"github.com/hossein-225/Iranian-bank-gateways/internal/errors"
	"go.uber.org/zap"
)

func (s *PaymentService) SendRequest(ctx context.Context, amount int64, callbackURL, invoiceID, payload string) (string, error) {

	paymentRequest := PaymentRequest{
		TerminalID:  s.TerminalID,
		Amount:      amount,
		CallbackURL: callbackURL,
		InvoiceID:   invoiceID,
		Payload:     payload,
	}

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", fmt.Errorf("خطا در تبدیل درخواست به JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Saderat.RequestURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return "", fmt.Errorf("خطا در ساخت درخواست HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست HTTP: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			s.Logger.Error("Error closing response body:", zap.Error(err))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("خطا: %s", errors.HandleServiceErrors(resp.StatusCode))
	}

	var responseMap map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", fmt.Errorf("خطا در پردازش پاسخ: %w", err)
	}

	accessToken, ok := responseMap["AccessToken"].(string)

	if !ok {
		return "", fmt.Errorf("خطا: %s", errors.HandleServiceErrors(10))
	}

	return accessToken, nil

}
