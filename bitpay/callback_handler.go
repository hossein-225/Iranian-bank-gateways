package bitpay

import (
	"fmt"
	"net/http"
)

func (b *BitPayIR) HandleCallback(w http.ResponseWriter, r *http.Request) {
	transID := r.URL.Query().Get("trans_id")
	idGet := r.URL.Query().Get("id_get")

	if transID == "" || idGet == "" {
		http.Error(w, "Invalid callback parameters", http.StatusBadRequest)
		return
	}

	result, err := b.Verify(transID, idGet)
	if err != nil {
		http.Error(w, fmt.Sprintf("Verification failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Transaction verified successfully: %+v", result)

	// check the response code and do anything you need
}
