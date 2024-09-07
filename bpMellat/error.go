package bpMellat

import "errors"

var bankErrors = map[string]string{
	"0":   "",
	"11":  "شماره کارت نامعتبر است",
	"12":  "موجودی کافی نیست",
	"13":  "رمز نادرست است",
	"14":  "تعداد دفعات وارد کردن رمز بیش از حد مجاز است",
	"15":  "کارت نامعتبر است",
	"16":  "دفعات برداشت وجه بیش از حد مجاز است",
	"17":  "کاربر از انجام تراکنش منصرف شده است",
	"18":  "تاریخ انقضای کارت گذشته است",
	"19":  "مبلغ برداشت وجه بیش از حد مجاز است",
	"111": "صادر کننده کارت نامعتبر است",
	"112": "خطای سوییچ صادر کننده کارت",
	"113": "پاسخی از صادر کننده کارت دریافت نشد",
	"114": "دارنده این کارت مجاز به انجام این تراکنش نیست",
	"21":  "پذیرنده نامعتبر است",
	"23":  "خطای امنیتی رخ داده است",
	"24":  "اطلاعات کاربری پذیرنده نامعتبر است",
	"25":  "مبلغ نامعتبر است",
	"31":  "پاسخ نامعتبر است",
	"32":  "فرمت اطلاعات وارد شده صحیح نمی باشد",
	"33":  "حساب نامعتبر است",
	"34":  "خطای سیستمی",
	"35":  "تاریخ نامعتبر است",
	"41":  "شماره درخواست تکراری است",
	"42":  "تراکنش Sale یافت نشد",
	"43":  "قبلا درخواست Verify داده شده است",
	"44":  "درخواست Verify یافت نشد",
	"45":  "تراکنش Settle شده است",
	"46":  "تراکنش Settle نشده است",
	"47":  "تراکنش Settle یافت نشد",
	"48":  "تراکنش Reverse شده است",
	"49":  "تراکنش Refund یافت نشد",
	"412": "شناسه قبض نادرست است",
	"413": "شناسه پرداخت نادرست است",
	"414": "سازمان صادر کننده قبض نامعتبر است",
	"415": "مدت زمان مجاز برای انجام تراکنش به پایان رسیده است",
	"416": "خطا در ثبت اطلاعات",
	"417": "شناسه پرداخت کننده نامعتبر است",
	"418": "اشکال در تعریف اطلاعات مشتری",
	"419": "تعداد دفعات ورود اطلاعات از حد مجاز گذشته است",
	"421": "IP نامعتبر است",
	"51":  "تراکنش تکراری است",
	"54":  "تراکنش مرجع موجود نیست",
	"55":  "تراکنش نامعتبر است",
	"61":  "خطا در واریز",
	"62":  "مسير بازگشت به سايت در دامنه ثبت شده برای پذيرنده قرار ندارد",
	"98":  "سقف استفاده از رمز ايستا به پايان رسيده است",
	"995": "تعلق کارت بانکي به مشتری احراز نشد",
}

func getBankError(result string) error {
	if errMsg, ok := bankErrors[result]; ok {
		if errMsg == "" {
			return nil
		}
		return errors.New(errMsg)
	}
	return errors.New("خطای ناشناخته")
}
