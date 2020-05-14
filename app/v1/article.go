package v1

import (
	"fmt"
	"hello_gin/app/global/response"
	"hello_gin/app/services"

	"github.com/gin-gonic/gin"
)

const (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	USER_HEADER_BUCKET   string = "qm-plus-img"
)

// 新增文章
func Add(c *gin.Context) {
	err, article := services.AddArticle("this is title", "this is title")
	if err != nil {
		response.FailWithDetailed(0, "error", fmt.Sprintf("新增失败，%v", err), c)
	} else {
		response.OkWithData(article, c)
		// response.OkWithMessage("成功", c)
	}
}
