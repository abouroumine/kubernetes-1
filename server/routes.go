package server

import (
	ctrl "abouroumine.com/kubernetes-1/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) NewRouter() {
	gin.SetMode(gin.ReleaseMode)

	s.S = gin.New()

	s.S.Use(gin.Logger())
	s.S.Use(gin.Recovery())

	hello := new(ctrl.HelloController)
	s.S.GET("/hello", hello.Hello)
	s.S.POST("/word", hello.ReceiveWorld)

	s.S.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not Found",
		})
	})
}
