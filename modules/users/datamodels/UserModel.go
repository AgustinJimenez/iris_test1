package datamodels

import (
	"time"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the shake
// of the example, we will use this datamodel
// as the only one User model in our application.
type User struct {
	Id        int64     `json:"id" form:"id"`
	Firstname string    `json:"firstname" form:"firstname" xorm:"varchar(90)"`
	Lastname  string    `json:"lastname" form:"lastname" xorm:"varchar(90)"`
	Email     string    `json:"email" form:"email" xorm:"varchar(60)"`
	Username  string    `json:"username" form:"username" xorm:"varchar(60)"`
	Password  string    `json:"password" form:"password" xorm:"password" xorm:"varchar(200)"`
	CreatedAt time.Time `json:"created_at" form:"created_at" xorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" xorm:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

// IsValid can do some very very simple "low-level" data validations.
func (u *User) IsValid() bool {
	return u.Id > 0
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func (this *User) GeneratePassword() {
	hashed, err := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("\n ERROR ON USER MODEL \n")
		panic(err)
	}

	this.Password = string(hashed)
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func (this *User) CreatedAtFormated() string {
	return this.CreatedAt.Format("2006-01-02 15:04:05")
}
func (this *User) UpdatedAtFormated() string {
	return this.UpdatedAt.Format("2006-01-02 15:04:05")
}
