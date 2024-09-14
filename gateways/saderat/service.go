package saderat

import (
	"fmt"

	config "github.com/hossein-225/Iranian-bank-gateways/configs"
	"go.uber.org/zap"
)

type PaymentRequest struct {
	Amount      int64  `json:"Amount"`
	CallbackURL string `json:"CallbackURL"`
	InvoiceID   string `json:"InvoiceID"`
	TerminalID  int64  `json:"TerminalID"`
	Payload     string `json:"Payload"`
}

type PaymentResponse struct {
	Status      int    `json:"Status"`
	AccessToken string `json:"AccessToken"`
}

type PaymentCallback struct {
	RespCode       int    `json:"respcode"`       // کد نتیجه تراکنش
	RespMsg        string `json:"respmsg"`        // پیام نتیجه تراکنش
	Amount         int64  `json:"amount"`         // مبلغ کسر شده از حساب مشتری
	InvoiceID      string `json:"invoiceid"`      // شماره فاکتور پذیرنده
	Payload        string `json:"Payload"`        // اطلاعات تکمیلی ارسال شده توسط پذیرنده
	TerminalID     int64  `json:"terminalid"`     // شماره ترمینال
	TraceNumber    int64  `json:"tracenumber"`    // شماره پیگیری تراکنش
	RRN            int64  `json:"rrn"`            // شماره سند بانکی
	DatePaid       string `json:"datePaid"`       // زمان و تاریخ تراکنش
	DigitalReceipt string `json:"digitalreceipt"` // رسید دیجیتال تراکنش
	CardNumber     string `json:"cardnumber"`     // شماره کارت پرداخت کننده (mask شده)
	IssuerBank     string `json:"issuerbank"`     // نام بانک صادر کننده کارت
}

type AdviceRequest struct {
	DigitalReceipt string `json:"digitalreceipt"`
	TerminalID     int64  `json:"Tid"`
}

type AdviceResponse struct {
	Status   string `json:"Status"`
	ReturnID int    `json:"ReturnId"`
	Message  string `json:"Message"`
}

type RollbackRequest struct {
	DigitalReceipt string `json:"digitalreceipt"`
	TerminalID     int64  `json:"Tid"`
}

type RollbackResponse struct {
	Status   string `json:"Status"`
	ReturnID string `json:"ReturnId"`
	Message  string `json:"Message"`
}

type PaymentService struct {
	Logger     *zap.Logger
	TerminalID int64
}

func NewPaymentService(terminalId int64) (*PaymentService, error) {

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
		TerminalID: terminalId,
		Logger:     logger,
	}, nil

}
