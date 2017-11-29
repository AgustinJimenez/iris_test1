package repositories

import (
	"errors"

	"../../../libraries/xorm"
	"../datamodels"
	"github.com/go-xorm/xorm"
)

type UserRepository struct {
	User datamodels.User
	Db   xorm_mysql.Engine
}

func (this *UserRepository) StartDB() (*xorm.Engine, error) {
	DB, err := this.Db.New()
	if err != nil {
		return new(xorm.Engine), err
	}
	err = DB.Sync2(new(datamodels.User))
	if err != nil {
		return new(xorm.Engine), err
	}

	return DB, nil
}

func (this *UserRepository) Store(new_user datamodels.User) (datamodels.User, error) {

	DB, err := this.StartDB()
	if err != nil {
		return new_user, err
	}
	affected, err := DB.Insert(&new_user)

	if err != nil {
		return new_user, err
	}
	if affected == 0 {
		return new_user, errors.New("unable to create this user")
	} else {
		return new_user, nil
	}

}

func (this *UserRepository) GetByUsernameAndPassword(username string, password string) (datamodels.User, error) {

	valid_password := false

	DB, err := this.StartDB()

	if err != nil {
		return datamodels.User{}, err
	}

	user := &datamodels.User{}
	has, err := DB.Where("username = ?", username).Get(user)

	if err != nil {
		return datamodels.User{}, err
	}

	if has {
		valid_password = user.ValidatePassword(password)
	} else {
		return *user, errors.New("User not found")
	}

	if valid_password != true {
		return *user, errors.New("Invalid Password")
	} else {
		return *user, nil
	}

}

func (this *UserRepository) GetById(id int64) (datamodels.User, error) {

	DB, err := this.StartDB()
	if err != nil {
		return datamodels.User{}, err
	}

	user := &datamodels.User{}
	has, err := DB.Id(3).Get(user)

	if err != nil {
		return *user, err
	}

	if has {
		return *user, nil
	} else {
		return *user, errors.New("User not found")
	}

}
