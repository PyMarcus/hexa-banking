package app

import (
	"encoding/json"
	"net/http"

	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/logger"
	"github.com/PyMarcus/go_banking_api/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	serviceInjected service.DefaultAccountService
}

func (acc AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		customer_id := mux.Vars(r)["customer_id"]
		request.CustomerId = customer_id
		response, err := acc.serviceInjected.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			logger.Info("Created user " + response.AccountId)
			writeResponse(w, http.StatusCreated, response)
		}
	}
}
