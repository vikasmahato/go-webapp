package main

import (
"github.com/vikasmahato/golang-webapp-starter-kit/controllers"
"github.com/vikasmahato/golang-webapp-starter-kit/services/listener"
"github.com/vikasmahato/golang-webapp-starter-kit/services/middleware"
)

/*
Add routes here using AddRoute and AddRouteWithMiddleware.
*/

func setupRoutes(httpListener *listener.HTTPListenerService, appContext *middleware.AppContext) {
	httpListener.
		AddStaticRoute("/assets/", "./www/assets").
		AddRoute("/greeting", controllers.HelloWorld, "GET")
}
