package main

import (
	"fmt"
	"github.com/keitax/airlog/internal/di"
)

func main() {
	con := di.Container{}
	fmt.Println(con)
}
