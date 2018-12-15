package helpers

import (
	"go_development/src/api/v1/encoders/connectors"
)

// RequestConfig ...
type RequestConfig struct {
	EncoderName string
	Integration bool
	EncoderURL  string
	Path        string
	Method      string
	Parameters  map[string]interface{}
}

// Single ... return encoder response, error
func Single(connector connectors.VDMSConnector, config RequestConfig) (string, error) {
	res, err := connector.Request(connectors.VDMSConfig{
		EncoderURL:  config.EncoderURL,
		Path:        config.Path,
		Method:      config.Method,
		Integration: config.Integration,
		Parameters:  config.Parameters,
	})
	return res, err
}
