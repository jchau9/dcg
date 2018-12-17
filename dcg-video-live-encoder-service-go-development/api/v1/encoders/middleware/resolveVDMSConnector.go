package middleware

import (
	"net/http"
	"net/url"
)

// ResolveVDMSConnector ...
func ResolveVDMSConnector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form, _ := url.ParseQuery(r.URL.RawQuery)
		apiKey := "eAs5wYI6DTGDl4e9i4CiqNs9AtM3P7UeKhzhGmFw"
		userID := "8baebcb1115a4bb78fa90c40ae8d81aa"
		integrationURL := "https://services.uplynk.com"
		form.Add("apiKey", apiKey)
		form.Add("userID", userID)
		form.Add("integrationURL", integrationURL)
		r.URL.RawQuery = form.Encode()
		next.ServeHTTP(w, r)
	})
}
