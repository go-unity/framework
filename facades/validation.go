package facades

import (
	"github.com/go-unity/framework/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
