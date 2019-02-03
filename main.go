package main

import (
	"TrainingCamp/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/add", addUser).Methods("POST").Queries()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func logUser(user model.User) {
	log.Printf("Add user. Name: %v Sname: %v Email: %v", user.Name, user.Sname, user.Email)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{r.FormValue("name"), r.FormValue("sname"), r.FormValue("email")}
	logUser(user)
	w.Write([]byte("succeed"))
}
