package app

import (
	"encoding/json"
	"net/http"
  "fmt"

	"github.com/ThePianoDentist/kettle-stub-backend/utils"

	"github.com/ThePianoDentist/kettle-stub-backend/app_context"

	"github.com/ThePianoDentist/kettle-stub-backend/storage"
  "go.uber.org/zap"
)

func PostDrinker(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d storage.Drinker
	if err := decoder.Decode(&d); err != nil {
		utils.ErrorResp(appCtx.Lgr, w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	// I think reading body is weird/dumb. and defering before reading body leads to panic in some scenarios.
	// (add stack overflow link here if find/know)
	defer r.Body.Close()
	appCtx.Drinkers = append(appCtx.Drinkers, d.Name)
	err := appCtx.FcmController.SendFCM(appCtx.Token, fmt.Sprintf("%s demands Tea!", d.Name), "")
	if err != nil {
		appCtx.Lgr.Error("error publishing fcm message", zap.Error(err))
	}
	utils.SuccessResp(appCtx.Lgr, w, 200, map[string][]string{"drinkers": appCtx.Drinkers})
}
