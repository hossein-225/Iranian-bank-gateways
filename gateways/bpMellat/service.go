package bpmellat

import (
	"fmt"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

type BpMellat struct {
	Logger       *zap.Logger
	UserName     string `xml:"userName"`
	UserPassword string `xml:"userPassword"`
	TerminalID   int    `xml:"terminalId"`
}

type BpPayRequest struct {
	LocalDate       string `xml:"localDate"`
	LocalTime       string `xml:"localTime"`
	AdditionalData  string `xml:"additionalData"`
	CallBackURL     string `xml:"callBackUrl"`
	OrderID         int64  `xml:"orderId"`
	SaleOrderID     int64  `xml:"saleOrderId"`
	SaleReferenceID int64  `xml:"saleReferenceId"`
	Amount          int64  `xml:"amount"`
	PayerID         int64  `xml:"payerId"`
}

type BpRequest struct {
	OrderID         int64 `xml:"orderId"`
	SaleOrderID     int64 `xml:"saleOrderId"`
	SaleReferenceID int64 `xml:"saleReferenceId"`
}

func NewService(terminalID int, userName, userPassword string) (*BpMellat, error) {
	if err := config.LoadConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return &BpMellat{
		TerminalID:   terminalID,
		UserName:     userName,
		UserPassword: userPassword,
		Logger:       logger,
	}, nil
}
