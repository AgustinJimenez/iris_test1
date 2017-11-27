package services

import (
	"fmt"

	"../datamodels"
	"../repositories"
)

type UserService interface {
	Create(user datamodels.User)
}

/*
type userService struct {
	repo repositories.UserRepository
}
*/

func Create(user datamodels.User) datamodels.User {
	fmt.Printf("\nHERE IS SERVICE CREATE=======================>\n")
	if user.Id > 0 || user.Firstname == "" || user.Lastname == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		fmt.Printf("\n ERROR ON USER SERVICE \n")
		panic("unable to create this user")

	}
	user.GeneratePassword()
	return repositories.Store(user)
}
