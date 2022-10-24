package qrcode

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	pkg_api "wecome-ipad/util/http"
)

type qrcodeParams struct {
	Token string `json:"token"`
}

type QrcodeGenerated struct {
	Code string `json:"code"`
	Data struct {
		UUID        string `json:"uuid"`
		QrCode      string `json:"qr_code"`
		ExpiredTime string `json:"expired_time"`
		IsNewDevice bool   `json:"is_new_device"`
	} `json:"data"`
}

func QrcodeController(c *gin.Context) {

	var qrcodeParams qrcodeParams

	if err := c.BindJSON(&qrcodeParams); err != nil {
		c.JSON(http.StatusForbidden, "token miss")
		return
	}

	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.GetJson("/login/qr_code", qrcodeParams.Token)

	resp := &QrcodeGenerated{}
	_ = json.Unmarshal([]byte(ret), resp)

	c.JSON(http.StatusOK, resp)

}
