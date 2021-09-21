package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type Server struct {
	S  *gin.Engine
	DB *pg.DB
}
