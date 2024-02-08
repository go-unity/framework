package console

import (
	"testing"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"

	"github.com/stretchr/testify/assert"
)

func TestRuleMakeCommand(t *testing.T) {
	requestMakeCommand := &RuleMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := requestMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("Uppercase").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/rules/uppercase.go"))

	mockContext.On("Argument", 0).Return("User/Phone").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/rules/user/phone.go"))
	assert.True(t, file.Contain("internal/rules/user/phone.go", "package user"))
	assert.True(t, file.Contain("internal/rules/user/phone.go", "type Phone struct"))
	assert.Nil(t, file.Remove("internal"))
}
