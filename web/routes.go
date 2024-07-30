package web

import (
	"net/http"
	"order_service/web/middlewares"
)

func InitRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	// mux.Handle(
	// 	"POST /register",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.Register),
	// 	),
	// )

}
