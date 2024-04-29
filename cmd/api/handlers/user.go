package handlers

import (
	"net/http"

	"github.com/dkZzzz/quality_hub/cmd/api/rpc"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/userpb"
	"github.com/gin-gonic/gin"
)

var (
	param_error             = "参数错误"
	service_discovery_error = "服务发现失败"
)

func toResponse(rsp *userpb.Resp) (int, gin.H) {
	if rsp.Token != nil {
		return int(rsp.Code), gin.H{
			"code": rsp.Code,
			"msg":  rsp.Msg,
			"data": gin.H{"token": rsp.Token},
		}
	}
	return int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Msg,
		"data": nil,
	}
}

// Login 登录
// 从GIN接受参数，发送给rpc服务
func Login(c *gin.Context) {
	var param LoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}
	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := userpb.UserLoginReq{
		Username: param.Username,
		Password: param.Password,
	}
	rsp, _ := rpc.Login(&req)
	c.JSON(toResponse(rsp))
}

// Register 注册
// 把注册的上下文发送给rpc服务
func Register(c *gin.Context) {
	var param RegisterParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	rsp, _ := rpc.Register(&userpb.UserRegisterReq{
		Username: param.Username,
		Password: param.Password,
		Email:    param.Email,
	})
	c.JSON(toResponse(rsp))
}

// Logout 退出登录
// 把退出登录的上下文发送给rpc服务
func Logout(c *gin.Context) {
	var param LogoutParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	rsp, _ := rpc.Logout(&userpb.UserLogoutReq{
		Username: param.Username,
		Token:    param.Token,
	})
	c.JSON(toResponse(rsp))
}

// ModifyUsername 修改用户名
// 把修改用户名的上下文发送给rpc服务
func ModifyUsername(c *gin.Context) {
	var param ModifyUsernameParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	rsp, _ := rpc.ModifyUsername(&userpb.UserModifyUsernameReq{
		Username:    param.Username,
		NewUsername: param.NewUsername,
		Token:       param.Token,
	})
	c.JSON(toResponse(rsp))
}

// ModifyEmail 修改邮箱
// 把修改邮箱的上下文发送给rpc服务
func ModifyEmail(c *gin.Context) {
	var param ModifyEmailParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	rsp, _ := rpc.ModifyEmail(&userpb.UserModifyEmailReq{
		Username: param.Username,
		NewEmail: param.NewEmail,
		Token:    param.Token,
	})
	c.JSON(toResponse(rsp))
}

// ModifyPassword 修改密码
// 把修改密码的上下文发送给rpc服务
func ModifyPassword(c *gin.Context) {
	var param ModifyPasswordParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("user")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	rsp, _ := rpc.ModifyPassword(&userpb.UserModifyPasswordReq{
		Username:    param.Username,
		Password:    param.OldPassword,
		NewPassword: param.NewPassword,
		Token:       param.Token,
	})
	c.JSON(toResponse(rsp))
}
