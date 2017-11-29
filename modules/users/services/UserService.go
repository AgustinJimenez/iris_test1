package services

import (
	"time"

	"../datamodels"
	"../repositories"
)

type UserServiceInterface interface {
	Create(password string, new_user datamodels.User) (datamodels.User, error)
	GetById(id int64) (datamodels.User, error)
	GetByUsernameAndPassword(username string, password string) (datamodels.User, error)
}

type UserService struct {
	repo repositories.UserRepository
}

func (this *UserService) Create(password string, new_user datamodels.User) (datamodels.User, error) {

	/*
		if new_user.Id > 0 || new_user.Firstname == "" || new_user.Lastname == "" || new_user.Username == "" || new_user.Email == "" || password == "" {
			return new_user, errors.New("unable to create this user")

		}
	*/
	hashed, err := new_user.GeneratePassword(password)

	if err != nil {
		return new_user, err
	}

	new_user.Password = hashed
	new_user.CreatedAt = time.Now()

	return this.repo.Store(new_user)
}

func (this *UserService) GetById(id int64) (datamodels.User, error) {

	return this.repo.GetById(id)
}

func (this *UserService) GetByUsernameAndPassword(username string, password string) (datamodels.User, error) {

	return this.repo.GetByUsernameAndPassword(username, password)
}
