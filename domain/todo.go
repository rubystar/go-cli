package domain

import "time"

type Todo struct {
  Id int64 `json:"id"`
  Text string `json:"text"`
  Completed bool `json:"completed"`
}

func NewTodo(text string) *Todo {
  return &Todo {
    Id: time.Now().UnixNano(),
    Text: text,
    Completed: false,
  }
}
