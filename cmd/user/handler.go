package user

import (
	"context"
	"strconv"
	"time"

	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/pkg/jwt"
	"github.com/dkZzzz/quality_hub/proto/userpb"
)

// rpc服务的具体接口业务逻辑

var (
	param_error       string = "参数错误"
	create_user_error string = "创建用户失败"
	username_error    string = "用户名错误"
	password_error    string = "密码错误"
	logout_error      string = "登出错误"
	token_error       string = "鉴权token错误"
	modify_error      string = "修改失败"
)

var (
	register_success string = "注册成功"
	login_success    string = "登录成功"
	logout_success   string = "退出登录成功"
	modify_success   string = "修改成功"
)

type UserServerImpl struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServerImpl) Register(ctx context.Context, in *userpb.UserRegisterReq) (*userpb.Resp, error) {
	if in.Username == "" || in.Password == "" || in.Email == "" {
		return &userpb.Resp{
			Code: 400,
			Msg:  param_error,
		}, nil
	}
	// 创建用户
	userID, err := mysql.CreateUser(ctx, in.Username, in.Password, in.Email, mysql.Argon2ParamVar)
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  create_user_error,
		}, nil
	}

	// 创建完自动登陆
	token, err := jwt.GenerateJWT(in.Username, strconv.Itoa(userID))
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  create_user_error,
		}, nil
	}

	err = redis.Client.WithContext(ctx).Set(in.Username, token, 24*time.Hour).Err()
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  create_user_error,
		}, nil
	}

	return &userpb.Resp{
		Code:  200,
		Token: &token,
		Msg:   register_success,
	}, nil
}

func (s *UserServerImpl) Login(ctx context.Context, in *userpb.UserLoginReq) (*userpb.Resp, error) {
	if in.Username == "" || in.Password == "" {
		return &userpb.Resp{
			Code: 400,
			Msg:  param_error,
		}, nil
	}

	userID, err := mysql.CheckUser(ctx, in.Username, in.Password)
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  password_error,
		}, nil
	}

	// 登录
	token, err := jwt.GenerateJWT(in.Username, strconv.Itoa(userID))
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  password_error,
		}, nil

	}

	err = redis.Client.WithContext(ctx).Set(in.Username, token, 24*time.Hour).Err()
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  create_user_error,
		}, nil
	}

	return &userpb.Resp{
		Code:  200,
		Token: &token,
		Msg:   login_success,
	}, nil
}

func (s *UserServerImpl) Logout(ctx context.Context, in *userpb.UserLogoutReq) (*userpb.Resp, error) {
	if in.Username == "" {
		return &userpb.Resp{
			Code: 500,
			Msg:  username_error,
		}, nil
	}

	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &userpb.Resp{
			Code: 403,
			Msg:  token_error,
		}, nil
	}
	err := redis.Client.WithContext(ctx).Del(in.Username).Err()
	if err != nil {
		return &userpb.Resp{
			Code: 500,
			Msg:  logout_error,
		}, nil
	}
	return &userpb.Resp{
		Code: 200,
		Msg:  logout_success,
	}, nil
}

func (s *UserServerImpl) ModifyUsername(ctx context.Context, in *userpb.UserModifyUsernameReq) (*userpb.Resp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &userpb.Resp{
			Code: 403,
			Msg:  token_error,
		}, nil
	}

	err := mysql.ModifyUsername(ctx, in.Username, in.NewUsername)
	if err != nil {
		return &userpb.Resp{
			Code: 200,
			Msg:  modify_error,
		}, nil
	}

	return &userpb.Resp{
		Code: 200,
		Msg:  modify_success,
	}, nil
}

func (s *UserServerImpl) ModifyEmail(ctx context.Context, in *userpb.UserModifyEmailReq) (*userpb.Resp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &userpb.Resp{
			Code: 403,
			Msg:  token_error,
		}, nil
	}

	err := mysql.ModifyEmail(ctx, in.Username, in.NewEmail)
	if err != nil {
		return &userpb.Resp{
			Code: 200,
			Msg:  modify_error,
		}, nil
	}

	return &userpb.Resp{
		Code: 200,
		Msg:  modify_success,
	}, nil
}

func (s *UserServerImpl) ModifyPassword(ctx context.Context, in *userpb.UserModifyPasswordReq) (*userpb.Resp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &userpb.Resp{
			Code: 403,
			Msg:  token_error,
		}, nil

	}

	err := mysql.ModifyPassword(ctx, in.Username, in.NewPassword)
	if err != nil {
		return &userpb.Resp{
			Code: 200,
			Msg:  modify_error,
		}, nil
	}

	return &userpb.Resp{
		Code: 200,
		Msg:  modify_success,
	}, nil
}
