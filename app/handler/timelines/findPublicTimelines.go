package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const defaultLimit = 40
const maxLimit = 80

func (h *handler) findPublicTimelines(w http.ResponseWriter, r *http.Request) {
	only_media, err := strconv.ParseBool(r.FormValue("only_media"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	since_id, err := strconv.Atoi(r.FormValue("since_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = defaultLimit
	}

	ctx := r.Context()

	if maxLimit < limit {
		limit = maxLimit
	}

	dto, err := h.timelineUsecase.FindPublicTimelines(ctx, only_media, since_id, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timelines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
