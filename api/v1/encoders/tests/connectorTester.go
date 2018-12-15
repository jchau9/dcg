package tests

import (
	"fmt"
	"go_development/src/api/v1/encoders/connectors"
	"go_development/src/api/v1/encoders/helpers"
	"log"
)

// TestBaseConnector ...
func TestBaseConnector() string {
	connector := connectors.BaseConnector{Timeout: 3000}
	url := "https://www.example.com/"
	res, err := connector.Request(connectors.Config{
		URL:    url,
		Method: "GET",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return res
}

// TestVDMSConnector ...
func TestVDMSConnector() string {
	apiKey := "eAs5wYI6DTGDl4e9i4CiqNs9AtM3P7UeKhzhGmFw"
	userID := "8baebcb1115a4bb78fa90c40ae8d81aa"
	integrationURL := "https://services.uplynk.com"
	baseConnector := connectors.BaseConnector{Timeout: 3000}
	connector := connectors.VDMSConnector{BaseConnector: baseConnector, IntegrationURL: integrationURL, APIKey: apiKey, UserID: userID}
	res, err := helpers.Single(connector, helpers.RequestConfig{
		Path:   "/api2/channel/update",
		Method: "POST",
		Parameters: map[string]interface{}{
			"id":        "0378e046a93b488ebc57300fa1fb158a",
			"slicer_id": "slce199_fxd",
		},
		Integration: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return res
}
