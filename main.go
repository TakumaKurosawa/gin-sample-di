package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"github.com/TakumaKurosawa/gin-sample/server/restserver"
)

func main() {
	g := gin.Default()
	e := echo.New()
	srv := restserver.NewHttpServer(g, e)
	srv.NewRouter()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
