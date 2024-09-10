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

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	bpmellaterror "github.com/hossein-225/Iranian-bank-gateways/internal/errors"
)

type SoapSettleEnvelope struct {
	XMLName xml.Name       `xml:"Envelope"`
	Body    SoapSettleBody `xml:"Body"`
}

type SoapSettleBody struct {
	BpSettleRequestResponse BpSettleRequestResponse `xml:"bpSettleRequestResponse"`
}

type BpSettleRequestResponse struct {
	Return string `xml:"return"`
}

func (req *BpMellat) BpSettleRequest(ctx context.Context, input BpRequest) error {
	soapEnvelope := `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <web:bpSettleRequest>
         <terminalId>` + strconv.Itoa(req.TerminalID) + `</terminalId>
         <userName>` + req.UserName + `</userName>
         <userPassword>` + req.UserPassword + `</userPassword>
         <orderId>` + strconv.FormatInt(input.OrderID, 10) + `</orderId>
         <saleOrderId>` + strconv.FormatInt(input.SaleOrderID, 10) + `</saleOrderId>
         <saleReferenceId>` + strconv.FormatInt(input.SaleReferenceID, 10) + `</saleReferenceId>
      </web:bpSettleRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppConfig.Mellat.URL, bytes.NewBufferString(soapEnvelope))
	if err != nil {
		return fmt.Errorf("خطا در ساخت درخواست HTTP: %w", err)
	}

	httpReq.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست HTTP: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			req.Logger.Error("Error closing response body", zap.Error(cerr))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("خطا در خواندن پاسخ: %w", err)
	}

	var response SoapSettleEnvelope
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("خطا در تجزیه XML: %w", err)
	}

	parts := strings.Split(response.Body.BpSettleRequestResponse.Return, ",")
	if len(parts) < 1 {
		return errors.New("پاسخ نامعتبر از سرور بانک")
	} else if parts[0] != "0" {
		return fmt.Errorf("%w", bpmellaterror.GetBPMellatError(parts[0]))
	}

	return nil
}
