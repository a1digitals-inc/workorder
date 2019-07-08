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


func GetArticle(ctx *gin.Context)  {

	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := 0
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistTagByTagId(id) {
			data = models.GetArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		code = e.INVALID_PARAMS
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMssg(code),
		"data" : data,
	})
}

func GetArticles(c *gin.Context)  {

}

func AddArticle(c *gin.Context)  {

}

func EditArticle(c *gin.Context)  {

}

func DeleteArticle(c *gin.Context)  {

}