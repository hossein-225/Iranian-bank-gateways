package saman

import (
	"fmt"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

//nolint:tagliatelle // bank need to use PascalCase
type PaymentRequest struct {
	TerminalID       string `json:"TerminalId"`
	ResNum           string `json:"ResNum"`
	RedirectURL      string `json:"RedirectUrl"`
	CellNumber       string `json:"CellNumber"`
	Action           string `json:"action"`
	HashedCardNumber string `json:"HashedCardNumber,omitempty"`
	TokenExpiryInMin int    `json:"TokenExpiryInMin,omitempty"`
	Amount           int    `json:"Amount"`
}

//nolint:tagliatelle // bank need to use PascalCase
type CallbackResponse struct {
	MID              string `json:"MID"`
	State            string `json:"State"`
	Status           string `json:"Status"`
	RRN              string `json:"RRN"`
	RefNum           string `json:"RefNum"`
	ResNum           string `json:"ResNum"`
	TerminalID       string `json:"TerminalId"`
	TraceNo          string `json:"TraceNo"`
	SecurePan        string `json:"SecurePan"`
	HashedCardNumber string `json:"HashedCardNumber"`
	Amount           int    `json:"Amount"`
	Wage             int    `json:"Wage,omitempty"`
}

//nolint:tagliatelle // bank need to use PascalCase
type VerifyRequest struct {
	RefNum         string `json:"RefNum"`
	TerminalNumber int64  `json:"TerminalNumber"`
}

//nolint:tagliatelle // bank need to use PascalCase
type VerifyResponse struct {
	ResultDescription string `json:"ResultDescription"`
	TransactionDetail struct {
		RRN             string `json:"RRN"`
		RefNum          string `json:"RefNum"`
		MaskedPan       string `json:"MaskedPan"`
		HashedPan       string `json:"HashedPan"`
		StraceDate      string `json:"StraceDate"`
		StraceNo        string `json:"StraceNo"`
		TerminalNumber  int32  `json:"TerminalNumber"`
		OriginalAmount  int64  `json:"OrginalAmount"`
		AffectiveAmount int64  `json:"AffectiveAmount"`
	} `json:"TransactionDetail"`
	Success    bool `json:"Success"`
	ResultCode int  `json:"ResultCode"`
}

type PaymentService struct {
	Logger     *zap.Logger
	TerminalID string
}

func NewPaymentService(terminalID string) (*PaymentService, error) {
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
