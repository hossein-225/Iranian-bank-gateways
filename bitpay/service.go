package bitpay

const (
	RequestURL = "https://bitpay.ir/payment/gateway-send"
	VerifyURL  = "https://bitpay.ir/payment/gateway-result-second"
	PayURL     = "https://bitpay.ir/payment/"
	APIKey     = "your token"

	CallbackURL = "your callback URL"
)

type BitPayIR struct {
	API      string
	Redirect string
}

type BitPayRequest struct {
	Amount      int
	OrderID     string
	Name        string
	Email       string
	Description string
}
