package foundation

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/go-unity/framework/auth"
	"github.com/go-unity/framework/cache"
	frameworkconfig "github.com/go-unity/framework/config"
	"github.com/go-unity/framework/console"
	"github.com/go-unity/framework/contracts/database/orm"
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/crypt"
	"github.com/go-unity/framework/database"
	"github.com/go-unity/framework/database/gorm"
	"github.com/go-unity/framework/event"
	"github.com/go-unity/framework/filesystem"
	"github.com/go-unity/framework/grpc"
	"github.com/go-unity/framework/hash"
	"github.com/go-unity/framework/http"
	frameworklog "github.com/go-unity/framework/log"
	"github.com/go-unity/framework/mail"
	cachemocks "github.com/go-unity/framework/mocks/cache"
	configmocks "github.com/go-unity/framework/mocks/config"
	consolemocks "github.com/go-unity/framework/mocks/console"
	ormmocks "github.com/go-unity/framework/mocks/database/orm"
	logmocks "github.com/go-unity/framework/mocks/log"
	queuemocks "github.com/go-unity/framework/mocks/queue"
	routemocks "github.com/go-unity/framework/mocks/route"
	"github.com/go-unity/framework/queue"
	"github.com/go-unity/framework/schedule"
	supportdocker "github.com/go-unity/framework/support/docker"
	"github.com/go-unity/framework/support/env"
	"github.com/go-unity/framework/support/file"
	frameworktranslation "github.com/go-unity/framework/translation"
	"github.com/go-unity/framework/validation"
)

type ApplicationTestSuite struct {
	suite.Suite
	app *Application
}

func TestApplicationTestSuite(t *testing.T) {
	assert.Nil(t, file.Create(".env", "APP_KEY=12345678901234567890123456789012"))

	suite.Run(t, new(ApplicationTestSuite))

	assert.Nil(t, file.Remove(".env"))
}

func (s *ApplicationTestSuite) SetupTest() {
	s.app = &Application{
		Container:     NewContainer(),
		publishes:     make(map[string]map[string]string),
		publishGroups: make(map[string]map[string]string),
	}
	App = s.app
}

func (s *ApplicationTestSuite) TestGetEnvPathInProductionEnvironment() {
	os.Setenv("ENV", "production")
	os.Setenv("AWS_SECRESTS_NAME", "nome_do_secreto")
	envPath := getEnvPath()
	s.Equal("nome_do_secreto", envPath)
}

func (s *ApplicationTestSuite) TestGetEnvPathInNonProductionEnvironment() {
	os.Unsetenv("ENV")
	os.Unsetenv("AWS_SECRESTS_NAME")
	envPath := getEnvPath()
	s.Equal(".env", envPath)
}

func (s *ApplicationTestSuite) TestPath() {
	s.Equal(filepath.Join("internal", "gounity.go"), s.app.Path("gounity.go"))
}

func (s *ApplicationTestSuite) TestBasePath() {
	s.Equal("gounity.go", s.app.BasePath("gounity.go"))
}

func (s *ApplicationTestSuite) TestConfigPath() {
	s.Equal(filepath.Join("config", "gounity.go"), s.app.ConfigPath("gounity.go"))
}

func (s *ApplicationTestSuite) TestDatabasePath() {
	s.Equal(filepath.Join("database", "gounity.go"), s.app.DatabasePath("gounity.go"))
}

func (s *ApplicationTestSuite) TestStoragePath() {
	s.Equal(filepath.Join("storage", "gounity.go"), s.app.StoragePath("gounity.go"))
}

func (s *ApplicationTestSuite) TestPublicPath() {
	s.Equal(filepath.Join("public", "gounity.go"), s.app.PublicPath("gounity.go"))
}

