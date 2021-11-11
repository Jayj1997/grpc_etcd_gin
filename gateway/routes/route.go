/*
 * @Author       : jayj
 * @Date         : 2021-11-05 13:58:49
 * @Description  :
 */
package routes

import (
	"gateway/middleware"
	"gateway/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadGin() *gin.Engine {

	g := gin.Default()

	g.Use(middleware.Cors())

	getRoute(g)

	g.NoRoute(func(ctx *gin.Context) {
		res.Error(ctx, http.StatusNotFound, res.UrlNotFound)
	})

	return g
}

func getRoute(g *gin.Engine) {

	v1 := g.Group("/v1")
	addPingRoute(v1)

}
