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

type SoapVerifyEnvelope struct {
	XMLName xml.Name       `xml:"Envelope"`
	Body    SoapVerifyBody `xml:"Body"`
}

type SoapVerifyBody struct {
	BpVerifyRequestResponse BpVerifyRequestResponse `xml:"bpVerifyRequestResponse"`
}

type BpVerifyRequestResponse struct {
	Return string `xml:"return"`
}

func (req *BpMellat) BpVerifyRequest() error {
	soapEnvelope := `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <web:bpVerifyRequest>
         <terminalId>` + strconv.Itoa(req.TerminalId) + `</terminalId>
         <userName>` + req.UserName + `</userName>
         <userPassword>` + req.UserPassword + `</userPassword>
         <orderId>` + strconv.FormatInt(req.OrderId, 10) + `</orderId>
         <saleOrderId>` + strconv.FormatInt(req.SaleOrderId, 10) + `</saleOrderId>
         <saleReferenceId>` + strconv.FormatInt(req.SaleReferenceId, 10) + `</saleReferenceId>
      </web:bpVerifyRequest>
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

	var response SoapVerifyEnvelope
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	parts := strings.Split(response.Body.BpVerifyRequestResponse.Return, ",")
	if len(parts) < 1 {
		return errors.New("پاسخ نامعتبر از سرور بانک")
	}

	if err = getBankError(parts[0]); err != nil {
		return err
	}

	return nil
}
