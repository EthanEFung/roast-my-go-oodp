package main

import (
	"log"
	"net/http"
)

func main() {
	ctrl := todayController{}
	cacheCtrl := cachedTodayController{ctrl}
	http.Handle("/", cacheCtrl)
	log.Fatal(http.ListenAndServe(":3000", nil))
}