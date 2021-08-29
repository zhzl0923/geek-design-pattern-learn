package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	data := []interface{}{1, 3, 5, 7, 8}
	iterator := NewArrayIterator(data)
	// i 用于测试
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.Current().(int))
		t.Log(iterator.Current())
		iterator.Next()
		i++
	}
}
