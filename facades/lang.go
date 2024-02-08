package facades

import (
	"context"

	"github.com/go-unity/framework/contracts/translation"
)

func Lang(ctx context.Context) translation.Translator {
	return App().MakeLang(ctx)
}
