package main
import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct{
	Name string
	Method string
	Pattern string 
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes =Routes{
	Route{"Index","GET","/",Index},
	Route{"TodoIndex","GET","/todos",TodoIndex},
	Route{"TodoItem","GET","/todos/{todoId}",TodoItem},
	Route{"TodoCreate","POST","/todos",CreateTodo}
}