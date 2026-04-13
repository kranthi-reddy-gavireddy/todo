package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kranthi-reddy-gavireddy/internal/api/models"
)

type ITodo interface {
	Create(*models.Todo) error

	Update(oldTitle string, data *models.Todo) (*models.Todo, error)
	GetByID(id uuid.UUID) models.Todo
	GetAll() []models.Todo
	Delete(todo *models.Todo) error
}

var data = make(map[uuid.UUID]models.Todo)

var uniqueTitle = make(map[string]bool)

type Todo struct{}

func (t *Todo) Create(todo *models.Todo) error {

	if uniqueTitle[todo.Title] {
		return errors.New("title must be unique")
	}
	uniqueTitle[todo.Title] = true

	data[todo.ID] = *todo

	return nil
}

func (t *Todo) Update(oldTitle string, todo *models.Todo) (*models.Todo, error) {
	data[todo.ID] = *todo

	if oldTitle != todo.Title && uniqueTitle[todo.Title] {
		return nil, errors.New("title must be unique")
	}
	delete(uniqueTitle, todo.Title)
	uniqueTitle[todo.Title] = true

	return todo, nil
}

func (t *Todo) GetByID(id uuid.UUID) models.Todo {

	return data[id]
}

func (t *Todo) GetAll() []models.Todo {
	todos := make([]models.Todo, 0, len(data))

	for _, todo := range data {
		todos = append(todos, todo)
	}

	return todos
}

func (t *Todo) Delete(todo *models.Todo) error {
	delete(uniqueTitle, todo.Title)
	delete(data, todo.ID)

	return nil
}

func New() ITodo {
	return &Todo{}
}
