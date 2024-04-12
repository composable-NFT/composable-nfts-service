package api

import (
	"composable-nfts-service/logging"
	"composable-nfts-service/pkg/app"
	"composable-nfts-service/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TestGet
// @Summary TestGet
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /test_get [get]
func TestGet(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.Query("p1")
	password := c.Query("p2")

	logging.Debug("username:", username, "password:", password)
	//fmt.Println("username:", username, "password:", password)

	appG.Response(
		http.StatusOK,
		e.SUCCESS,
		map[string]string{
			"message": "success",
		},
	)
}
