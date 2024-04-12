package routers

import (
	_ "composable-nfts-service/docs"
	"composable-nfts-service/routers/api"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	testApi := r.Group("/test")
	{
		testApi.GET("/test_get", api.TestGet)
	}

	return r

}
