package main

import (
	"order_service/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	app.Wait()
	app.Cleanup()
}
