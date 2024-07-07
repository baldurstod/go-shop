package handlers

import (
	"net/http"
)

type RecoveryHandler struct {
	Handler http.Handler
}

func (h *RecoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//h.log(err)
		}
	}()

	h.Handler.ServeHTTP(w, r)
}
