package app

import (
	"net/http"


	"github.com/ThePianoDentist/toast-notification/app_context"
	"github.com/ThePianoDentist/toast-notification/utils"
)

func Poll(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
  if appCtx.Status != "CONFIRMED"{
    utils.SuccessResp(appCtx.Lgr, w, 200, map[string]int{"n": 0})
    return
  }
  utils.SuccessResp(appCtx.Lgr, w, 200, map[string]int{"n": len(appCtx.Drinkers)})
}
