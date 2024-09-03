package bpMellat

const (
	url        = "https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"
	GatewayURL = "https://bpm.shaparak.ir/pgwchannel/startpay.mellat"
)

type BpMellat struct {
	TerminalId   int    `xml:"terminalId"`
	UserName     string `xml:"userName"`
	UserPassword string `xml:"userPassword"`
}

type BpPayRequest struct {
	OrderId         int64 `xml:"orderId"`
	SaleOrderId     int64 `xml:"saleOrderId"`
	SaleReferenceId int64 `xml:"saleReferenceId"`

	Amount         int64  `xml:"amount"`
	LocalDate      string `xml:"localDate"`
	LocalTime      string `xml:"localTime"`
	AdditionalData string `xml:"additionalData"`
	CallBackUrl    string `xml:"callBackUrl"`
	PayerId        int64  `xml:"payerId"`
}

type CallbackResponse struct {
	RefId           string `json:"RefId"`
	ResCode         string `json:"ResCode"`
	SaleOrderId     int64  `json:"SaleOrderId"`
	SaleReferenceId int64  `json:"SaleReferenceId"`
	CardHolderInfo  string `json:"CardHolderInfo"`
}

func NewService(terminalId int, userName, userPassword string) *BpMellat {

	return &BpMellat{
		TerminalId:   terminalId,
		UserName:     userName,
		UserPassword: userPassword,
	}

}
