package app

import (
	"encoding/json"
	"net/http"

	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/PyMarcus/go_banking_api/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	serviceInjected service.DefaultCustomerService
}

// postgres, postgres, your_password
func router() *mux.Router {
	mux := mux.NewRouter()
	getMethodsHandlers(mux)
	return mux
}

func getMethodsHandlers(mux *mux.Router) {
	// customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	mux.HandleFunc("/api/v1/customers", customerHandler.getCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/customers/{customer_id:[0-9]+}", customerHandler.getCustomerById).Methods(http.MethodGet)
}

func (ch *CustomerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if(status == ""){
		customers, err := ch.serviceInjected.GetAllCustomers()
		if err != nil {
			writeResponse(w, err.Code, err)
		} else {
			writeResponse(w, http.StatusOK, customers)
		}
	}else{
		ch.getCustomersByStatus(w, status)
	}	
}

func (ch *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]
	customers, err := ch.serviceInjected.GetCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomersByStatus(w http.ResponseWriter, status string){
	var customers []domain.Customer
	var err *errs.AppError

	if(status == "active"){
		customers, err = ch.serviceInjected.GetActiveCustomers()
	}else{
		customers, err = ch.serviceInjected.GetInactiveCustomers()
	}
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil{
		panic(err.Error())
	}
}
