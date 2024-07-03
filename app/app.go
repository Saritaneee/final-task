package app

type UserLogin struct {
	Name     string `form:"name"`
	Username string `form:"username"`
	Email    string `form:"email" valid:"required"`
	Password string `form:"password" valid:"required"`
}

type UserUpdate struct {
	Name     string `form:"name" valid:"required"`
	Username string `form:"username" valid:"required"`
	Email    string `form:"email" valid:"required"`
	Password string `form:"password" valid:"required"`
}
