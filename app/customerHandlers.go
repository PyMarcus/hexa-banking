package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PyMarcus/go_banking_api/config"
	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/PyMarcus/go_banking_api/logger"
	"github.com/PyMarcus/go_banking_api/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type CustomerHandler struct {
	serviceInjected service.DefaultCustomerService
}

// postgres, postgres, your_password
func router() *mux.Router {
	mux := mux.NewRouter()
	methodsHandlers(mux)
	return mux
}

func methodsHandlers(mux *mux.Router) {
	// customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDBClient()
	customerRepo := domain.NewCustomerRepositoryDb(dbClient)
	accountRepo := domain.NewAccountRepository(dbClient)
	transactionRepo := domain.NewTransactionRepository(dbClient)
	// handler
	customerHandler := CustomerHandler{service.NewCustomerService(customerRepo)}
	accountHandler :=  AccountHandler{service.NewAccountService(accountRepo)}
	transactionHandler :=  TransactionHandler{service.NewTransactionService(transactionRepo)}

	// multiplex customers
	mux.HandleFunc("/api/v1/customers", customerHandler.getCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/customers/{customer_id:[0-9]+}", customerHandler.getCustomerById).Methods(http.MethodGet)

	// multiplex account
	mux.HandleFunc("/api/v1/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)

	// multiplex transaction
	mux.HandleFunc("/api/v1/customers/{customer_id:[0-9]+}/account/{destiny_id:[0-9]+}", transactionHandler.NewTransaction).Methods(http.MethodPost)

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
	var customers []dto.CustomerResponse
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

func getDBClient() *sqlx.DB{
	user := config.GlobalConfig.DBUser
	password := config.GlobalConfig.DBPassword
	host := config.GlobalConfig.DBHost
	database := config.GlobalConfig.DBDatabase

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, password, host, database)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Error("Fail to connect into database err: " + err.Error())
		db.Close()
		panic(err)
	}
	return db
}