package main  

import (
	"github.com/PyMarcus/go_banking_api/app"
	"github.com/PyMarcus/go_banking_api/logger"
)

func main(){
	logger.Info("Starting application...")
	app.Start()
}