package bitpay

const (
	RequestURL = "https://bitpay.ir/payment/gateway-send"
	VerifyURL  = "https://bitpay.ir/payment/gateway-result-second"
	PayURL     = "https://bitpay.ir/payment/"
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

func NewService(token, callbackURL string) *BitPayIR {

	return &BitPayIR{
		API: token,
		Redirect: callbackURL,
	}

}