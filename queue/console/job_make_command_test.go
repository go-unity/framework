package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/go-unity/framework/mocks/console"
	"github.com/go-unity/framework/support/file"
)

func TestJobMakeCommand(t *testing.T) {
	jobMakeCommand := &JobMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := jobMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("GounityJob").Once()
	err = jobMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/jobs/gounity_job.go"))

	mockContext.On("Argument", 0).Return("Gounity/Job").Once()
	err = jobMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/jobs/gounity/job.go"))
	assert.True(t, file.Contain("internal/jobs/gounity/job.go", "package gounity"))
	assert.True(t, file.Contain("internal/jobs/gounity/job.go", "type Job struct"))
	assert.Nil(t, file.Remove("internal"))
}
