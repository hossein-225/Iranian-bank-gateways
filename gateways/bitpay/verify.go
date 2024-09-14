package bitpay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	neturl "net/url"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	bitpayerror "github.com/hossein-225/Iranian-bank-gateways/internal/errors"
)

func (b *BitPayIR) Verify(ctx context.Context, transID, idGet string) (map[string]any, error) {
	data := neturl.Values{}
	data.Set("api", b.API)
	data.Set("trans_id", transID)
	data.Set("id_get", idGet)
	data.Set("json", "1")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.BitPay.VerifyURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.Logger.Error("error occurred: %w", zap.Error(err))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	status, ok := result["status"].(float64)
	if !ok {
		return nil, errors.New("خطا در تأیید: نوع status نادرست است")
	}

	if status != 1 {
		return nil, fmt.Errorf("خطا در تأیید: %s", bitpayerror.GetBitPayVerifyError(status))
	}

	return result, nil
}
