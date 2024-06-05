package main

import (
	"github.com/PyMarcus/banking_auth/app"
	"github.com/PyMarcus/banking_auth/logger"
)

func main() {
	logger.Info("running...")
	app.Start()
}
