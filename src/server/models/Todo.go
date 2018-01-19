package main

import "time"

type Todo struct{
	TodoId int32 `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"hascompleted"`
	Due time.Time `json:"duetime"`
}
type Todos []Todo
