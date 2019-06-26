package v1

import (
	"ginapi/models"
	"ginapi/pkg/e"
	"ginapi/pkg/setting"
	"ginapi/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
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
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdAt := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("最长 100 字符")

}

////新增文章标签
//func AddTag(c *gin.Context) {
//	name := c.Query("name")
//	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
//	createdBy := c.Query("created_by")
//
//	valid := validation.Validation{}
//	valid.Required(name, "name").Message("名称不能为空")
//	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
//	valid.Required(createdBy, "created_by").Message("创建人不能为空")
//	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
//	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//
//	code := e.INVALID_PARAMS
//	if ! valid.HasErrors() {
//		if ! models.ExistTagByName(name) {
//			code = e.SUCCESS
//			models.AddTag(name, state, createdBy)
//		} else {
//			code = e.ERROR_EXIST_TAG
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code" : code,
//		"msg" : e.GetMsg(code),
//		"data" : make(map[string]string),
//	})
//}


