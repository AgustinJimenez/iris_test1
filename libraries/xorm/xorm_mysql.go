package xorm_mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Engine struct {
	*xorm.Engine
}

func (this *Engine) New() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "root:root@/iris?charset=utf8")
}
