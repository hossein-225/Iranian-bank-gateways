package zarinpal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (z *ZarinPalService) Inquiry(authority string) (*InquiryResponseData, error) {
	inquiryReq := InquiryRequest{
		MerchantID: z.API,
		Authority:  authority,
	}

	jsonData, err := json.Marshal(inquiryReq)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(InquiryURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var inquiryResp InquiryResponse
	if err := json.NewDecoder(resp.Body).Decode(&inquiryResp); err != nil {
		return nil, err
	}

	if inquiryResp.Errors.Code != 0 {
		return nil, GetErrorMessage(inquiryResp.Errors.Code, inquiryResp.Errors.Validations)
	} else if inquiryResp.Data.Code != 100 {
		return nil, GetErrorMessage(inquiryResp.Data.Code, "خطایی رخ داده است")
	}

	return &inquiryResp.Data, nil
}
