package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"
)

func TestListenerMakeCommand(t *testing.T) {
	listenerMakeCommand := &ListenerMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := listenerMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("gounityListen").Once()
	err = listenerMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/listeners/gounity_listen.go"))

	mockContext.On("Argument", 0).Return("gounity/listen").Once()
	err = listenerMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/listeners/gounity/listen.go"))
	assert.True(t, file.Contain("internal/listeners/gounity/listen.go", "package gounity"))
	assert.True(t, file.Contain("internal/listeners/gounity/listen.go", "type Listen struct {"))
	assert.Nil(t, file.Remove("internal"))
}
