package app

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PyMarcus/banking_auth/config"
	"github.com/PyMarcus/banking_auth/domain"
	"github.com/PyMarcus/banking_auth/logger"
	"github.com/PyMarcus/banking_auth/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Start() {
	config.NewGlobalConfigSet()
	addr := fmt.Sprintf("%s:%s", config.GlobalConfig.ServerIP, config.GlobalConfig.ServerPort)
	authRepo := domain.NewAuthRepository(getDBsettings())
	authService := service.NewAuthService(authRepo)

	router := mux.NewRouter()
	ah := AuthHandler{s: authService}
	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/private_route", ah.RequireAuth(ah.Pv)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(addr, router))
}

func getDBsettings() *sqlx.DB {
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
