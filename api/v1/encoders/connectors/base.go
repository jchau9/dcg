package connectors

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*BaseConnector is an interface for a base connector that supports any provider.
* Supports HTTPS requests to any URL. For any new provider you should create a new connector class extending this one.
 */
type BaseConnector struct {
	Timeout int
}

/*Config is an interface for a request configuration object*/
type Config struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Payload string            `json:"payload"`
}

// Request ...
func (b BaseConnector) Request(config Config) (string, error) {
	// TODO: more error checking
	request, err := http.NewRequest(config.Method, config.URL, strings.NewReader(config.Payload))

	if err != nil {
		log.Fatal(err)
	}

	// set headers['content-length'] = payload.length if headers['content-length'] has no value
	if config.Payload != "" && (config.Headers["content-length"] == "" || config.Headers["Content-Length"] == "") {
		config.Headers["content-length"] = strconv.Itoa(len(config.Payload))
		request.Header.Set("content-length", strconv.Itoa(len(config.Payload)))
	}

	request.Header.Set("Content-Type", config.Headers["Content-Type"])
	client := &http.Client{}
	res, err := client.Do(request)
	var data []byte
	if err != nil {
		fmt.Println("http request has failed with error: ", err)
	} else {
		data, _ = ioutil.ReadAll(res.Body)
	}
	defer res.Body.Close()
	return string(data), err
}
