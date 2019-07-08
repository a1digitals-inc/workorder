package v1

import (
	"ginapi/models"
	"ginapi/pkg/e"
	"ginapi/pkg/setting"
	"ginapi/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 增加工单
func AddTicket(c *gin.Context) {

}

// 删除工单
func DeleteTicket(c *gin.Context) {

}

// 修改工单
func EditTicket(c *gin.Context) {

}

// 工单详情
func GetTicket(c *gin.Context) {
	id := c.Query("id")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if id != "" {
		maps["id"] = id
	}

	if arg := c.Query("status"); arg != "" {
		maps["status"] = com.StrTo(arg).Exist()
	}
	data["list"] = models.GetTickets(util.GetPage(c), setting.PAGE_SIZE, maps)
	data["total"] = models.GetTotalTickets(maps)
	c.JSON(http.StatusOK, gin.H{
		"code":e.SUCCESS,
		"msg":e.GetMssg(e.SUCCESS),
		"data" : data,
	})
}




