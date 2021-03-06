package app

import (
	"net/http"


	"github.com/ThePianoDentist/kettle-stub-backend/app_context"
	"github.com/ThePianoDentist/kettle-stub-backend/utils"
)

func Poll(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
  if appCtx.Status != "CONFIRMED"{
    utils.SuccessResp(appCtx.Lgr, w, 444, map[string]int{"n": 0})
    return
  }
  utils.SuccessResp(appCtx.Lgr, w, 200 + len(appCtx.Drinkers), map[string]int{"n": len(appCtx.Drinkers)})
}
