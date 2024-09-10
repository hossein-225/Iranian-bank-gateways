package zarinpal

import (
	"fmt"

	"go.uber.org/zap"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
)

type ZarinPalService struct {
	Logger *zap.Logger
	API    string
}

type PaymentRequest struct {
	Metadata    map[string]string `json:"metadata"`
	MerchantID  string            `json:"merchant_id"`
	Description string            `json:"description"`
	CallbackURL string            `json:"callback_url"`
	Currency    string            `json:"currency"`
	Amount      int               `json:"amount"`
}

type PaymentRequestDto struct {
	Description string `json:"description"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Currency    string `json:"currency"`
	OrderID     string `json:"order_id"`
	CallbackURL string `json:"callback_url"`
	Amount      int    `json:"amount"`
}

type PaymentResponseData struct {
	Authority string `json:"authority"`
	CardHash  string `json:"card_hash"`
	CardPan   string `json:"card_pan"`
	Message   string `json:"message"`
	FeeType   string `json:"fee_type"`
	Code      int    `json:"code"`
	RefID     int    `json:"ref_id"`
	Fee       int    `json:"fee"`
}

type ErrorResponse struct {
	Validations map[string][]string `json:"validations"`
	Message     string              `json:"message"`
	Code        int                 `json:"code"`
}

type PaymentResponse struct {
	Errors ErrorResponse       `json:"errors"`
	Data   PaymentResponseData `json:"data"`
}

type VerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	OrderID    string `json:"order_id"`
	Amount     int    `json:"amount"`
}

type VerifyResponseData struct {
	CardPan  string `json:"card_pan"`
	CardHash string `json:"card_hash"`
	FeeType  string `json:"fee_type"`
	RefID    int    `json:"ref_id"`
	Code     int    `json:"code"`
	Fee      int    `json:"fee"`
}

type VerifyResponse struct {
	Errors ErrorResponse      `json:"errors"`
	Data   VerifyResponseData `json:"data"`
}

type InquiryRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
}

type InquiryResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type InquiryResponse struct {
	Errors ErrorResponse       `json:"errors"`
	Data   InquiryResponseData `json:"data"`
}

type UnverifiedResponseData struct {
	Authorities []struct {
		Authority   string `json:"authority"`
		CallbackURL string `json:"callback_url"`
		Referer     string `json:"referer"`
		Date        string `json:"date"`
		Amount      int    `json:"amount"`
	} `json:"authorities"`
}

type UnverifiedResponse struct {
	Errors ErrorResponse          `json:"errors"`
	Data   UnverifiedResponseData `json:"data"`
}

func NewService(token string) (*ZarinPalService, error) {
	err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return &ZarinPalService{
		API:    token,
		Logger: logger,
	}, nil
}
