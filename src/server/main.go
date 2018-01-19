package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/todos",TodoList)
	router.HandleFunc("/todos/{todoId}",TodoItem)

	log.Fatal(http.ListenAndServe(":8010",router)) 
	

}

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome ~Index")
}

func TodoList(w http.ResponseWriter,r *http.Request){

}

func TodoItem(w http.ResponseWriter ,r *http.Request){

}



