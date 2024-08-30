package bpMellat

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func (req *BpMellat) BpSettleRequest(input BpPayRequest) error {
	soapEnvelope := `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <web:bpSettleRequest>
         <terminalId>` + strconv.Itoa(req.TerminalId) + `</terminalId>
         <userName>` + req.UserName + `</userName>
         <userPassword>` + req.UserPassword + `</userPassword>
         <orderId>` + strconv.FormatInt(input.OrderId, 10) + `</orderId>
         <saleOrderId>` + strconv.FormatInt(input.SaleOrderId, 10) + `</saleOrderId>
         <saleReferenceId>` + strconv.FormatInt(input.SaleReferenceId, 10) + `</saleReferenceId>
      </web:bpSettleRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	resp, err := http.Post(url, "text/xml", bytes.NewBuffer([]byte(soapEnvelope)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response SoapSettleEnvelope
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	parts := strings.Split(response.Body.BpSettleRequestResponse.Return, ",")
	if len(parts) < 1 {
		return errors.New("پاسخ نامعتبر از سرور بانک")
	}

	if err = getBankError(parts[0]); err != nil {
		return err
	}

	return nil
}
