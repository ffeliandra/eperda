package main

import (
	"github.com/ffeliandra/eperdabe/controller"
	"github.com/ffeliandra/eperdabe/model"
)

func main() {
	model.ConnectDatabase()
	controller.RouteInit()
}
