package main

import (
	"RESTful_API/service"
)

func main() {
	app := &service.App{}

	app.Initialize()
	app.Run(":9027")
}