func (s *ApplicationTestSuite) TestPublishes() {
	s.app.Publishes("github.com/go-unity/sms", map[string]string{
		"config.go": "config.go",
	})
	s.Equal(1, len(s.app.publishes["github.com/go-unity/sms"]))
	s.Equal(0, len(s.app.publishGroups))

	s.app.Publishes("github.com/go-unity/sms", map[string]string{
		"config.go":  "config1.go",
		"config1.go": "config1.go",
	}, "public", "private")
	s.Equal(2, len(s.app.publishes["github.com/go-unity/sms"]))
	s.Equal("config1.go", s.app.publishes["github.com/go-unity/sms"]["config.go"])
	s.Equal(2, len(s.app.publishGroups["public"]))
	s.Equal("config1.go", s.app.publishGroups["public"]["config.go"])
	s.Equal(2, len(s.app.publishGroups["private"]))
}

func (s *ApplicationTestSuite) TestAddPublishGroup() {
	s.app.addPublishGroup("public", map[string]string{
		"config.go": "config.go",
	})
	s.Equal(1, len(s.app.publishGroups["public"]))

	s.app.addPublishGroup("public", map[string]string{
		"config.go":  "config1.go",
		"config1.go": "config1.go",
	})
	s.Equal(2, len(s.app.publishGroups["public"]))
	s.Equal("config1.go", s.app.publishGroups["public"]["config.go"])
}

func (s *ApplicationTestSuite) TestMakeArtisan() {
	serviceProvider := &console.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeArtisan())
}

func (s *ApplicationTestSuite) TestMakeAuth() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "auth.defaults.guard").Return("user").Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})
	s.app.Singleton(cache.Binding, func(app foundation.Application) (any, error) {
		return &cachemocks.Cache{}, nil
	})
	s.app.Singleton(database.BindingOrm, func(app foundation.Application) (any, error) {
		return &ormmocks.Orm{}, nil
	})

	serviceProvider := &auth.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeAuth(http.Background()))
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeCache() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "cache.default").Return("memory").Once()
	mockConfig.On("GetString", "cache.stores.memory.driver").Return("memory").Once()
	mockConfig.On("GetString", "cache.prefix").Return("gounity").Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})
	s.app.Singleton(frameworklog.Binding, func(app foundation.Application) (any, error) {
		return &logmocks.Log{}, nil
	})

	serviceProvider := &cache.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeCache())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeConfig() {
	serviceProvider := &frameworkconfig.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeConfig())
}

func (s *ApplicationTestSuite) TestMakeCrypt() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "app.key").Return("12345678901234567890123456789012").Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	serviceProvider := &crypt.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeCrypt())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeEvent() {
	s.app.Singleton(queue.Binding, func(app foundation.Application) (any, error) {
		return &queuemocks.Queue{}, nil
	})

	serviceProvider := &event.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeEvent())
}

func (s *ApplicationTestSuite) TestMakeGate() {
	serviceProvider := &auth.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeGate())
}

func (s *ApplicationTestSuite) TestMakeGrpc() {
	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return &configmocks.Config{}, nil
	})

	serviceProvider := &grpc.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeGrpc())
}

func (s *ApplicationTestSuite) TestMakeHash() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "hashing.driver", "argon2id").Return("argon2id").Once()
	mockConfig.On("GetInt", "hashing.argon2id.time", 4).Return(4).Once()
	mockConfig.On("GetInt", "hashing.argon2id.memory", 65536).Return(65536).Once()
	mockConfig.On("GetInt", "hashing.argon2id.threads", 1).Return(1).Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	serviceProvider := &hash.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeHash())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeLang() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "app.locale").Return("en").Once()
	mockConfig.On("GetString", "app.fallback_locale").Return("en").Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	serviceProvider := &frameworktranslation.ServiceProvider{}
	serviceProvider.Register(s.app)
	ctx := http.Background()

	s.NotNil(s.app.MakeLang(ctx))
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeLog() {
	serviceProvider := &frameworklog.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeLog())
}

func (s *ApplicationTestSuite) TestMakeMail() {
	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return &configmocks.Config{}, nil
	})
	s.app.Singleton(queue.Binding, func(app foundation.Application) (any, error) {
		return &queuemocks.Queue{}, nil
	})

	serviceProvider := &mail.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeMail())
}

