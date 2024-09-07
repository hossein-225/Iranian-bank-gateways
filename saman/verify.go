package saman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strconv"
)

func (s *PaymentService) Verify(refNum string) (*VerifyResponse, error) {

	data := neturl.Values{}
	data.Set("TerminalNumber", s.TerminalId)
	data.Set("RefNum", refNum)

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

	var result VerifyResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ResultCode != 0 {
		return nil, fmt.Errorf("خطا در تایید: %s", GetVerifyAndReverseErrorMessage(result.ResultCode))
	}

	return &result, nil
}
