package app_context

import (
	"net/http"

	"github.com/ThePianoDentist/kettle-stub-backend/fcm_client"

	"go.uber.org/zap"
)

// Just a simple wrapper so we can pass the global state (i.e. hub) into every request.
// Handler just needs to match the interface (so needs a ServeHTTP method)
type CtxHandler struct {
	AppCtx         *AppContext
	CtxHandlerFunc func(*AppContext, http.ResponseWriter, *http.Request)
}

func (h *CtxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.CtxHandlerFunc(h.AppCtx, w, r)
}

type AppContext struct {
	Lgr           *zap.Logger
	Token         string
	FcmController *fcm_client.FCMController
  Drinkers []string
  Status string // could be enum but fine for stub
}
