package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

type statusStruct struct {
	ID        int
	Account   object.Account
	Content   string
	CreatedAt time.Time
}

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

	var status statusStruct
	var statuses []statusStruct
	for _, e := range dto.Statuses {
		status.ID = e.Status.ID
		status.Account = *e.Account
		status.Content = e.Status.Content
		status.CreatedAt = e.Status.CreatedAt
		statuses = append(statuses, status)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
