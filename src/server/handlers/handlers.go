package main
import(
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome ~Index")
}

func TodoList(w http.ResponseWriter,r *http.Request){
	todos :=Todos{
		Todo{Id:1,Name:"Write presentation"},
		Todo{Id:2,Name:"Host meetup"},
	}
	if err :=json.NewDecoder(todos);err !=nil{
		panic(err)
	}
}

func TodoItem(w http.ResponseWriter ,r *http.Request){
	vars :=mux.Vars(r)
	todoId :=vars["todoId"]
	fmt.Fprintln(w,"Todo Item:",todoId)
}
