package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"
)

func TestEventMakeCommand(t *testing.T) {
	eventMakeCommand := &EventMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := eventMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("gounityEvent").Once()
	err = eventMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/events/gounity_event.go"))

	mockContext.On("Argument", 0).Return("gounity/Event").Once()
	err = eventMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/events/gounity/event.go"))
	assert.True(t, file.Contain("internal/events/gounity/event.go", "package gounity"))
	assert.True(t, file.Contain("internal/events/gounity/event.go", "type Event struct {"))
	assert.Nil(t, file.Remove("internal"))
}
