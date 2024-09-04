package saman

const (
	RequestURL = "https://sep.shaparak.ir/OnlinePG/OnlinePG"
	VerifyURL  = "https://sep.shaparak.ir/verifyTxnRandomSessionkey/ipg/VerifyTransaction"
	PayURL     = "https://sep.shaparak.ir/OnlinePG/SendToken"
	ReverseURL = "https://sep.shaparak.ir/verifyTxnRandomSessionkey/ipg/ReverseTransaction"
)

type PaymentRequest struct {
	Action           string `json:"action"`
	TerminalId       string `json:"TerminalId"`
	Amount           int    `json:"Amount"`
	ResNum           string `json:"ResNum"`
	RedirectUrl      string `json:"RedirectUrl"`
	CellNumber       string `json:"CellNumber"`
	TokenExpiryInMin int    `json:"TokenExpiryInMin,omitempty"`
	HashedCardNumber string `json:"HashedCardNumber,omitempty"`
}

type CallbackResponse struct {
	MID              string `json:"MID"`
	State            string `json:"State"`
	Status           string `json:"Status"`
	RRN              string `json:"RRN"`
	RefNum           string `json:"RefNum"`
	ResNum           string `json:"ResNum"`
	TerminalId       string `json:"TerminalId"`
	TraceNo          string `json:"TraceNo"`
	Amount           int    `json:"Amount"`
	Wage             int    `json:"Wage,omitempty"`
	SecurePan        string `json:"SecurePan"`
	HashedCardNumber string `json:"HashedCardNumber"`
}

type VerifyRequest struct {
	RefNum         string `json:"RefNum"`
	TerminalNumber int64  `json:"TerminalNumber"`
}

type VerifyResponse struct {
	TransactionDetail struct {
		RRN             string `json:"RRN"`
		RefNum          string `json:"RefNum"`
		MaskedPan       string `json:"MaskedPan"`
		HashedPan       string `json:"HashedPan"`
		TerminalNumber  int32  `json:"TerminalNumber"`
		OriginalAmount  int64  `json:"OrginalAmount"`
		AffectiveAmount int64  `json:"AffectiveAmount"`
		StraceDate      string `json:"StraceDate"`
		StraceNo        string `json:"StraceNo"`
	} `json:"TransactionDetail"`
	ResultCode        int    `json:"ResultCode"`
	ResultDescription string `json:"ResultDescription"`
	Success           bool   `json:"Success"`
}

type PaymentService struct {
	TerminalId string
	Username   string
	Password   string
}

func NewPaymentService(TerminalId, username, password string) *PaymentService {
	return &PaymentService{
		TerminalId: TerminalId,
		Username:   username,
		Password:   password,
	}
}

func (s *PaymentService) CreateRequest(amount int, resNum, cellNumber, redirectUrl string) *PaymentRequest {
	return &PaymentRequest{
		Action:      "token",
		TerminalId:  s.TerminalId,
		Amount:      amount,
		ResNum:      resNum,
		RedirectUrl: redirectUrl,
		CellNumber:  cellNumber,
	}
}
