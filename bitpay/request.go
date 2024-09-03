package bitpay

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strconv"
)

func (b *BitPayIR) Request(request BitPayRequest) (string, error) {
	data := neturl.Values{}
	data.Set("api", b.API)
	data.Set("redirect", b.Redirect)
	data.Set("amount", strconv.Itoa(request.Amount))
	data.Set("factorId", request.OrderID)
	data.Set("name", request.Name)
	data.Set("email", request.Email)
	data.Set("description", request.Description)

	resp, err := http.PostForm(RequestURL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseStr := string(body)
	if i, err := strconv.Atoi(responseStr); i <= 0 || err != nil {
		return "", fmt.Errorf("خطا در درخواست: %s", HandleErrorCode(responseStr))
	} else {
		addr := PayURL + "gateway-" + responseStr + "-get"
		return addr, nil
	}
}
