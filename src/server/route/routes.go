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

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route :=range routes{
		router.Methods(route.Method).
			   Path(route.Pattern).	
			   Name(route.Method).
			   Handler(route.HandlerFunc)
	}
	return router
}

var routes =Routes{
	Route{"Index","Get","/",Index},
	Route{"TodoIndex","Get","/todos",TodoIndex},
	Route{"TodoItem","Get","/todos/{todoId}",TodoItem},
}