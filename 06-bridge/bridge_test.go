package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSevereNotification_Notify(t *testing.T) {
	sender := NewSmsSender([]string{"123214325"})
	notification := NewSevereNotification(sender)
	err := notification.Notify("hello")
	assert.Nil(t, err)
}
