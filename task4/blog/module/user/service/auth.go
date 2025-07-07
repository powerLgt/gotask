package service

import (
	"blog/module/common/model"
	"blog/module/user/form"
	"blog/pkg/database"
	"blog/pkg/utils"
	"errors"
)

type AuthService struct{}

// 用户登录
func (s *AuthService) Login(form form.LoginForm) (user model.User, token string, err error) {
	// 查找用户
	err = database.DB.Where("username = ?", form.Username).First(&user).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	// 验证密码
	if !utils.CheckPassword(user.Password, form.Password) {
		err = errors.New("用户名或密码错误")
		return
	}

	// 生成JWT令牌
	token, err = utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		err = errors.New("生成令牌失败")
		return
	}

	return
}

// 用户注册
func (s *AuthService) Register(form form.RegisterForm) (user model.User, err error) {
	// 检查用户名是否已存在
	var existingUser model.User
	if database.DB.Where("username = ? OR email = ?", form.Username, form.Email).First(&existingUser).Error == nil {
		err = errors.New("用户名或邮箱已存在")
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(form.Password)
	if err != nil {
		err = errors.New("密码加密失败")
		return
	}

	// 创建用户
	user = model.User{
		Username: form.Username,
		Password: hashedPassword,
		Email:    form.Email,
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		err = errors.New("用户创建失败")
		return
	}

	return
}
