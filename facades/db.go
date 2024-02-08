package facades

import (
	"github.com/go-unity/framework/contracts/database/orm"
)

func Orm() orm.Orm {
	return App().MakeOrm()
}
