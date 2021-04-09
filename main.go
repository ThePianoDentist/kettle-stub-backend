package main

import (
	"github.com/ThePianoDentist/kettle-stub-backend/app"

	"go.uber.org/zap"
)

func main() {
	lgr, _ := zap.NewProduction()
	defer lgr.Sync()
	a := app.NewApp(lgr)

	a.Run("0.0.0.0:8080")
}
