package orm

import (
	"github.com/locpham24/go-coffee/infra"
)

func InitOrmInstances() {
	User = InitUserOrm(infra.GetDB().DB)
	//orm.Order = orm.Order.InitWithDB(infra.GetDB().DB)
	//orm.Point = orm.Point.InitWithDB(infra.GetDB().DB)
	// ....
}
