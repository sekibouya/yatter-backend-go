package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FindStatusByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	acc_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	dto, err := h.statusUsecase.FindStatusByID(ctx, acc_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Statuses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
