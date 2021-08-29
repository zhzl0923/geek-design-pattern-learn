package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountState(t *testing.T) {
	account := NewAccount("zhang", 2000)
	account.Withdraw(1000)
	assert.Equal(t, account.state.GetStateName(), "Normal State")
	account.Withdraw(1100)
	assert.Equal(t, account.state.GetStateName(), "Overdraft State")
	account.Withdraw(1900)
	assert.Equal(t, account.state.GetStateName(), "Restricted State")
	account.Deposit(3000)
	assert.Equal(t, account.state.GetStateName(), "Normal State")
}
