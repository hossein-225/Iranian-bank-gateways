package zarinpal

const (
	PaymentURL    = "https://api.zarinpal.com/pg/v4/payment/request.json"
	VerifyURL     = "https://api.zarinpal.com/pg/v4/payment/verify.json"
	InquiryURL    = "https://api.zarinpal.com/pg/v4/payment/inquiry.json"
	UnverifiedURL = "https://api.zarinpal.com/pg/v4/payment/unVerified.json"
	PayURL        = "https://www.zarinpal.com/pg/StartPay/"
)

type ZarinPalService struct {
	API      string
	Redirect string
}

type PaymentRequest struct {
	MerchantID  string            `json:"merchant_id"`
	Amount      int               `json:"amount"`
	Description string            `json:"description"`
	CallbackURL string            `json:"callback_url"`
	Currency    string            `json:"currency"`
	Metadata    map[string]string `json:"metadata"`
}

type PaymentRequestDto struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Currency    string `json:"currency"`
	OrderID     string `json:"order_id"`
}

type PaymentResponseData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Authority string `json:"authority"`
	CardHash  string `json:"card_hash"`
	CardPan   string `json:"card_pan"`
	RefID     int    `json:"ref_id"`
	FeeType   string `json:"fee_type"`
	Fee       int    `json:"fee"`
}

type ErrorResponse struct {
	Code        int                 `json:"code"`
	Message     string              `json:"message"`
	Validations map[string][]string `json:"validations"`
}

type PaymentResponse struct {
	Data   PaymentResponseData `json:"data"`
	Errors ErrorResponse       `json:"errors"`
}

type VerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	Amount     int    `json:"amount"`
	OrderID    string `json:"order_id"`
}

type VerifyResponseData struct {
	Code     int    `json:"code"`
	RefID    int    `json:"ref_id"`
	CardPan  string `json:"card_pan"`
	CardHash string `json:"card_hash"`
	FeeType  string `json:"fee_type"`
	Fee      int    `json:"fee"`
}

type VerifyResponse struct {
	Data   VerifyResponseData `json:"data"`
	Errors ErrorResponse      `json:"errors"`
}

type InquiryRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
}

type InquiryResponseData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type InquiryResponse struct {
	Data   InquiryResponseData `json:"data"`
	Errors ErrorResponse       `json:"errors"`
}

type UnverifiedResponseData struct {
	Authorities []struct {
		Authority   string `json:"authority"`
		Amount      int    `json:"amount"`
		CallbackURL string `json:"callback_url"`
		Referer     string `json:"referer"`
		Date        string `json:"date"`
	} `json:"authorities"`
}

type UnverifiedResponse struct {
	Data   UnverifiedResponseData `json:"data"`
	Errors ErrorResponse          `json:"errors"`
}

func NewService(token, callbackURL string) *ZarinPalService {
	return &ZarinPalService{
		API:      token,
		Redirect: callbackURL,
	}
}
