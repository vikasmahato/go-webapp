package main

import (
	"github.com/vikasmahato/golang-webapp-starter-kit/services/listener"
	"github.com/vikasmahato/golang-webapp-starter-kit/services/middleware"
)

func setupMiddleware(httpListener *listener.HTTPListenerService, appContext *middleware.AppContext) {
	httpListener.
		AddMiddleware(appContext.Logger).
		AddMiddleware(appContext.StartAppContext).
		AddMiddleware(appContext.AccessControl).
		AddMiddleware(appContext.OptionsHandler)
}