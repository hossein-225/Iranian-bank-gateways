package saman

import "fmt"

func GetErrorMessage(code int) string {
	switch code {
	case 2:
		return "پرداخت با موفقیت انجام شد"
	case 1:
		return "کاربر از پرداخت انصراف داده است"
	case 3:
		return "پذیرنده فروشگاهی نامعتبر است"
	case 4:
		return "کاربر در بازه زمانی تعیین شده پاسخی ارسال نکرده است"
	case 5:
		return "پارامترهای ارسالی نامعتبر است"
	case 8:
		return "آدرس سرور پذیرنده نامعتبر است"
	case 10:
		return "توکن ارسال شده یافت نشد"
	case 11:
		return "با این شماره ترمینال فقط تراکنش های توکنی قابل پرداخت هستند."
	case 12:
		return "شماره ترمینال ارسال شده یافت نشد"
	case 21:
		return "محدودیت های مدل چند حسابی رعایت نشده"
	default:
		return fmt.Sprintf("کد خطا: %d. خطای ناشناخته.", code)
	}
}

func GetVerifyAndReverseErrorMessage(code int) string {
	switch code {
	case -9999:
		return "دریافت خطای استثنا"
	case -9998:
		return "دریافت تایم اوت 65 ثانیه ای"
	case -106:
		return "آدرس آی پی درخواستی غیرمجاز می باشد."
	case -105:
		return "ترمینال ارسالی در سیستم موجود نمی باشد."
	case -104:
		return "ترمینال ارسالی غیرفعال می باشد."
	case -18:
		return "IP Address فروشنده نامعتبر است."
	case -14:
		return "چنین تراکنشی تعریف نشده است."
	case -11:
		return "طول ورودی ها کمتر از حد مجاز است."
	case -10:
		return "رسید دیجیتالی به صورت Base64 نیست.(حاوی کاراکترهای غیرمجاز است)"
	case -8:
		return "طول ورودی ها بیشتر از حد مجاز است."
	case -7:
		return "رسید دیجیتال تهی است."
	case -6:
		return "بیش از نیم ساعت از زمان اجرای تراکنش گذشته است."
	case -2:
		return "تراکنش یافت نشد."
	case 0:
		return "موفق"
	case 1:
		return "کاربر از پرداخت انصراف داده است"
	case 2:
		return "درخواست تکراری می باشد."
	case 5:
		return "تراکنش برگشت خورده می باشد."
	default:
		return fmt.Sprintf("کد خطا: %d. خطای ناشناخته.", code)
	}
}
