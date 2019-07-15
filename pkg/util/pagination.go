package util

import (
	"ginapi/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// trans pagination to start number
func GetPage(c *gin.Context) (result int) {
	result = 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PAGE_SIZE
	}
	return
}
