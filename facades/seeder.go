package facades

import (
	"github.com/go-unity/framework/contracts/database/seeder"
)

func Seeder() seeder.Facade {
	return App().MakeSeeder()
}
