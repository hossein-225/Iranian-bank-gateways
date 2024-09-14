package saman

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/hossein-225/Iranian-bank-gateways/internal/errors"
	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	neturl "net/url"
)

func (ps *PaymentService) Verify(ctx context.Context, refNum string) (*VerifyResponse, error) {
	data := neturl.Values{}
	data.Set("TerminalNumber", ps.TerminalID)
	data.Set("RefNum", refNum)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Saman.VerifyURL, strings.NewReader(data.Encode()))
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

	if code, err := strconv.Atoi(string(body)); err == nil {
		return nil, fmt.Errorf("خطا در تایید: %s", errors.GetSamanVerifyAndReverseError(code))
	}

	var result VerifyResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.ResultCode != 0 {
		return nil, fmt.Errorf("خطا در تایید: %s", errors.GetSamanVerifyAndReverseError(result.ResultCode))
	}

	return &result, nil
}
