package console

import (
	"testing"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"

	"github.com/stretchr/testify/assert"
)

func TestPolicyMakeCommand(t *testing.T) {
	policyMakeCommand := &PolicyMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := policyMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("UserPolicy").Once()
	err = policyMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/policies/user_policy.go"))

	mockContext.On("Argument", 0).Return("User/AuthPolicy").Once()
	err = policyMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/policies/user/auth_policy.go"))
	assert.True(t, file.Contain("internal/policies/user/auth_policy.go", "package user"))
	assert.True(t, file.Contain("internal/policies/user/auth_policy.go", "type AuthPolicy struct {"))

	assert.Nil(t, file.Remove("internal"))
}
