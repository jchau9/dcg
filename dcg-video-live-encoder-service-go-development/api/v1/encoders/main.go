package encoders

import (
	"go_development/src/api/v1/encoders/controllers"
	"go_development/src/api/v1/encoders/middleware"

	"github.com/gorilla/mux"
)

// EncoderRoutes ...
func EncoderRoutes(router *mux.Router) {
	// middleware
	router.Use(middleware.EncoderDataInjector, middleware.ResolveVDMSConnector)

	// setPublicFeedSource
	router.HandleFunc("/setPublicFeedSource", controllers.SetPublicFeedSource).Methods("POST")
}
