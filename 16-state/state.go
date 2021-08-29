package state

import (
	"fmt"
	"strconv"
)

type Account struct {
	state   AccountState
	owner   string
	balance float64
}

func NewAccount(owner string, balance float64) *Account {
	account := &Account{owner: owner, balance: balance}
	account.state = NewNormalState(account)
	return account
}

func (a *Account) SetBalance(balance float64) {
	a.balance = balance
}

func (a *Account) SetState(state AccountState) {
	a.state = state
}

func (a *Account) Deposit(amount float64) {
	fmt.Println(a.owner + "存款" + float64ToString(amount))
	a.state.Deposit(amount)
	fmt.Println("当前余额为：" + float64ToString(a.balance))
	fmt.Println("当前账户状态：" + a.state.GetStateName())
}

func (a *Account) Withdraw(amount float64) {
	fmt.Println(a.owner + "取款" + float64ToString(amount))
	a.state.Withdraw(amount)
	fmt.Println("当前余额为：" + float64ToString(a.balance))
	fmt.Println("当前账户状态：" + a.state.GetStateName())
}

type AccountState interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetStateName() string
}

type NormalState struct {
	acc *Account
}

func NewNormalState(acc *Account) *NormalState {
	return &NormalState{acc: acc}
}

func (s NormalState) Deposit(amount float64) {
	s.acc.SetBalance(s.acc.balance + amount)
}

func (s NormalState) Withdraw(amount float64) {
	balance := s.acc.balance - amount

	if balance < -2000 {
		fmt.Println("操作受限")
		return
	}

	s.acc.SetBalance(balance)
	if balance < 0 && balance > -2000 {
		s.acc.SetState(NewOverdraftState(s.acc))
		return
	}

	if balance == -2000 {
		s.acc.SetState(NewRestrictedState(s.acc))
		return
	}

}

func (s NormalState) GetStateName() string {
	return "Normal State"
}

type OverdraftState struct {
	acc *Account
}

func NewOverdraftState(acc *Account) *OverdraftState {
	return &OverdraftState{acc: acc}
}

func (s OverdraftState) Deposit(amount float64) {
	s.acc.SetBalance(s.acc.balance + amount)
	if s.acc.balance >= 0 {
		s.acc.SetState(NewNormalState(s.acc))
	}
}

func (s OverdraftState) Withdraw(amount float64) {
	balance := s.acc.balance - amount

	if balance < -2000 {
		fmt.Println("操作受限")
		return
	}
	s.acc.SetBalance(s.acc.balance - amount)
	if balance == -2000 {
		s.acc.SetState(NewRestrictedState(s.acc))
		return
	}
}

func (s OverdraftState) GetStateName() string {
	return "Overdraft State"
}

type RestrictedState struct {
	acc *Account
}

func NewRestrictedState(acc *Account) *RestrictedState {
	return &RestrictedState{acc: acc}
}

func (s RestrictedState) Deposit(amount float64) {
	s.acc.SetBalance(s.acc.balance + amount)
	if s.acc.balance > -2000 && s.acc.balance < 0 {
		s.acc.SetState(NewOverdraftState(s.acc))
		return
	}
	if s.acc.balance >= 0 {
		s.acc.SetState(NewNormalState(s.acc))
		return
	}
}

func (s RestrictedState) Withdraw(amount float64) {
	fmt.Print("账号受限，取款失败！")
}

func (s RestrictedState) GetStateName() string {
	return "Restricted State"
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 10, 64)
}
