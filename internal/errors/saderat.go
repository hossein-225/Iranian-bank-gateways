package errors

import "fmt"

func HandleCallBackErrors(code string) string {
	switch code {
	case "0":
		return "تراکنش با موفقیت انجام شده است."
	case "-1":
		return "کاربر دکمه انصراف را در صفحه پرداخت فشرده است."
	case "-2":
		return "زمان انجام تراکنش برای کاربر به اتمام رسیده است."
	default:
		return fmt.Sprintf("کد خطا: %s. خطای ناشناخته.", code)
	}
}

func HandleServiceErrors(code int) string {
	switch code {
	case -1:
		return "تراکنش پیدا نشد."
	case -2:
		return "در زمان دریافت توکن به دلیل عدم وجود (عدم انطباق) IP و یا به دلیل بسته بودن خروجی پورت 8081 از سمت Host این خطا ایجاد می گردد.	تراکنش قبلا Reverse شده است."
	case -3:
		return "Total Error خطای عمومی - خطای Exception ها"
	case -4:
		return "امکان انجام درخواست برای این تراکنش وجود ندارد."
	case -5:
		return "آدرس IP نامعتبر میباشد."
	case -6:
		return "عدم فعال بودن سرویس برگشت تراکنش برای پذیرنده"
	case 10:
		return "توکن ارسال شده یافت نشد"
	default:
		return fmt.Sprintf("کد خطا: %d. خطای ناشناخته.", code)
	}
}
