package services

import (
	"main.go/dto"
	"main.go/models"
	"main.go/repository"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}