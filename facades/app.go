package facades

import (
	foundationcontract "github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
