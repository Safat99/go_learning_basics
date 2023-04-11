package router

import (
	"mongoGo/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/movies/all", controller.GetAllMyMovies)
	router.POST("/movies/create", controller.CreateMovie)
	router.PUT("/movies/:id", controller.MarkedAsWatched)
	router.DELETE("/movies/:id", controller.DeleteMovie)

	return router
}
