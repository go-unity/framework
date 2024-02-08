package facades

import (
	"github.com/go-unity/framework/contracts/testing"
)

func Testing() testing.Testing {
	return App().MakeTesting()
}
