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

type SoapReversalEnvelope struct {
	XMLName xml.Name         `xml:"Envelope"`
	Body    SoapReversalBody `xml:"Body"`
}

type SoapReversalBody struct {
	BpReversalRequestResponse BpReversalRequestResponse `xml:"bpReversalRequestResponse"`
}

type BpReversalRequestResponse struct {
	Return string `xml:"return"`
}

func (req *BpMellat) BpReversalRequest() error {
	soapEnvelope := `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <web:bpReversalRequest>
         <terminalId>` + strconv.Itoa(req.TerminalId) + `</terminalId>
         <userName>` + req.UserName + `</userName>
         <userPassword>` + req.UserPassword + `</userPassword>
         <orderId>` + strconv.FormatInt(req.OrderId, 10) + `</orderId>
         <saleOrderId>` + strconv.FormatInt(req.SaleOrderId, 10) + `</saleOrderId>
         <saleReferenceId>` + strconv.FormatInt(req.SaleReferenceId, 10) + `</saleReferenceId>
      </web:bpReversalRequest>
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

	var response SoapReversalEnvelope
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	parts := strings.Split(response.Body.BpReversalRequestResponse.Return, ",")
	if len(parts) < 1 {
		return errors.New("پاسخ نامعتبر از سرور بانک")
	}

	if err = getBankError(parts[0]); err != nil {
		return err
	}

	return nil
}
