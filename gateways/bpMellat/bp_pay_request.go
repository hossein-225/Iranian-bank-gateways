package bpmellat

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	bpmellaterror "github.com/hossein-225/Iranian-bank-gateways/internal/errors"
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SoapBody `xml:"Body"`
}

type SoapBody struct {
	BpPayRequestResponse BpPayRequestResponse `xml:"bpPayRequestResponse"`
}

type BpPayRequestResponse struct {
	Return string `xml:"return"`
}

func (req *BpMellat) BpPayRequest(ctx context.Context, input *BpPayRequest) (string, error) {
	now := time.Now()
	sendDate := now.Format("20060102")
	sendTime := now.Format("150405")

	if input.LocalDate == "" {
		input.LocalDate = sendDate
	}

	if input.LocalTime == "" {
		input.LocalTime = sendTime
	}

	soapEnvelope := `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <web:bpPayRequest>
         <terminalId>` + strconv.Itoa(req.TerminalID) + `</terminalId>
         <userName>` + req.UserName + `</userName>
         <userPassword>` + req.UserPassword + `</userPassword>
         <orderId>` + strconv.FormatInt(input.OrderID, 10) + `</orderId>
         <amount>` + strconv.FormatInt(input.Amount, 10) + `</amount>
         <localDate>` + input.LocalDate + `</localDate>
         <localTime>` + input.LocalTime + `</localTime>
         <additionalData>` + input.AdditionalData + `</additionalData>
         <callBackUrl>` + input.CallBackURL + `</callBackUrl>
         <payerId>` + strconv.FormatInt(input.PayerID, 10) + `</payerId>
      </web:bpPayRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Mellat.URL, bytes.NewBufferString(soapEnvelope))
	if err != nil {
		return "", fmt.Errorf("خطا در ساخت درخواست HTTP: %w", err)
	}

	httpReq.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست HTTP: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			req.Logger.Error("Error closing response body", zap.Error(cerr))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("خطا در خواندن پاسخ: %w", err)
	}

	var response SoapEnvelope
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("خطا در تجزیه XML: %w", err)
	}

	parts := strings.Split(response.Body.BpPayRequestResponse.Return, ",")
	if len(parts) < 2 {
		return "", errors.New("پاسخ نامعتبر از سرور بانک")
	}

	if bpmellaterror.GetBPMellatError(parts[0]) != nil {
		return "", fmt.Errorf("%w", bpmellaterror.GetBPMellatError(parts[0]))
	}

	paymentURL := fmt.Sprintf("%s?RefId=%s", config.AppConfig.Mellat.GatewayURL, parts[1])

	return paymentURL, nil
}
