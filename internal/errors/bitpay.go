package errors

func GetBitPayRequestError(code string) string {
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

func GetBitPayVerifyError(code float64) string {
	switch code {
	case -1:
		return "API ارسالی با نوع API تعریف شده در Bitpay سازگار نیست. API باید یک رشته 52 کاراکتری باشد."
	case -2:
		return "id_trans ارسال شده، داده عددی نمی باشد. این متغیر باید یک داده صحیح و عددی باشد مانند 123456."
	case -3:
		return "get_id ارسال شده، داده عددی نمی باشد. این متغیر باید یک داده صحیح و عددی باشد مانند 123456."
	case -4:
		return "چنین تراکنشی در پایگاه داده وجود ندارد یا موفقیت آمیز نبوده است. ممکن است تراکنش ناموفق باشد یا API اشتباه ارسال شده باشد."
	case 11:
		return "تراکنش از قبل وریفاي شده است."
	default:
		return "خطای ناشناخته."
	}
}
