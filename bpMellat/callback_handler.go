package bpMellat

import (
	"log"
	"net/http"
	"strconv"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	refId := r.FormValue("RefId")
	resCode := r.FormValue("ResCode")
	saleOrderId, err := strconv.ParseInt(r.FormValue("SaleOrderId"), 10, 64)
	if err != nil {
		log.Printf("Invalid SaleOrderId: %v", err)
		http.Error(w, "Invalid SaleOrderId", http.StatusBadRequest)
		return
	}
	saleReferenceId, err := strconv.ParseInt(r.FormValue("SaleReferenceId"), 10, 64)
	if err != nil {
		log.Printf("Invalid SaleReferenceId: %v", err)
		http.Error(w, "Invalid SaleReferenceId", http.StatusBadRequest)
		return
	}
	cardHolderInfo := r.FormValue("CardHolderInfo")

	response := CallbackResponse{
		RefId:           refId,
		ResCode:         resCode,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
		CardHolderInfo:  cardHolderInfo,
	}

	log.Printf("Received callback: %+v", response)

	// check the response code and do anything you need
}
