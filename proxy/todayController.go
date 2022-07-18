package main

import (
	"net/http"
)

type todayController struct {
	
}

// serves file
func (controller todayController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "today.json")
}