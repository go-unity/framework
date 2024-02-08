package facades

import "github.com/go-unity/framework/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
