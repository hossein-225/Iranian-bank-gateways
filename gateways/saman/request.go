package saman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hossein-225/Iranian-bank-gateways/internal/errors"
	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

func (ps *PaymentService) SendRequest(ctx context.Context, amount int, resNum, cellNumber, redirectURL string) (string, error) {
	paymentRequest := &PaymentRequest{
		Action:      "token",
		TerminalID:  ps.TerminalID,
		Amount:      amount,
		ResNum:      resNum,
		RedirectURL: redirectURL,
		CellNumber:  cellNumber,
	}

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", fmt.Errorf("خطا در تبدیل درخواست به JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Saman.RequestURL, bytes.NewBuffer(jsonData))
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
			ps.Logger.Error("Error closing response body:", zap.Error(err))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("خطا: %s", errors.GetSamanError(resp.StatusCode))
	}

	var responseMap map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", fmt.Errorf("خطا در پردازش پاسخ: %w", err)
	}

	token, ok := responseMap["token"].(string)
	if !ok {
		return "", fmt.Errorf("خطا: %s", errors.GetSamanError(10))
	}

	paymentURL := fmt.Sprintf("%s?token=%s", config.AppConfig.Saman.PayURL, token)

	return paymentURL, nil
}