package bitpay

import (
	"context"
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

func (b *BitPayIR) Request(ctx context.Context, request *BitPayRequest) (string, error) {
	data := neturl.Values{}
	data.Set("api", b.API)
	data.Set("redirect", request.CallbackURL)
	data.Set("amount", strconv.Itoa(request.Amount))
	data.Set("factorId", request.OrderID)
	data.Set("name", request.Name)
	data.Set("email", request.Email)
	data.Set("description", request.Description)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.BitPay.RequestURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.Logger.Error("Error closing response body: %w", zap.Error(err))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	responseStr := string(body)
	if i, err := strconv.Atoi(responseStr); i <= 0 || err != nil {
		return "", fmt.Errorf("خطا در درخواست: %s", errors.GetBitPayRequestError(responseStr))
	}

	addr := config.AppConfig.BitPay.PayURL + "gateway-" + responseStr + "-get"

	return addr, nil
}
