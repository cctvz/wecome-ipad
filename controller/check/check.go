package check

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	pkg_api "wecome-ipad/util/http"
)

type CheckGenerated struct {
	Code int `json:"code"`
	Data struct {
		WxID         string `json:"wx_id"`
		HeadImageURL string `json:"head_image_url"`
		NickName     string `json:"nick_name"`
		Status       string `json:"status"`
	} `json:"data"`
}

func Check(c *gin.Context) {

	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.GetJson("/login/check", map[string]string{})

	respCheck := &CheckGenerated{}
	_ = json.Unmarshal([]byte(ret), respCheck)
	c.JSON(http.StatusOK, respCheck)
}
