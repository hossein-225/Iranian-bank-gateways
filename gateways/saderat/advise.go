package saderat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	neturl "net/url"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	"github.com/hossein-225/Iranian-bank-gateways/internal/errors"
	"go.uber.org/zap"
)

func (ps *PaymentService) ConfirmTransaction(ctx context.Context, digitalReceipt string, amount int64) (*AdviceResponse, error) {
	data := neturl.Values{}
	data.Set("Tid", strconv.FormatInt(ps.TerminalID, 10))
	data.Set("digitalreceipt", digitalReceipt)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Saderat.AdviseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			ps.Logger.Error("failed to close response body", zap.Error(closeErr))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result AdviceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if int64(result.ReturnID) != amount {
		return nil, fmt.Errorf("خطا در تایید: %s", errors.HandleServiceErrors(result.ReturnID))
	}

	return &result, nil
}
