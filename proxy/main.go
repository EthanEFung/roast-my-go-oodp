package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		in this example imagine that we have a http server with an endpoint that serves NBA
		data in the form of JSON. But lets assume that we are using a 3rd party service, and
		need to rate limit our calls to said service. A typical way we can deal with this matter
		is through caching the response.  
	*/
	ctrl := todayController{}
	/*
		by using the cacheTodayController, what we'll do is save the response for the day as a file
		that way each day we will make only one call to our 3rd party service.
	*/
	cacheCtrl := cachedTodayController{ctrl}
	http.Handle("/", cacheCtrl)
	log.Fatal(http.ListenAndServe(":3000", nil))
}