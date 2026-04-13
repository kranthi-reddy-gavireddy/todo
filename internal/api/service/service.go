package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kranthi-reddy-gavireddy/internal/api/models"
	"github.com/kranthi-reddy-gavireddy/internal/api/repository"
)

type ITodo interface {
	Create(req models.CreateTodoRequest) (*models.TodoResponse, error)
	Update(req models.UpdateTodoRequest, id uuid.UUID) (*models.TodoResponse, error)
	GetByID(id uuid.UUID) (*models.TodoResponse, error)
	GetAll() ([]models.TodoResponse, error)
	Delete(id uuid.UUID) error
}

type Todo struct {
	repo repository.ITodo
}

func (t *Todo) Create(m models.CreateTodoRequest) (*models.TodoResponse, error) {
	req := &models.Todo{
		ID:        uuid.New(),
		Title:     m.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := t.repo.Create(req); err != nil {
		return nil, err
	}

	res := &models.TodoResponse{
		ID:          req.ID.String(),
		Title:       req.Title,
		IsCompleted: req.IsCompleted,
	}

	return res, nil
}

func (t *Todo) Update(req models.UpdateTodoRequest, id uuid.UUID) (*models.TodoResponse, error) {
	data := t.repo.GetByID(id)

	if data == (models.Todo{}) {
		return nil, errors.New("todo not found")
	}

	data.UpdatedAt = time.Now()
	data.Title = *req.UpdatedTitle
	data.IsCompleted = *req.IsCompleted

	res, err := t.repo.Update(*req.PreviousTitle, &data)
	if err != nil {
		return nil, err
	}

	return &models.TodoResponse{
		ID:          res.ID.String(),
		Title:       res.Title,
		IsCompleted: res.IsCompleted,
	}, nil
}

func (t *Todo) GetByID(id uuid.UUID) (*models.TodoResponse, error) {
	data := t.repo.GetByID(id)

	if data == (models.Todo{}) {
		return nil, errors.New("todo not found")
	}

	res := &models.TodoResponse{
		ID:          data.ID.String(),
		Title:       data.Title,
		IsCompleted: data.IsCompleted,
	}

	return res, nil
}

func (t *Todo) GetAll() ([]models.TodoResponse, error) {
	data := t.repo.GetAll()

	if len(data) == 0 {
		return nil, errors.New("no todos found")
	}

	res := make([]models.TodoResponse, 0, len(data))
	for _, todo := range data {
		res = append(res, models.TodoResponse{
			ID:          todo.ID.String(),
			Title:       todo.Title,
			IsCompleted: todo.IsCompleted,
		})
	}

	return res, nil
}

func (t *Todo) Delete(id uuid.UUID) error {
	data := t.repo.GetByID(id)

	if data == (models.Todo{}) {
		return errors.New("todo not found")
	}

	return t.repo.Delete(&data)
}

func New(repo repository.ITodo) ITodo {
	return &Todo{repo: repo}
}
