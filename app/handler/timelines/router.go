package timelines

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	timelineUsecase usecase.Timeline
}

// Create Handler for `/v1/timelines/public`
func NewRouter(u usecase.Timeline) http.Handler {
	r := chi.NewRouter()
	h := &handler{
		u,
	}
	r.Get("/public", h.findPublicTimelines)
	return r
}
