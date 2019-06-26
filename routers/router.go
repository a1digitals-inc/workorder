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
		apiv1.POST("/ticket", v1.AddTicket)
		apiv1.DELETE("/ticket/:id", v1.DeleteTicket)
		apiv1.PUT("/ticket/:id", v1.EditTicket)
		apiv1.GET("/tickets", v1.GetTicket)
	}

	return r
}
