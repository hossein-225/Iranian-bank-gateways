package bpMellat

const (
	url         = "https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"
	GatewayURL  = "https://bpm.shaparak.ir/pgwchannel/startpay.mellat"
	CallBackUrl = "your callback URL"
)

type BpMellat struct {
	TerminalId      int    `xml:"terminalId"`
	UserName        string `xml:"userName"`
	UserPassword    string `xml:"userPassword"`
	OrderId         int64  `xml:"orderId"`
	SaleOrderId     int64  `xml:"saleOrderId"`
	SaleReferenceId int64  `xml:"saleReferenceId"`
}

type BpPayRequest struct {
	TerminalId   int    `xml:"terminalId"`
	UserName     string `xml:"userName"`
	UserPassword string `xml:"userPassword"`
	OrderId      int64  `xml:"orderId"`

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

func NewService(orderId, saleOrderId, saleReferenceId int64) *BpMellat {

	// complete it and it's better to read it from env/config file
	terminalId := 0
	userName := ""
	userPassword := ""

	return &BpMellat{
		TerminalId:      terminalId,
		UserName:        userName,
		UserPassword:    userPassword,
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}
}
