package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory_CountNumOfFiles(t *testing.T) {
	num := NewDirectory().CountNumOfFiles()
	assert.Equal(t, num, 3)
}

func TestDirectory_CountSizeOfFiles(t *testing.T) {
	size := NewDirectory().CountSizeOfFiles()
	t.Log(size)
}
