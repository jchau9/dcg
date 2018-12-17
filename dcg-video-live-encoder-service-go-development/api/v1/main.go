package version1

import (
	"go_development/src/api/v1/encoders"

	"github.com/gorilla/mux"
)

// Version1 ...
func Version1(router *mux.Router) {
	encoders.EncoderRoutes(router)
}
