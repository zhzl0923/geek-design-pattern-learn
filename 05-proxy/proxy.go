package proxy

import (
	"fmt"
	"time"
)

type IUserController interface {
	Login(username, password string) error
}

type UserController struct {
}

func (u *UserController) Login(username, password string) error {
	return nil
}

type UserControllerProxy struct {
	userController *UserController
}

func NewUserControllerProxy(userController *UserController) *UserControllerProxy {
	return &UserControllerProxy{userController: userController}
}

func (p *UserControllerProxy) Login(username, password string) error {
	start := time.Now()

	if err := p.userController.Login(username, password); err != nil {
		return err
	}
	fmt.Printf("user login cost time: %s", time.Since(start))
	return nil
}
