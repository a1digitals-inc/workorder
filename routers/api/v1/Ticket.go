package v1

import (
	"ginapi/models"
	"ginapi/pkg/e"
	"ginapi/pkg/setting"
	"ginapi/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
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

func AddTag(c *gin.Context)  {
	name :=c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustUint8()
	createdAt := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("最长 100 字符")
	valid.Required(createdAt, "created_at").Message("时间不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只能为0 和 1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTicketByName(name) {
			models.AddTicket(name,state,createdAt)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":e.GetMssg(code),
		"data":make(map[string]string),
	})
}


