package file

import (
	"os"
	"testing"

	"github.com/go-unity/framework/testing/file"

	"github.com/stretchr/testify/assert"
)

func TestClientOriginalExtension(t *testing.T) {
	assert.Equal(t, ClientOriginalExtension("logo.jpg"), "jpg")
}

func TestContain(t *testing.T) {
	assert.True(t, Contain("../constant.go", "const Version"))
}

func TestCreate(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/goravel/gounity.txt"
	assert.Nil(t, Create(path, `goravel`))
	assert.Equal(t, 1, file.GetLineNum(path))
	assert.True(t, Exists(path))
	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/goravel"))
}

func TestExists(t *testing.T) {
	assert.True(t, Exists("file.go"))
}

func TestExtension(t *testing.T) {
	extension, err := Extension("file.go")
	assert.Nil(t, err)
	assert.Equal(t, "txt", extension)
}

func TestLastModified(t *testing.T) {
	ti, err := LastModified("../../logo.jpg", "UTC")
	assert.Nil(t, err)
	assert.NotNil(t, ti)
}

func TestMimeType(t *testing.T) {
	mimeType, err := MimeType("../../logo.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "image/jpeg", mimeType)
}

func TestRemove(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/goravel/gounity.txt"
	assert.Nil(t, Create(path, `goravel`))

	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/goravel"))
}

func TestSize(t *testing.T) {
	size, err := Size("../../logo.jpg")
	assert.Nil(t, err)
	assert.Equal(t, int64(114598), size)
}
