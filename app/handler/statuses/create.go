package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string
	Medias  []map[string]interface{}
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	account_info := auth.AccountOf(r.Context()) // 認証情報を取得する

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.statusUsecase.AddStatus(ctx, int(account_info.ID), req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authentication", fmt.Sprintf("username %v", account_info.Username))
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//panic(fmt.Sprintf("Must Implement Status Creation And Check Acount Info %v", account_info))

}
