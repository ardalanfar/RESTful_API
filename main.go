package main

import (
	"Service_Restful/service"
)

func main() {
	app := &service.App{}

	app.Initialize()
	app.Run(":9027")
}
