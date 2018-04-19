package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name     string
	Password string
}

type UserRepository struct {
	session *mgo.Session
}

func (u *UserRepository) SaveUser(user User) error {

	c := u.session.DB("auth").C("user")
	err := c.Insert(user)
	return err

}

func (u *UserRepository) DeleteUser(name string) error {
	c := u.session.DB("auth").C("user")
	err := c.Remove(bson.M{"name": name})
	return err
}

func (u *UserRepository) GetUser(name string) (User, error) {
	c := u.session.DB("auth").C("user")
	user1 := User{}
	err := c.Find(bson.M{"name": name}).One(&user1)
	return user1, err
}

func (u *UserRepository) CloseSession() {
	u.session.Close()
}

func NewUseRepository(session *mgo.Session) UserRepository {
	return UserRepository{session}
}