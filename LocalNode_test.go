package lun

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var localNode *LocalNode

func init() {
	localNode = NewLocalNode()
}

func TestLocalRun(t *testing.T) {
	output := localNode.Run("echo abc")
	assert.Equal(t, "abc\n", string(output))
}

func TestLocalSafeRun(t *testing.T) {
	output, err := localNode.SafeRun("echo abc")
	assert.NoError(t, err)
	assert.Equal(t, "abc\n", string(output))
}

func TestLocalRunError(t *testing.T) {
	assert.Panics(t, func() { localNode.Run("exit 1") })
}

func TestLocalSafeRunError(t *testing.T) {
	_, err := localNode.SafeRun("exit 1")
	assert.Error(t, err)
}

func TestLocalRunCD(t *testing.T) {
	_ = localNode.Run("cd /")
	output := localNode.Run("pwd")
	assert.Equal(t, "/\n", string(output))
	_ = localNode.Run("cd ~")
	output = localNode.Run("pwd")
	assert.Equal(t, os.Getenv("HOME")+"\n", string(output))
}
