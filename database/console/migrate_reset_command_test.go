package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ormcontract "github.com/go-unity/framework/contracts/database/orm"
	"github.com/go-unity/framework/database/gorm"
	configmock "github.com/go-unity/framework/mocks/config"
	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/env"
)

func TestMigrateResetCommand(t *testing.T) {
	if env.IsWindows() {
		t.Skip("Skipping tests of using docker")
	}

	var (
		mockConfig *configmock.Config
		query      ormcontract.Query
	)

	if err := testDatabaseDocker.Fresh(); err != nil {
		t.Fatal(err)
	}

	beforeEach := func() {
		mockConfig = &configmock.Config{}
	}

	tests := []struct {
		name  string
		setup func()
	}{
		{
			name: "mysql",
			setup: func() {
				var err error
				docker := gorm.NewMysqlDocker(testDatabaseDocker)
				query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createMysqlMigrations()

			},
		},
		{
			name: "postgresql",
			setup: func() {
				var err error
				docker := gorm.NewPostgresqlDocker(testDatabaseDocker)
				query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createPostgresqlMigrations()
			},
		},
		{
			name: "sqlserver",
			setup: func() {
				var err error
				docker := gorm.NewSqlserverDocker(testDatabaseDocker)
				query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqlserverMigrations()
			},
		},
		{
			name: "sqlite",
			setup: func() {
				var err error
				docker := gorm.NewSqliteDocker("goravel")
				query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqliteMigrations()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()

			mockContext := &consolemocks.Context{}

			migrateCommand := NewMigrateCommand(mockConfig)
			assert.Nil(t, migrateCommand.Handle(mockContext))

			migrateResetCommand := NewMigrateResetCommand(mockConfig)
			assert.Nil(t, migrateResetCommand.Handle(mockContext))

			var agent Agent
			err := query.Where("name", "goravel").FirstOrFail(&agent)
			assert.Error(t, err)

			removeMigrations()
		})
	}
}
