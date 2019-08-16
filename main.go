package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keitam913/airlog/di"
)

func main() {
	dc := di.Container{}

	g := dc.Gin()
	g.Use(gin.Recovery(), gin.Logger())

	port := dc.Config().Port

	if err := g.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
