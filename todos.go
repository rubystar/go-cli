package main

import (
	"cli/domain"
	"cli/services/todo"
	"flag"
	"fmt"
	"os"
)


func main() {
  createCmd := flag.NewFlagSet("create", flag.ExitOnError)
  todoText := createCmd.String("text", "blank todo", "what needs to be done?")

  doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
  doneTodoId := doneCmd.String("id", "", "what is the id of todo that is completed?")

  deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
  todoId := deleteCmd.String("id", "", "what is the id of todo that needs to be deleted?")

  indexCmd := flag.NewFlagSet("index", flag.ExitOnError)

  if len(os.Args) < 2 {
    fmt.Println("At least one sub command is required. allowed sub commands: 'create', 'delete', 'index', 'done'")
    os.Exit(1)
  }

  cmd := os.Args[1]

  switch cmd {
  case "create":
    createCmd.Parse(os.Args[2:])
    newTodo := domain.NewTodo(*todoText)
    todo.Add(newTodo)

  case "index":
    indexCmd.Parse(os.Args[2:])
    todo.GetAll()

  case "delete":
    deleteCmd.Parse(os.Args[2:])

    if ok := todo.Delete(*todoId); !ok {
      fmt.Println("couldn't delete todo")
      os.Exit(1)
    }

  case "done":
    doneCmd.Parse(os.Args[2:])

    if ok := todo.Done(*doneTodoId); !ok {
      fmt.Println("couldn't mark as done")
      os.Exit(1)
    }

  default:
    fmt.Println("unknown command")
    os.Exit(1)
  }

}
