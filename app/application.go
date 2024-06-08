package app

import (
	"order_service/config"
	"order_service/db"
	"order_service/web"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	config.LoadConfig()
	db.InitDB()
}

func (app *Application) Run() {
	web.StartServer(&app.wg)
}

func (app *Application) Wait() {
	app.wg.Wait()
}

func (app *Application) Cleanup() {
	db.CloseDB()
}
