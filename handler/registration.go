package handler

import (
	"net/http"
	"ssoServer/repository"
	"encoding/json"
	"log"
	"ssoServer/jwt"
	"fmt"
	"gopkg.in/mgo.v2"
)

const KEY = "DataArt"

 func (u *userHandler) UserRegistration(writer http.ResponseWriter, request *http.Request) {
	var user repository.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	_, err = u.userRepo.GetUser(user.Name)
	if err != nil {
		err = u.userRepo.SaveUser(user)
		if err != nil {
			log.Println(err)
			http.Error(writer, "can not register user", http.StatusInternalServerError)
		}
	}

	token, err := jwt.GetToken(KEY)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(writer, token)
}

type userHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(s *mgo.Session) userHandler {
	  return userHandler{
		 repository.NewUseRepository(s),
	}
}