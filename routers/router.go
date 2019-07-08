package routers

import (
	"ginapi/pkg/setting"
	"ginapi/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RUN_MODE)

	apiv1 := r.Group("/api/v1")
	{
		//apiv1.POST("/ticket", v1.AddTicket)
		//apiv1.DELETE("/ticket/:id", v1.DeleteTicket)
		//apiv1.PUT("/ticket/:id", v1.EditTicket)
		//apiv1.GET("/tickets", v1.GetTicket)

		// 标签
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTags)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 文章
		apiv1.GET("/articles",v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return r
}
