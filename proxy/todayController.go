package main

import (
	"net/http"
)

/*
	todayController adheres to the http.Handler interface and looks for a `today.json`
	file written on the server. It is responsible for serving the file as a response to
	incoming http requests.
*/
type todayController struct {
}


// serves file
func (controller todayController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "today.json")
}