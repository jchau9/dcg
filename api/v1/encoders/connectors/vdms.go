package connectors

import (
	"bytes"
	"compress/zlib"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// VDMSConnector ...
type VDMSConnector struct {
	BaseConnector  BaseConnector
	IntegrationURL string
	APIKey         string
	UserID         string
}

// VDMSConfig ...
type VDMSConfig struct {
	EncoderURL  string
	Path        string
	Method      string
	Parameters  map[string]interface{}
	Integration bool
}

// generate signature to make appropriate request
func (v VDMSConnector) getSignedPayload(method string, parameters map[string]interface{}, isIntegration bool) string {
	// get current time stamp
	curTstamp := int(time.Now().Unix())
	// check if integration api, if it is, connect to vdms
	if !isIntegration {
		cnonce := rand.Intn(1000)
		parameters["cnonce"] = cnonce
		parameters["timestamp"] = curTstamp / cnonce
		// authentication of request
		h := sha1.New()
		h.Write([]byte(v.APIKey))
		authToken := hex.EncodeToString(h.Sum(nil))
		// generate signature
		sigInput := method + ":" + parameters["timestamp"].(string) + ":" + parameters["cnonce"].(string) + ":" + authToken

		h2 := sha1.New()
		h2.Write([]byte(sigInput))
		parameters["sig"] = base64.StdEncoding.EncodeToString(h2.Sum(nil))

		b, _ := json.Marshal(parameters)
		return string(b)
	}
	parameters["_owner"] = v.UserID
	parameters["_timestamp"] = curTstamp

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	b, _ := json.Marshal(parameters)
	w.Write(b)
	w.Close()
	msg := strings.Replace(base64.StdEncoding.EncodeToString(buf.Bytes()), " ", "", -1)

	mac := hmac.New(sha256.New, []byte(v.APIKey))
	mac.Write([]byte(msg))
	sig := hex.EncodeToString(mac.Sum(nil))
	return "msg=" + url.QueryEscape(msg) + "&sig=" + url.QueryEscape(sig)
}

// Request ...
func (v VDMSConnector) Request(config VDMSConfig) (string, error) {
	var reqURL *url.URL
	var contentType string
	if config.Integration {
		reqURL, _ = url.Parse(v.IntegrationURL)
		contentType = "application/x-www-form-urlencoded"
	} else {
		reqURL, _ = url.Parse(config.EncoderURL)
		contentType = "application/json"
	}
	reqURL.Path = config.Path

	payload := v.getSignedPayload(config.Method, config.Parameters, config.Integration)
	res, err := v.BaseConnector.Request(Config{
		URL:    reqURL.String(),
		Method: config.Method,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
		Payload: payload,
	})
	return res, err
}
