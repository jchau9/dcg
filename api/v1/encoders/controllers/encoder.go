package controllers

import (
	"fmt"
	"go_development/src/api/v1/encoders/connectors"
	"go_development/src/api/v1/encoders/helpers"
	"go_development/src/response"
	"net/http"
)

func initConnector(r *http.Request) connectors.VDMSConnector {
	apiKey := r.FormValue("apiKey")
	userID := r.FormValue("userID")
	integrationURL := r.FormValue("integrationURL")

	baseConnector := connectors.BaseConnector{Timeout: 3000}
	connector := connectors.VDMSConnector{BaseConnector: baseConnector, IntegrationURL: integrationURL, APIKey: apiKey, UserID: userID}
	return connector
}

// SetPublicFeedSource ...
func SetPublicFeedSource(w http.ResponseWriter, r *http.Request) {
	channelGUID := r.FormValue("channelGUID")
	encoderID := r.FormValue("encoderId")
	connector := initConnector(r)

	res, _ := helpers.Single(connector, helpers.RequestConfig{
		Path:   "/api2/channel/update",
		Method: "POST",
		Parameters: map[string]interface{}{
			"id":        channelGUID,
			"slicer_id": encoderID,
		},
		Integration: true,
	})

	fmt.Println(res)
	response.HandleError(w, r, "placeholder") //placeholder
}
