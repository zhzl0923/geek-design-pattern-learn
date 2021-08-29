package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	coffee := Coffee{}
	milk := &Milk{drink: coffee}
	desc := milk.Desc()
	cost := milk.Cost()
	assert.Equal(t, desc, "牛奶咖啡")
	assert.Equal(t, cost, 2)
}
