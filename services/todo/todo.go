package todo

import (
	"cli/domain"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const FilePath = "data/todos.json"

func GetAll() {
  fmt.Println(string(readFromFile()))
}

func Add(todo *domain.Todo) {
  todos := getTodosFromFile()
  todos = append(todos, *todo)

  writeToFile(todos)
}

func Done(todoId string) bool {
  todos := getTodosFromFile()

  for i, t := range todos {
    if strconv.FormatInt(t.Id, 10) == todoId {
      t.Completed = true
      todos[i] = t
      writeToFile(todos)
      return true
    }
  }
  return false
}

func Delete(todoId string) bool {
  todos := getTodosFromFile()

  for i, t := range todos {
    if strconv.FormatInt(t.Id, 10) == todoId {
      todos := append(todos[:i], todos[i+1:]...)
      writeToFile(todos)
      return true
    }
  }
  return false
}

func getTodosFromFile() []domain.Todo {
  todosData := readFromFile()

  var todos []domain.Todo
  if err := json.Unmarshal(todosData, &todos); err != nil {
    fmt.Printf("Error unmarshalling existing todos %s", err)
    os.Exit(1)
  }
  return todos
}

func readFromFile() []byte {
  todosData, err := os.ReadFile(FilePath)
  if err != nil {
    fmt.Printf("Error reading all todos %s", err)
    os.Exit(1)
  }
  return todosData
}

func writeToFile(todos []domain.Todo) {
  jsonTodos, err := json.MarshalIndent(todos, "", "  ")
  if err != nil {
    fmt.Printf("error parsing todos into json %s", err)
    os.Exit(1)
  }

  os.WriteFile(FilePath, jsonTodos, 0777)
}