package bitpay

func HandleErrorCode(code string) string {
	switch code {
	case "-1":
		return "API ارسالی با نوع API تعریف شده در Bitpay سازگار نیست."
	case "-2":
		return "مقدار amount به عنوان عددی کمتر از حداقل مقدار (1000 ریال) ارسال شده است."
	case "-3":
		return "مقدار redirect نباید نال باشد."
	case "-4":
		return "API شما نادرست است یا آدرس ثبت شده در Bitpay با آدرس ارسال شده همخوانی ندارد."
	case "-5":
		return "مشکل در اتصال به درگاه، لطفا مجدداً تلاش کنید یا از طریق تیکت پیگیری کنید."
	default:
		return "خطای ناشناخته."
	}
}
