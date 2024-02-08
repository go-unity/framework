package facades

import (
	"github.com/go-unity/framework/contracts/auth"
	"github.com/go-unity/framework/contracts/auth/access"
	"github.com/go-unity/framework/contracts/http"
)

func Auth(ctx http.Context) auth.Auth {
	return App().MakeAuth(ctx)
}

func Gate() access.Gate {
	return App().MakeGate()
}
