package proxy

import (
	"testing"
)

func TestUserControllerProxy_Login(t *testing.T) {
	type fields struct {
		userController *UserController
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test proxy login",
			fields: fields{
				userController: &UserController{},
			},
			args:    args{username: "test", password: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UserControllerProxy{
				userController: tt.fields.userController,
			}
			if err := p.Login(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UserControllerProxy.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
