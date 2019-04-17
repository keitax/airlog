package main

import (
	"fmt"
	"github.com/keitax/airlog/internal/di"
)

func main() {
	dc := di.Container{}

	g := dc.Gin()
	port := dc.Config().Port

	if err := g.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
