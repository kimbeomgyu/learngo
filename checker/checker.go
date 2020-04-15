package checker

import (
	"net/http"
)

// CheckErr is error check
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckStatusCode is status code check
func CheckStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		panic("Request failed with Status: " + string(res.StatusCode))
	}
}
