package saman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strconv"
)

func (s *PaymentService) Verify(refNum string) (map[string]interface{}, error) {

	data := neturl.Values{}
	data.Set("terminalId", s.TerminalId)
	data.Set("refNum", refNum)

	resp, err := http.PostForm(VerifyURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if code, err := strconv.Atoi(string(body)); err == nil {
		return nil, fmt.Errorf("خطا در تایید: %s", GetVerifyAndReverseErrorMessage(code))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// Verification result
	if resultCode, ok := result["ResultCode"].(float64); ok && resultCode != 0 {
		return nil, fmt.Errorf("خطا در تایید: %s", GetVerifyAndReverseErrorMessage(int(resultCode)))
	}

	return result, nil
}
