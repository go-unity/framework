package cache

import (
	"github.com/go-unity/framework/contracts/cache"
	"github.com/go-unity/framework/contracts/config"
	"github.com/go-unity/framework/contracts/log"
)

type CacheStorage struct {
	cache.Driver
	config config.Config
	driver Driver
	log    log.Log
	stores map[string]cache.Driver
}

func NewCacheStorage(config config.Config, log log.Log, store string) (*CacheStorage, error) {
	driver := NewDriverImpl(config)
	instance, err := driver.New(store)
	if err != nil {
		return nil, err
	}

	return &CacheStorage{
		Driver: instance,
		config: config,
		driver: driver,
		log:    log,
		stores: map[string]cache.Driver{
			store: instance,
		},
	}, nil
}

func (app *CacheStorage) Store(name string) cache.Driver {
	if driver, exist := app.stores[name]; exist {
		return driver
	}

	instance, err := app.driver.New(name)
	if err != nil {
		app.log.Error(err)

		return nil
	}

	app.stores[name] = instance

	return instance
}
