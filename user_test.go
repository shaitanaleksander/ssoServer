package main

import (
	"testing"
	"ssoServer/repository"
	"gopkg.in/mgo.v2"
	"fmt"
)

var urlDB = "mongodb://admin:shaitan13@cluster0-shard-00-00-rf2me.mongodb.net:27017,cluster0-shard-00-01-rf2me.mongodb.net:27017,cluster0-shard-00-02-rf2me.mongodb.net:27017/admin?&replicaSet=Cluster0-shard-0&authSource=admin"
var session *mgo.Session
var userRepo repository.UserRepository

func init() {
	session = repository.CreateSession(urlDB)
	userRepo = repository.NewUseRepository(session)
}

func TestSaveUser(t *testing.T) {
	defer userRepo.CloseSession()
	user := repository.User{
		Name:     "koly",
		Password: "ytrewq",
	}
	err := userRepo.SaveUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUser(t *testing.T) {
	defer userRepo.CloseSession()
	user := "koly"
	err := userRepo.DeleteUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUser(t *testing.T) {
	defer userRepo.CloseSession()
	u, err := userRepo.GetUser("any")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}


