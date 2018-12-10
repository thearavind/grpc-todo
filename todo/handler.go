package todo
import (
  "log"
  "golang.org/x/net/context"
  "github.com/satori/go.uuid"
)
// Server represents the gRPC server
type Server struct {
  Todos []*TodoObject
}
// SayHello generates response to a Ping request
func (s *Server) AddTodo(ctx context.Context, newTodo *AddTodoParams) (*TodoObject, error) {
  log.Printf("Received new task %s", newTodo.Task)
  todoObject := &TodoObject{
    Id: uuid.NewV1().String(),
    Task: newTodo.Task,
  }
  s.Todos = append(s.Todos, todoObject)
  return todoObject, nil
}

func (s *Server) GetTodos(ctx context.Context, _ *GetTodoParams) (*TodoResponse, error) {
  log.Printf("get tasks")
  return &TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, delTodo *DeleteTodoParams) (*DeleteResponse, error) {
  var updatedTodos []*TodoObject
  for index, todo := range s.Todos {
    if(todo.Id == delTodo.Id) {
      updatedTodos = append(s.Todos[:index], s.Todos[index + 1:]...)
        break;
    }
  }
  s.Todos = updatedTodos
  return &DeleteResponse{Message: "success"}, nil
}