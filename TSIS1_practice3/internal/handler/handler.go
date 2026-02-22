package handler

import "net/http"

type Handler struct {
	users *UsersHandler
}

func NewHandler(uc UsersUC) *Handler {
	return &Handler{users: NewUsersHandler(uc)}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/users", h.users.Users)
	mux.HandleFunc("/users/", h.users.UserByID)
}
