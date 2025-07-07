package controller

import (
	"blog/module/common/vo"
	"blog/module/user/form"
	"blog/module/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authService = &service.AuthService{}

// 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
	var loginForm form.LoginForm

	// 绑定JSON数据
	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		ctx.JSON(http.StatusBadRequest, vo.Fail("请求参数错误: "+err.Error()))
		return
	}

	// 调用服务层进行登录验证
	user, token, err := authService.Login(loginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	// 返回成功结果和令牌
	data := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	ctx.JSON(http.StatusOK, vo.Success("登录成功", data))
}

// 用户注册
func (c *AuthController) Register(ctx *gin.Context) {
	var registerForm form.RegisterForm

	// 绑定JSON数据
	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		ctx.JSON(http.StatusBadRequest, vo.Fail("请求参数错误: "+err.Error()))
		return
	}

	// 调用服务层进行用户注册
	user, err := authService.Register(registerForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, vo.Success("注册成功", user))
}
