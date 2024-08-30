package bitpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strconv"
)

func (b *BitPayIR) Verify(transID, idGet string) (map[string]interface{}, error) {
	data := neturl.Values{}
	data.Set("api", b.API)
	data.Set("trans_id", transID)
	data.Set("id_get", idGet)
	data.Set("json", "1")

	resp, err := http.PostForm(VerifyURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseStr := string(body)
	if _, err := strconv.Atoi(responseStr); err == nil {
		return nil, fmt.Errorf("خطا در تأیید: %s", HandleErrorCode(responseStr))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
