package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string                   `json:"status"`
	Medias  []map[string]interface{} `json:"omitempty"`
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

	var status statusStruct
	var statuses []statusStruct
	status.ID = dto.Status.ID
	status.Account = *dto.Account
	status.Content = dto.Status.Content
	status.CreatedAt = dto.Status.CreatedAt
	statuses = append(statuses, status)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//panic(fmt.Sprintf("Must Implement Status Creation And Check Acount Info %v", account_info))

}
