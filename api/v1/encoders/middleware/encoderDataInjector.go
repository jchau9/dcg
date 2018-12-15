package middleware

import (
	"net/http"
	"net/url"
)

// EncoderDataInjector ...
func EncoderDataInjector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form, _ := url.ParseQuery(r.URL.RawQuery)
		channelGUID := "0378e046a93b488ebc57300fa1fb158a"
		encoderID := "slce199_fxd"
		form.Add("channelGUID", channelGUID)
		form.Add("encoderId", encoderID)
		r.URL.RawQuery = form.Encode()
		next.ServeHTTP(w, r)
	})
}