func (s *ApplicationTestSuite) TestMakeOrm() {
	if env.IsWindows() {
		s.T().Skip("Skipping tests of using docker")
	}

	databaseDocker, err := supportdocker.InitDatabase()
	if err != nil {
		log.Fatalf("Init docker error: %s", err)
	}

	mysqlDocker := gorm.NewMysqlDocker(databaseDocker)
	_, err = mysqlDocker.New()
	s.Nil(err)

	config := databaseDocker.Mysql.Config()
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "database.default").Return("mysql").Once()
	mockConfig.On("Get", "database.connections.mysql.read").Return(nil).Once()
	mockConfig.On("Get", "database.connections.mysql.write").Return(nil).Once()
	mockConfig.On("GetString", "database.connections.mysql.driver").Return(orm.DriverMysql.String()).Twice()
	mockConfig.On("GetString", "database.connections.mysql.charset").Return("utf8mb4").Once()
	mockConfig.On("GetString", "database.connections.mysql.loc").Return("Local").Once()
	mockConfig.On("GetString", "database.connections.mysql.database").Return(config.Database).Once()
	mockConfig.On("GetString", "database.connections.mysql.host").Return("localhost").Once()
	mockConfig.On("GetString", "database.connections.mysql.username").Return(config.Username).Once()
	mockConfig.On("GetString", "database.connections.mysql.password").Return(config.Password).Once()
	mockConfig.On("GetString", "database.connections.mysql.prefix").Return("").Once()
	mockConfig.On("GetInt", "database.connections.mysql.port").Return(config.Port).Once()
	mockConfig.On("GetBool", "database.connections.mysql.singular").Return(true).Once()
	mockConfig.On("GetBool", "app.debug").Return(true).Once()
	mockConfig.On("GetInt", "database.pool.max_idle_conns", 10).Return(10)
	mockConfig.On("GetInt", "database.pool.max_open_conns", 100).Return(100)
	mockConfig.On("GetInt", "database.pool.conn_max_idletime", 3600).Return(3600)
	mockConfig.On("GetInt", "database.pool.conn_max_lifetime", 3600).Return(3600)

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	serviceProvider := &database.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeOrm())
	s.Nil(databaseDocker.Stop())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeQueue() {
	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return &configmocks.Config{}, nil
	})

	serviceProvider := &queue.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeQueue())
}

func (s *ApplicationTestSuite) TestMakeRateLimiter() {
	serviceProvider := &http.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeRateLimiter())
}

func (s *ApplicationTestSuite) TestMakeRoute() {
	mockConfig := &configmocks.Config{}

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	mockRoute := &routemocks.Route{}
	s.app.Singleton("gounity.route", func(app foundation.Application) (any, error) {
		return mockRoute, nil
	})

	s.NotNil(s.app.MakeRoute())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeSchedule() {
	s.app.Singleton(console.Binding, func(app foundation.Application) (any, error) {
		return &consolemocks.Artisan{}, nil
	})
	s.app.Singleton(frameworklog.Binding, func(app foundation.Application) (any, error) {
		return &logmocks.Log{}, nil
	})

	serviceProvider := &schedule.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeSchedule())
}

func (s *ApplicationTestSuite) TestMakeStorage() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "filesystems.default").Return("local").Once()
	mockConfig.On("GetString", "filesystems.disks.local.driver").Return("local").Once()
	mockConfig.On("GetString", "filesystems.disks.local.root").Return("").Once()
	mockConfig.On("GetString", "filesystems.disks.local.url").Return("").Once()

	s.app.Singleton(frameworkconfig.Binding, func(app foundation.Application) (any, error) {
		return mockConfig, nil
	})

	serviceProvider := &filesystem.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeStorage())
	mockConfig.AssertExpectations(s.T())
}

func (s *ApplicationTestSuite) TestMakeValidation() {
	serviceProvider := &validation.ServiceProvider{}
	serviceProvider.Register(s.app)

	s.NotNil(s.app.MakeValidation())
}
