package facades

import (
	"github.com/go-unity/framework/contracts/crypt"
)

func Crypt() crypt.Crypt {
	return App().MakeCrypt()
}
