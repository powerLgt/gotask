package form

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
	Username string `form:"username" binding:"required,min=3,max=50"`
	Password string `form:"password" binding:"required,min=6"`
	Email    string `form:"email" binding:"required,email,min=1,max=100"`
}
