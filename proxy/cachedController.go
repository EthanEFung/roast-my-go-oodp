package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
	cachedTodayController adheres to the http.Handler interface. It is responsible for calling
	the third party service for json data, and storing the response as a json file if the
	todayController has no file to serve or if the request is the first request of the day.
*/
type cachedTodayController struct {
	todayController todayController
}

func (s cachedTodayController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileInfo, err := os.Stat("today.json")
	if err != nil {
		err = fetchJSON()
	} else {
		now := time.Now()
		formatted := fileInfo.ModTime().Format("2006/01/02")
		midnight, _ := time.Parse("2006/01/02", formatted)
		if midnight.After(now) {
			err = fetchJSON()
		}
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("serving todays resource")
	s.todayController.ServeHTTP(w, r)
}

func fetchJSON() error {
	fmt.Println("fetching todays resource")
	res, err := http.Get("https://data.nba.net/data/10s/prod/v1/today.json")
	if err != nil {
		return err
	}
	file, err := os.Create("today.json")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err !=  nil {
		return err
	}
	return nil
}
