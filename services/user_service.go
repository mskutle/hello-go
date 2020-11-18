package services

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mskutle/hello-go/models"
)

var users []models.User

type UserService struct{
	mdb *mgo.Database
}

func NewUserService(db *mgo.Database) UserService {
	return UserService{ mdb: db}
}

func (s UserService) GetAll() (error, []models.User) {
	users := make([]models.User, 0)
	err := s.mdb.C("users").Find(nil).All(&users)

	return err, users
}

func (s UserService) AddUser(user models.User) (error, *models.User) {
	err := s.mdb.C("users").Insert(user)
	if err != nil {
		return err, nil
	}
	newlyCreatedUser := s.GetByUsername(user.Username)
	if newlyCreatedUser == nil {
		return errors.New("User not found}"), nil
	}

	return nil, newlyCreatedUser
}

func (s UserService) GetByUsername(username string) *models.User {
	users := make([]models.User, 1)
	err := s.mdb.C("users").Find(bson.M{"username": username }).All(&users)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(users) > 0 {
		return &users[0]
	}
	return nil
}

func (s UserService) Delete(username string) error {
	return s.mdb.C("users").Remove(bson.M{"username": username})
}
