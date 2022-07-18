package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

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
	s.todayController.ServeHTTP(w, r)
}

func fetchJSON() error {
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
