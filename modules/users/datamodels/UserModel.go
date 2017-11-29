package datamodels

import (
	base "../../core/datamodels"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int64  `json:"id" form:"id"`
	Firstname  string `xorm:"varchar(90)" json:"firstname" form:"firstname"`
	Lastname   string `xorm:"varchar(90)" json:"lastname" form:"lastname"`
	Email      string `xorm:"varchar(60)" json:"email" form:"email"`
	Username   string `xorm:"varchar(60) "json:"username" form:"username"`
	Password   []byte `xorm:"password" json:"password" form:"password"`
	base.Model `xorm:"extends"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsValid() bool {
	return u.Id > 0
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func (this *User) GeneratePassword(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword will check if passwords are matched.
func (this *User) ValidatePassword(password string) bool {

	if err := bcrypt.CompareHashAndPassword(this.Password, []byte(password)); err != nil {
		return false
	}
	return true
}

/*
func (this *User) CreatedAtFormated() string {
	return this.CreatedAt.Format("2006-01-02 15:04:05")
}
func (this *User) UpdatedAtFormated() string {
	return this.UpdatedAt.Format("2006-01-02 15:04:05")
}
*/
