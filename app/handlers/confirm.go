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
	utils.SuccessResp(appCtx.Lgr, w, 200, map[string][]string{"drinkers": appCtx.Drinkers})
}
