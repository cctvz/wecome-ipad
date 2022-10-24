package router

import (
	"github.com/gin-gonic/gin"
	"wecome-ipad/controller/qrcode"
)

func Router() *gin.Engine {

	r := gin.New()

	wecome := r.Group("/v1")
	{
		//获取二维码
		wecome.GET("/qrcode", qrcode.QrcodeController)
	}

	return r
}
