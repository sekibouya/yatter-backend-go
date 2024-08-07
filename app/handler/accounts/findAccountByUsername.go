package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type FindRequest struct {
	Username string
}

func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {
	var req FindRequest
	username := chi.URLParam(r, "username")

	if err := json.NewDecoder(strings.NewReader(fmt.Sprintf("{\"username\":\"%v\"}", username))).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.accountUsecase.Fetch(ctx, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
