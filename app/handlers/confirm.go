package app

import (
	"net/http"

	"github.com/ThePianoDentist/toast-notification/utils"

	"github.com/ThePianoDentist/toast-notification/app_context"
)

func Confirm(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
  appCtx.Status = "CONFIRMED"
	// I think reading body is weird/dumb. and defering before reading body leads to panic in some scenarios.
	// (add stack overflow link here if find/know)
	err := appCtx.FcmController.SendFCM(appCtx.Token, fmt.Sprintf("Boiling kottle for %v", appCtx.Drinkers), "")
	if err != nil {
		appCtx.Lgr.Error("error publishing fcm message", zap.Error(err))
	}
	utils.SuccessResp(appCtx.Lgr, w, 200, map[string][]string{"drinkers": appCtx.Drinkers})
}
