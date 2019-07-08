package v1

import (
	"ginapi/models"
	"ginapi/pkg/e"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetTags(ctx *gin.Context)  {

}

func AddTag(ctx *gin.Context)  {

	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	createdBy := ctx.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required("createdBy", "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0,1, "state").Message("状态只允许0或1")

	var code int

	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		code = e.INVALID_PARAMS
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":e.GetMssg(code),
		"data":make(map[string]string),
	})
}

func EditTags(ctx *gin.Context)  {
	id := com.StrTo(ctx.Param("id")).MustInt()
	name := ctx.Query("name")
	modifedBy := ctx.Query("modified_by")

	log.Println(modifedBy)

	valid := validation.Validation{}
	valid.Required(id,"id").Message("ID不能为空")
	valid.Required(modifedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifedBy,100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	state := -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	log.Println(valid.Errors)
	var code int
	if !valid.HasErrors() {
		if models.ExistTagByTagId(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifedBy

			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.INVALID_PARAMS
	}
	// TODO  make(map[string]string), 和 make(map[string]interface{}) 有什么区别
	ctx.JSON(http.StatusOK, gin.H{
		"code" :code,
		"msg" : e.GetMssg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := -1
	log.Println("id",id)
	if !valid.HasErrors() {
		if models.ExistTagByTagId(id) {
			models.DeleteTag(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}

	} else {
		code = e.INVALID_PARAMS
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : make(map[string]string),
	})

}
