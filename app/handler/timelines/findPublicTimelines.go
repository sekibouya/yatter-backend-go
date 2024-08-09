package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"yatter-backend-go/app/domain/object"
)

type timelineStruct struct {
	ID        int            `json:"id"`
	Account   object.Account `json:"account"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"create_at"`
}

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

	var timeline timelineStruct
	var timelines []timelineStruct
	for _, e := range dto.Timelines {
		timeline.ID = e.Status.ID
		timeline.Account = *e.Account
		timeline.Content = e.Status.Content
		timeline.CreatedAt = e.Status.CreatedAt
		timelines = append(timelines, timeline)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timelines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
