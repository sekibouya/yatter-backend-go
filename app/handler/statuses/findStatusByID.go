package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type FindRequest struct {
	AccountID int
}

func (h *handler) FindStatusByID(w http.ResponseWriter, r *http.Request) {
	var req FindRequest
	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(strings.NewReader(fmt.Sprintf("{\"id\":\"%v\"}", id))).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.statusUsecase.FindStatusByID(ctx, req.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
