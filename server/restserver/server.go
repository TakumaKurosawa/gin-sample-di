package restserver

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"github.com/TakumaKurosawa/gin-sample/server"
)

type httpServer struct {
	ginEngine  *gin.Engine
	echoEngine *echo.Echo
}

func NewHttpServer(ginEngine *gin.Engine, echoEngine *echo.Echo) server.Server {
	return &httpServer{
		ginEngine:  ginEngine,
		echoEngine: echoEngine,
	}
}

func (h *httpServer) NewRouter() {
	h.ginEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (h *httpServer) Run() error {
	if err := h.ginEngine.Run(); err != nil {
		return err
	}

	if err := h.echoEngine.Start(":8080"); err != nil {
		return err
	}

	return nil
}
