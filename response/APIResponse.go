package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleError ...
func HandleError(w http.ResponseWriter, r *http.Request, error string) {
	finalize := func(error string, status int) {
		w.WriteHeader(status)
		data, err := json.Marshal(map[string]string{"errorCode": "0", "errorMessage": error})
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
	finalize(error, 500)
}

// Respond ...
func Respond(w http.ResponseWriter, r *http.Request, message string) {
	finalize := func(message string, status int) {
		// set status num to what message should be sent to
		w.WriteHeader(status)
		// send message
		data, err := json.Marshal(map[string]string{"status": "success", "message": message})
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
	finalize(message, 200)
}
