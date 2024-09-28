package controller

import (
	"gin_demo/dto"
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func TestRegister(group *gin.RouterGroup) {
	test := &TestController{}
	group.GET("/ping", test.Ping)
	group.POST("/ping", test.PostTest)
}

// Ping
// @Summary ping test
// @Description ping test
// @Tags ping测试
// @ID /test/ping
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /test/ping [get]
func (test *TestController) Ping(c *gin.Context) {
	middleware.ResponseSuccess(c, "pong")
}

// PostTest
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /test/ping
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /test/ping [post]
func (test *TestController) PostTest(c *gin.Context) {
	input := &dto.AdminLoginInput{}
	if err := input.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
	}
	middleware.ResponseSuccess(c, &dto.AdminInfoOutput{
		Code:    200,
		Message: "success",
	})
}
