package datasource

import (
	xorm "./../../../libraries/xorm"
	model "./../datamodels"
)

var DB *xorm.Engine

func init() {

	DB := xorm.getInstance()
	err = DB.Sync2(new(model.User))
	if err != nil {
		panic(err)
	}

}

func all() {
	results, err := DB.QueryString("select * from users")
	if err != nil {
		panic(err)
	}
	return results
}
