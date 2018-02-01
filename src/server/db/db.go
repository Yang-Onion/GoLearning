package main
import "fmt"

var currentId int32

var todos Todos

func init(){

}

func FindTodo(id int32) Todo{
  for _,t:=range todos{
	  if t.Id ==id{
		  return t
	  }
  }
  return Todo{}
}

func CreateTodo(t Todo) Todo{
	currentId +=1
	t.Id = currentId
	todos =append(todos,t)
	return t
}

func DeleteTodo(id int32) error{
	for i,t :=range todos{
		if t.Id == id{
			todos =append(todos[:i],todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Can not find Todo with id of %d to delete ",id)
}