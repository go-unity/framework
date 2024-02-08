package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"
)

func TestMakeCommand(t *testing.T) {
	makeCommand := &MakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("CleanCache").Once()
	assert.Nil(t, makeCommand.Handle(mockContext))
	assert.True(t, file.Exists("internal/console/commands/clean_cache.go"))

	mockContext.On("Argument", 0).Return("Gounity/CleanCache").Once()
	assert.Nil(t, makeCommand.Handle(mockContext))
	assert.True(t, file.Exists("internal/console/commands/gounity/clean_cache.go"))
	assert.True(t, file.Contain("internal/console/commands/gounity/clean_cache.go", "package gounity"))
	assert.True(t, file.Contain("internal/console/commands/gounity/clean_cache.go", "type CleanCache struct"))

	assert.Nil(t, file.Remove("internal"))
}
