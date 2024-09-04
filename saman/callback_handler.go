package saman

import (
	"net/http"
	"strconv"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	response := CallbackResponse{
		MID:              r.FormValue("MID"),
		State:            r.FormValue("State"),
		Status:           r.FormValue("Status"),
		RRN:              r.FormValue("RRN"),
		RefNum:           r.FormValue("RefNum"),
		ResNum:           r.FormValue("ResNum"),
		TerminalId:       r.FormValue("TerminalId"),
		TraceNo:          r.FormValue("TraceNo"),
		SecurePan:        r.FormValue("SecurePan"),
		HashedCardNumber: r.FormValue("HashedCardNumber"),
	}

	amount, err := strconv.Atoi(r.FormValue("Amount"))
	if err == nil {
		response.Amount = amount
	}

	wage, err := strconv.Atoi(r.FormValue("Wage"))
	if err == nil {
		response.Wage = wage
	}

	if statusCode, err := strconv.Atoi(response.Status); err == nil {
		errorMessage := GetErrorMessage(statusCode)
		w.Write([]byte(errorMessage))
	} else {
		w.Write([]byte("کد وضعیت نامعتبر است"))
	}

	w.WriteHeader(http.StatusOK)
}
