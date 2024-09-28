package dto

import (
	"gin_demo/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,valid_username"`
	PassWord string `json:"password" form:"password" comment:"管理员密码" example:"123" validate:"required,valid_password"`
	Age      int    `json:"age" form:"age" comment:"管理员年龄" example:"20" validate:"required,valid_age"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminInfoOutput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
