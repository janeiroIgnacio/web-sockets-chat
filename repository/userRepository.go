package repository

import (
	"LefkasChat/config"
	"LefkasChat/models"
	"errors"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	SaveUser(message models.User)
	GetAllUsers()
	GetUserById(userId int)(models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (userRepo *userRepository) GetAllUsers() () {
}

func (userRepo *userRepository) GetUserById(userId int) (models.User, error) {
	user := models.User{}
	repo.db.First(&user, userId)
	if user.ID == 0 {
		return user, errors.New("Role not found at databse")
	}
	return user, nil
}

func (userRepo *userRepository) SaveUser(user models.User) {

}

var repo *userRepository

func GetUserRepository() UserRepository {
	if repo != nil{
		return repo
	}
	repo = &userRepository{
		db: config.GetDBInstance(),
	}
	return repo
}
