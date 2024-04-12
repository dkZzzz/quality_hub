package rpc

import (
	"context"

	"github.com/dkZzzz/quality_hub/proto/userpb"
	"github.com/dkZzzz/quality_hub/srv/usersrv"
)

var (
	userClient userpb.UserServiceClient
)

// 调用客户端，给服务端发送请求

// Login 登录
func Login(req *userpb.UserLoginReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// Register 注册
func Register(req *userpb.UserRegisterReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// Logout 退出登录
func Logout(req *userpb.UserLogoutReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.Logout(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// ModifyUsername 修改用户名
func ModifyUsername(req *userpb.UserModifyUsernameReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.ModifyUsername(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// ModifyEmail 修改邮箱
func ModifyEmail(req *userpb.UserModifyEmailReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.ModifyEmail(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// ModifyPassword 修改密码
func ModifyPassword(req *userpb.UserModifyPasswordReq) (*userpb.Resp, error) {
	userClient = usersrv.Client
	ctx := context.Background()
	rsp, err := userClient.ModifyPassword(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}
