package services

import (
	"github.com/mskutle/hello-go/models"
	"gopkg.in/mgo.v2"
)

var todos []models.Todo

type TodoService struct{
	mdb *mgo.Session
}

func NewTodoService(session *mgo.Session) TodoService {
	return TodoService{ mdb: session}
}

func (s TodoService) AddTodo(todo models.Todo) models.Todo {
	todos = append(todos, todo)
	return todo
}

func (s TodoService) GetAll() []models.Todo {
	return todos
}
