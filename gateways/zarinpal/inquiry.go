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

func (z *ZarinPalService) Inquiry(ctx context.Context, authority string) (*InquiryResponseData, error) {
	inquiryReq := InquiryRequest{
		MerchantID: z.API,
		Authority:  authority,
	}

	jsonData, err := json.Marshal(inquiryReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal inquiry request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Zarinpal.InquiryURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			z.Logger.Error("Error closing response body", zap.Error(err))
		}
	}()

	var inquiryResp map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&inquiryResp); err != nil {
		return nil, fmt.Errorf("failed to decode inquiry response: %w", err)
	}

	if errorsField, ok := inquiryResp["errors"].(map[string]any); ok {
		if code, ok := errorsField["code"].(float64); ok && code != 0 {
			return nil, fmt.Errorf("payment error: %w", zarinpalerror.GetZarinpalError(int(code), errorsField["validations"]))
		}
	}

	if dataField, ok := inquiryResp["data"].(map[string]any); ok {
		var inquiryData InquiryResponseData
		if status, ok := dataField["status"].(string); ok {
			inquiryData.Status = status
		}
		if message, ok := dataField["message"].(string); ok {
			inquiryData.Message = message
		}
		if code, ok := dataField["code"].(float64); ok {
			inquiryData.Code = int(code)
		}

		if inquiryData.Code != 100 {
			return nil, fmt.Errorf("inquiry failed: %s", inquiryData.Message)
		}

		return &inquiryData, nil
	}

	return nil, errors.New("unexpected response structure")
}
