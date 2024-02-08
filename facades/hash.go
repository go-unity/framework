package facades

import (
	"github.com/go-unity/framework/contracts/hash"
)

func Hash() hash.Hash {
	return App().MakeHash()
}
