package saderat

import (
	"fmt"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

//nolint:tagliatelle // bank need to use the below structure
type PaymentRequest struct {
	CallbackURL string `json:"CallbackURL"`
	InvoiceID   string `json:"InvoiceID"`
	Payload     string `json:"Payload"`
	Amount      int64  `json:"Amount"`
	TerminalID  int64  `json:"TerminalID"`
}

//nolint:tagliatelle // bank need to use the below structure
type PaymentResponse struct {
	AccessToken string `json:"AccessToken"`
	Status      int    `json:"Status"`
}

//nolint:tagliatelle // bank need to use the below structure
type PaymentCallback struct {
	RespMsg        string `json:"respmsg"`
	InvoiceID      string `json:"invoiceid"`
	Payload        string `json:"Payload"`
	DatePaid       string `json:"datePaid"`
	DigitalReceipt string `json:"digitalreceipt"`
	CardNumber     string `json:"cardnumber"`
	IssuerBank     string `json:"issuerbank"`
	RespCode       int    `json:"respcode"`
	Amount         int64  `json:"amount"`
	TerminalID     int64  `json:"terminalid"`
	TraceNumber    int64  `json:"tracenumber"`
	RRN            int64  `json:"rrn"`
}

//nolint:tagliatelle // bank need to use the below structure
type AdviceRequest struct {
	DigitalReceipt string `json:"digitalreceipt"`
	TerminalID     int64  `json:"Tid"`
}

//nolint:tagliatelle // bank need to use the below structure
type AdviceResponse struct {
	Status   string `json:"Status"`
	Message  string `json:"Message"`
	ReturnID int    `json:"ReturnId"`
}

//nolint:tagliatelle // bank need to use the below structure
type RollbackRequest struct {
	DigitalReceipt string `json:"digitalreceipt"`
	TerminalID     int64  `json:"Tid"`
}

//nolint:tagliatelle // bank need to use the below structure
type RollbackResponse struct {
	Status   string `json:"Status"`
	ReturnID string `json:"ReturnId"`
	Message  string `json:"Message"`
}

type PaymentService struct {
	Logger     *zap.Logger
	TerminalID int64
}

func NewPaymentService(terminalID int64) (*PaymentService, error) {
	err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Error syncing logger: %w", zap.Error(err))
		}
	}()

	return &PaymentService{
		TerminalID: terminalID,
		Logger:     logger,
	}, nil
}
