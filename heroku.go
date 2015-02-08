// +build !appengine

package main

import (
	"net/http"
)

// For Non-AppEngine Environments
func main() {
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
