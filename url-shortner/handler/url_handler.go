package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rahulp18/url-shortner/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

type shortenRequest struct {
	URL string `json:"url"`
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	code := h.service.Shorten(req.URL)

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": "http://localhost:8080/" + code,
	})
}
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	longURL, ok := h.service.Resolve(code)
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
