package routers

import (
	"composable-nfts-service/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	testApi := r.Group("/test")
	{
		testApi.GET("/test_get", api.TestGet)
	}

	return r

}
