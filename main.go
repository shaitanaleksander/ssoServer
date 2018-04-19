package main

import (
	"github.com/gorilla/mux"
	"ssoServer/handler"
	"net/http"
	"ssoServer/repository"
	"gopkg.in/mgo.v2"
)
const urlDB = "mongodb://admin:shaitan13@cluster0-shard-00-00-rf2me.mongodb.net:27017,cluster0-shard-00-01-rf2me.mongodb.net:27017,cluster0-shard-00-02-rf2me.mongodb.net:27017/admin?&replicaSet=Cluster0-shard-0&authSource=admin"
var session *mgo.Session

func init() {
	session = repository.CreateSession(urlDB)
}
func main() {
	defer session.Close()
	u := handler.NewUserHandler(session)
	r := mux.NewRouter()
	r.HandleFunc("/registration", u.UserRegistration).Methods("POST")
	http.ListenAndServe(":8081", r)
}

