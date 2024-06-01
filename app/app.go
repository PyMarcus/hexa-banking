package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PyMarcus/go_banking_api/config"
)

func Start() {
	config.NewGlobalConfigSet()
	addr := fmt.Sprintf("%s:%s", config.GlobalConfig.ServerIP, config.GlobalConfig.ServerPort)
	fmt.Println("Running on ", addr)
	log.Fatal(http.ListenAndServe(addr, router()))
}
