package foundation

import (
	"context"
	"fmt"
	"sync"

	"github.com/gookit/color"

	"github.com/go-unity/framework/auth"
	"github.com/go-unity/framework/cache"
	"github.com/go-unity/framework/config"
	"github.com/go-unity/framework/console"
	authcontract "github.com/go-unity/framework/contracts/auth"
	accesscontract "github.com/go-unity/framework/contracts/auth/access"
	cachecontract "github.com/go-unity/framework/contracts/cache"
	configcontract "github.com/go-unity/framework/contracts/config"
	consolecontract "github.com/go-unity/framework/contracts/console"
	cryptcontract "github.com/go-unity/framework/contracts/crypt"
	ormcontract "github.com/go-unity/framework/contracts/database/orm"
	seerdercontract "github.com/go-unity/framework/contracts/database/seeder"
	eventcontract "github.com/go-unity/framework/contracts/event"
	filesystemcontract "github.com/go-unity/framework/contracts/filesystem"
	foundationcontract "github.com/go-unity/framework/contracts/foundation"
	grpccontract "github.com/go-unity/framework/contracts/grpc"
	hashcontract "github.com/go-unity/framework/contracts/hash"
	httpcontract "github.com/go-unity/framework/contracts/http"
	logcontract "github.com/go-unity/framework/contracts/log"
	mailcontract "github.com/go-unity/framework/contracts/mail"
	queuecontract "github.com/go-unity/framework/contracts/queue"
	routecontract "github.com/go-unity/framework/contracts/route"
	schedulecontract "github.com/go-unity/framework/contracts/schedule"
	testingcontract "github.com/go-unity/framework/contracts/testing"
	translationcontract "github.com/go-unity/framework/contracts/translation"
	validationcontract "github.com/go-unity/framework/contracts/validation"
	"github.com/go-unity/framework/crypt"
	"github.com/go-unity/framework/database"
	"github.com/go-unity/framework/event"
	"github.com/go-unity/framework/filesystem"
	"github.com/go-unity/framework/grpc"
	"github.com/go-unity/framework/hash"
	"github.com/go-unity/framework/http"
	goravellog "github.com/go-unity/framework/log"
	"github.com/go-unity/framework/mail"
	"github.com/go-unity/framework/queue"
	"github.com/go-unity/framework/route"
	"github.com/go-unity/framework/schedule"
	"github.com/go-unity/framework/testing"
	"github.com/go-unity/framework/translation"
	"github.com/go-unity/framework/validation"
)

type instance struct {
	concrete any
	shared   bool
}

type Container struct {
	bindings  sync.Map
	instances sync.Map
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Bind(key any, callback func(app foundationcontract.Application) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: false})
}

func (c *Container) BindWith(key any, callback func(app foundationcontract.Application, parameters map[string]any) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: false})
}

func (c *Container) Instance(key any, ins any) {
	c.bindings.Store(key, instance{concrete: ins, shared: true})
}

func (c *Container) Make(key any) (any, error) {
	return c.make(key, nil)
}

func (c *Container) MakeArtisan() consolecontract.Artisan {
	instance, err := c.Make(console.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(consolecontract.Artisan)
}

func (c *Container) MakeAuth(ctx httpcontract.Context) authcontract.Auth {
	instance, err := c.MakeWith(auth.BindingAuth, map[string]any{
		"ctx": ctx,
	})
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(authcontract.Auth)
}

func (c *Container) MakeCache() cachecontract.Cache {
	instance, err := c.Make(cache.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(cachecontract.Cache)
}

func (c *Container) MakeConfig() configcontract.Config {
	instance, err := c.Make(config.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(configcontract.Config)
}

func (c *Container) MakeCrypt() cryptcontract.Crypt {
	instance, err := c.Make(crypt.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(cryptcontract.Crypt)
}

func (c *Container) MakeEvent() eventcontract.Instance {
	instance, err := c.Make(event.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(eventcontract.Instance)
}

func (c *Container) MakeGate() accesscontract.Gate {
	instance, err := c.Make(auth.BindingGate)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(accesscontract.Gate)
}

func (c *Container) MakeGrpc() grpccontract.Grpc {
	instance, err := c.Make(grpc.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(grpccontract.Grpc)
}

func (c *Container) MakeHash() hashcontract.Hash {
	instance, err := c.Make(hash.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(hashcontract.Hash)
}

func (c *Container) MakeLang(ctx context.Context) translationcontract.Translator {
	instance, err := c.MakeWith(translation.Binding, map[string]any{
		"ctx": ctx,
	})
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(translationcontract.Translator)
}

func (c *Container) MakeLog() logcontract.Log {
	instance, err := c.Make(goravellog.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(logcontract.Log)
}

func (c *Container) MakeMail() mailcontract.Mail {
	instance, err := c.Make(mail.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(mailcontract.Mail)
}

func (c *Container) MakeOrm() ormcontract.Orm {
	instance, err := c.Make(database.BindingOrm)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(ormcontract.Orm)
}

func (c *Container) MakeQueue() queuecontract.Queue {
	instance, err := c.Make(queue.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(queuecontract.Queue)
}

func (c *Container) MakeRateLimiter() httpcontract.RateLimiter {
	instance, err := c.Make(http.BindingRateLimiter)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(httpcontract.RateLimiter)
}

func (c *Container) MakeRoute() routecontract.Route {
	instance, err := c.Make(route.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(routecontract.Route)
}

func (c *Container) MakeSchedule() schedulecontract.Schedule {
	instance, err := c.Make(schedule.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(schedulecontract.Schedule)
}

func (c *Container) MakeStorage() filesystemcontract.Storage {
	instance, err := c.Make(filesystem.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(filesystemcontract.Storage)
}

func (c *Container) MakeTesting() testingcontract.Testing {
	instance, err := c.Make(testing.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(testingcontract.Testing)
}

func (c *Container) MakeValidation() validationcontract.Validation {
	instance, err := c.Make(validation.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(validationcontract.Validation)
}

func (c *Container) MakeView() httpcontract.View {
	instance, err := c.Make(http.BindingView)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(httpcontract.View)
}

func (c *Container) MakeSeeder() seerdercontract.Facade {
	instance, err := c.Make(database.BindingSeeder)

	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(seerdercontract.Facade)
}

func (c *Container) MakeWith(key any, parameters map[string]any) (any, error) {
	return c.make(key, parameters)
}

func (c *Container) Singleton(key any, callback func(app foundationcontract.Application) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: true})
}

func (c *Container) make(key any, parameters map[string]any) (any, error) {
	binding, ok := c.bindings.Load(key)
	if !ok {
		return nil, fmt.Errorf("binding not found: %+v", key)
	}

	if parameters == nil {
		instance, ok := c.instances.Load(key)
		if ok {
			return instance, nil
		}
	}

	bindingImpl := binding.(instance)
	switch concrete := bindingImpl.concrete.(type) {
	case func(app foundationcontract.Application) (any, error):
		concreteImpl, err := concrete(App)
		if err != nil {
			return nil, err
		}
		if bindingImpl.shared {
			c.instances.Store(key, concreteImpl)
		}

		return concreteImpl, nil
	case func(app foundationcontract.Application, parameters map[string]any) (any, error):
		concreteImpl, err := concrete(App, parameters)
		if err != nil {
			return nil, err
		}

		return concreteImpl, nil
	default:
		c.instances.Store(key, concrete)

		return concrete, nil
	}
}
