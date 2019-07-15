package v1

import (
	"ginapi/models"
	"ginapi/pkg/e"
	"ginapi/pkg/setting"
	"ginapi/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func GetArticle(ctx *gin.Context)  {
	id := com.StrTo(ctx.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.SUCCESS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : data,
	})
}

func GetArticles(ctx *gin.Context)  {
	valid := validation.Validation{}
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if arg:= ctx.Query("state"); arg != "" {
		maps["state"] = com.StrTo(arg).MustInt()
		valid.Range(maps["state"], 0, 1, "state").Message("状态只能为 0 或者 1")
	}

	if arg := ctx.Query("tag_id"); arg != "" {
		maps["tag_id"] = com.StrTo(arg).MustInt()
		valid.Min(maps["tag_id"], 1, "tag_id").Message("标签ID必须大于0")
	}
	code := e.SUCCESS
	if !valid.HasErrors() {
		data["lists"] = models.GetArticles(util.GetPage(ctx), setting.PAGE_SIZE, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Println(err)
		}
		code = e.INVALID_PARAMS
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg" : e.GetMssg(code),
		"data":data,
	})
}

func AddArticle(ctx *gin.Context)  {
	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	createdBy := ctx.Query("created_by")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc")
	valid.Required(content, "content")
	valid.Required(createdBy, "createdBy")
	valid.Range(state, 0 ,1, "state")

	code := e.SUCCESS
	if ! valid.HasErrors() {
		if models.ExistTagByTagId(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.INVALID_PARAMS
		for _,err := range valid.Errors {
			log.Println(err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : make(map[string]interface{}),
	})
}

func EditArticle(ctx *gin.Context)  {
	request := validateForEdit(ctx)
	code := e.SUCCESS
	if models.ExistArticleById(request["id"].(int)) {
		if models.ExistTagByTagId(request["tagId"].(int)) {
			data := make(map[string]interface{})
			data["tagId"] = request["tagId"].(int)
			if request["title"] != "" {
				data["title"] = request["title"].(string)
			}
			if request["desc"] != "" {
				data["desc"] = request["desc"].(string)
			}
			if request["content"] != "" {
				data["content"] = request["content"].(string)
			}

			data["modified_by"] = request["modifiedBy"].(string)
			models.EditArticle(request["id"].(int), data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : make(map[string]interface{}),
	})
}

func validateForEdit(ctx *gin.Context)(req map[string]interface{})  {
	req = make(map[string]interface{})
	valid := validation.Validation{}
	req["id"] = com.StrTo(ctx.Param("id")).MustInt()
	req["tagId"] = com.StrTo(ctx.Query("tag_id")).MustInt()
	req["title"] = ctx.Query("title")
	req["desc"] = ctx.Query("desc")
	req["content"] = ctx.Query("content")
	req["modifiedBy"] = ctx.Query("modified_by")

	valid.Min(req["id"] , 1, "id").Message("ID 必须大于 0")
	valid.MaxSize(req["title"] , 100, "title").Message("标题最长 100 字符")
	valid.MaxSize(req["content"] , 65535, "content").Message("内容最长 65535")
	valid.MaxSize(req["modifiedBy"] , 100, "modified_by").Message("修改人最长 100 字符")

	if arg:= ctx.Query("state"); arg != "" {
		req["state"] = com.StrTo(arg).MustInt()
		valid.Range(req["state"],0, 1,"state").Message("状态只允许0或1")
	}

	if valid.HasErrors() {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : e.INVALID_PARAMS,
			"msg" : e.GetMssg(e.INVALID_PARAMS),
			"data" : make(map[string]interface{}),
		})
		ctx.Abort()
	}
	return
}

func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			log.Println(err)
		}
	} else {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : make(map[string]interface{}),
	})
}