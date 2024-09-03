package bitpay

// const (
// 	RequestURL = "https://bitpay.ir/payment/gateway-send"
// 	VerifyURL  = "https://bitpay.ir/payment/gateway-result-second"
// 	PayURL     = "https://bitpay.ir/payment/"
// )

const (
	RequestURL = "https://bitpay.ir/payment-test/gateway-send"
	VerifyURL  = "https://bitpay.ir/payment-test/gateway-result-second"
	PayURL     = "https://bitpay.ir/payment-test/"
	APIKeyTest     = "adxcv-zzadq-polkjsad-opp13opoz-1sdf455aadzmck1244567"
	CallbackURL    = "https://localhost"
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
		API:      token,
		Redirect: callbackURL,
	}

}
