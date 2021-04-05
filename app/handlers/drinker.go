package app

import (
	"encoding/json"
	"net/http"

	"github.com/ThePianoDentist/toast-notification/utils"

	"github.com/ThePianoDentist/toast-notification/app_context"

	"github.com/ThePianoDentist/toast-notification/storage"
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
	utils.SuccessResp(appCtx.Lgr, w, 200, map[string][]string{"drinkers": appCtx.Drinkers})
}
