package router

import (
	"github.com/gin-gonic/gin"
	"wecome-ipad/controller/check"
	"wecome-ipad/controller/qrcode"
)

func Router() *gin.Engine {

	r := gin.New()

	wecome := r.Group("/v1")
	{
		//获取二维码
		wecome.GET("/qrcode", qrcode.QrcodeController)

		//检查用户是否登陆
		wecome.GET("/login/check", check.Check)
	}

	return r
}
