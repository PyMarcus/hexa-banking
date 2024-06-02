package app

import (
	"encoding/json"
	"net/http"

	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/logger"
	"github.com/PyMarcus/go_banking_api/service"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	serviceInjected service.DefaultTransactionService
}

func (acc TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		writeResponse(w, http.StatusBadRequest, err.Error())
	}else{
		customerId := mux.Vars(r)["customer_id"]
		destinyId := mux.Vars(r)["destiny_id"]
		request.DestinyAccountId = destinyId
		response, err := acc.serviceInjected.DoTransaction(request, customerId)
		if err != nil{
			writeResponse(w, err.Code, err.Message)
		}else{
			logger.Info("Transaction is completed for " + response.AccountId)
			writeResponse(w, http.StatusCreated, response)
		}
	}
}
