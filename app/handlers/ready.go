package app

import (
	"net/http"

	"github.com/ThePianoDentist/kettle-stub-backend/app_context"
	"go.uber.org/zap"
)

func Ready(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
	err := appCtx.FcmController.SendFCM(appCtx.Token, "Kettle boiled", "")
	if err != nil {
		appCtx.Lgr.Error("error publishing fcm message", zap.Error(err))
	}
  appCtx.Status = "" // reset so can accept another round
}
