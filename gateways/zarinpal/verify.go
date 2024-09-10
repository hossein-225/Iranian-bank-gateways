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

func (z *ZarinPalService) Verify(ctx context.Context, authority string, amount int, orderID string) (*VerifyResponseData, error) {
	verifyReq := VerifyRequest{
		MerchantID: z.API,
		Authority:  authority,
		Amount:     amount,
		OrderID:    orderID,
	}

	jsonData, err := json.Marshal(verifyReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal verify request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Zarinpal.VerifyURL, bytes.NewBuffer(jsonData))
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
		if closeErr := resp.Body.Close(); closeErr != nil {
			z.Logger.Error("error closing response body", zap.Error(closeErr))
		}
	}()

	var verifyResp map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&verifyResp); err != nil {
		return nil, fmt.Errorf("failed to decode verify response: %w", err)
	}

	if errorsField, ok := verifyResp["errors"].(map[string]any); ok {
		if code, ok := errorsField["code"].(float64); ok && code != 0 {
			return nil, fmt.Errorf("payment error: %w", zarinpalerror.GetZarinpalError(int(code), errorsField["validations"]))
		}
	}

	if dataField, ok := verifyResp["data"].(map[string]any); ok {
		var verifyData VerifyResponseData
		if cardPan, ok := dataField["card_pan"].(string); ok {
			verifyData.CardPan = cardPan
		}
		if cardHash, ok := dataField["card_hash"].(string); ok {
			verifyData.CardHash = cardHash
		}
		if feeType, ok := dataField["fee_type"].(string); ok {
			verifyData.FeeType = feeType
		}
		if refID, ok := dataField["ref_id"].(float64); ok {
			verifyData.RefID = int(refID)
		}
		if code, ok := dataField["code"].(float64); ok {
			verifyData.Code = int(code)
		}
		if fee, ok := dataField["fee"].(float64); ok {
			verifyData.Fee = int(fee)
		}

		return &verifyData, nil
	}

	return nil, errors.New("unexpected response structure")
}
