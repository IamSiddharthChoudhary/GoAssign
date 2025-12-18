package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	name string
	date string
}

func handlePost(res http.ResponseWriter, req *http.Request) {
	var usr User
	err := json.NewDecoder(req.Body).Decode(&usr)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

}

func handleUpadate(res http.ResponseWriter, r *http.Request) {

}

func handleDelete(res http.ResponseWriter, r *http.Request) {

}

func handleGet(res http.ResponseWriter, r *http.Request) {

}

func main() {
	multi := http.NewServeMux()
	multi.HandleFunc("POST /users", handlePost)
	multi.HandleFunc("PUY /users", handlePost)
	multi.HandleFunc("DELETE /users", handlePost)
	multi.HandleFunc("UPDATE /users", handlePost)

	fmt.Println("The sever running at 8000")

	http.ListenAndServe(":8000", multi)
}
