package helpers

import "net/http"

// GetBearerToken retrieves the bearer token from the request header
func GetBearerToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && bearerToken[:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}
