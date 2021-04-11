package entity

import (
	"github.com/locpham24/go-coffee/app/orm"
	"github.com/locpham24/go-coffee/infra"
)

func InitOrmInstances() {
	orm.User = orm.InitUserOrm(infra.GetDB().DB)
	//orm.Order = orm.Order.InitWithDB(infra.GetDB().DB)
	//orm.Point = orm.Point.InitWithDB(infra.GetDB().DB)
	// ....
}
