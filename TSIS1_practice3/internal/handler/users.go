package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"golang/internal/repository/_postgres/users"
	"golang/pkg/modules"
)

type UsersUC interface {
	GetUsers() ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(user modules.User) (int, error)
	UpdateUser(id int, user modules.User) error
	DeleteUser(id int) error
}

type UsersHandler struct {
	uc UsersUC
}

func NewUsersHandler(uc UsersUC) *UsersHandler {
	return &UsersHandler{uc: uc}
}

// /users -> GET, POST
func (h *UsersHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := h.uc.GetUsers()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, data)

	case http.MethodPost:
		var u modules.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Email) == "" {
			writeError(w, http.StatusBadRequest, "name and email are required")
			return
		}

		id, err := h.uc.CreateUser(u)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, map[string]int{"id": id})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// /users/{id} -> GET, PUT, DELETE
func (h *UsersHandler) UserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		u, err := h.uc.GetUserByID(id)
		if err != nil {
			if err == users.ErrUserNotFound {
				writeError(w, http.StatusNotFound, "user not found")
				return
			}
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, u)

	case http.MethodPut:
		var u modules.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Email) == "" {
			writeError(w, http.StatusBadRequest, "name and email are required")
			return
		}

		if err := h.uc.UpdateUser(id, u); err != nil {
			if err == users.ErrUserNotFound {
				writeError(w, http.StatusNotFound, "user not found")
				return
			}
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "updated"})

	case http.MethodDelete:
		if err := h.uc.DeleteUser(id); err != nil {
			if err == users.ErrUserNotFound {
				writeError(w, http.StatusNotFound, "user not found")
				return
			}
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
