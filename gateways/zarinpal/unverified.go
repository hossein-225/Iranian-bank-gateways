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

func (z *ZarinPalService) Unverified(ctx context.Context) (*UnverifiedResponseData, error) {
	req := struct {
		MerchantID string `json:"merchant_id"`
	}{
		MerchantID: z.API,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	reqPost, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Zarinpal.UnverifiedURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	reqPost.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqPost)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			z.Logger.Error("error closing response body", zap.Error(closeErr))
		}
	}()

	var unverifiedResp map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&unverifiedResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if errorsField, ok := unverifiedResp["errors"].(map[string]any); ok {
		if code, ok := errorsField["code"].(float64); ok && code != 0 {
			return nil, fmt.Errorf("payment error: %w", zarinpalerror.GetZarinpalError(int(code), errorsField["validations"]))
		}
	}

	if dataField, ok := unverifiedResp["data"].(map[string]any); ok {
		var unverifiedData UnverifiedResponseData
		if authorities, ok := dataField["authorities"].([]any); ok {
			for _, authority := range authorities {
				authorityMap, ok := authority.(map[string]any)
				if !ok {
					continue
				}

				auth := struct {
					Authority   string `json:"authority"`
					CallbackURL string `json:"callback_url"`
					Referer     string `json:"referer"`
					Date        string `json:"date"`
					Amount      int    `json:"amount"`
				}{}

				if authAuthority, ok := authorityMap["authority"].(string); ok {
					auth.Authority = authAuthority
				}
				if callbackURL, ok := authorityMap["callback_url"].(string); ok {
					auth.CallbackURL = callbackURL
				}
				if referer, ok := authorityMap["referer"].(string); ok {
					auth.Referer = referer
				}
				if date, ok := authorityMap["date"].(string); ok {
					auth.Date = date
				}
				if amount, ok := authorityMap["amount"].(float64); ok {
					auth.Amount = int(amount)
				}

				unverifiedData.Authorities = append(unverifiedData.Authorities, auth)
			}

			return &unverifiedData, nil
		}
	}

	return nil, errors.New("unexpected response structure")
}
