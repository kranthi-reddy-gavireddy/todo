package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kranthi-reddy-gavireddy/internal/api/models"
	"github.com/kranthi-reddy-gavireddy/internal/api/service"
)

type response map[string]any

type IHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.ITodo
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": "invalid request body"})
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		writeJSON(w, http.StatusBadRequest, response{"error": "title is required"})
		return
	}

	res, err := h.service.Create(req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, res)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := readID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": err.Error()})
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": "invalid request body"})
		return
	}

	if req.UpdatedTitle == nil || strings.TrimSpace(*req.UpdatedTitle) == "" {
		writeJSON(w, http.StatusBadRequest, response{"error": "updated_title is required"})
		return
	}

	if req.IsCompleted == nil {
		writeJSON(w, http.StatusBadRequest, response{"error": "is_completed is required"})
		return
	}

	if req.PreviousTitle == nil {
		existing, err := h.service.GetByID(id)
		if err != nil {
			writeJSON(w, http.StatusNotFound, response{"error": err.Error()})
			return
		}
		req.PreviousTitle = &existing.Title
	}

	res, err := h.service.Update(req, id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := readID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": err.Error()})
		return
	}

	res, err := h.service.GetByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, response{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAll()
	if err != nil {
		writeJSON(w, http.StatusNotFound, response{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := readID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{"error": err.Error()})
		return
	}

	if err := h.service.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, response{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, response{"message": "todo deleted successfully"})
}

func readID(r *http.Request) (uuid.UUID, error) {
	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		idParam = r.URL.Query().Get("id")
	}
	if idParam == "" {
		return uuid.Nil, errors.New("id is required")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, errors.New("invalid id")
	}

	return id, nil
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || payload == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func New(service service.ITodo) IHandler {
	return &Handler{service: service}
}
