package main

import (
	"encoding/json"
	"net/http"
)

type geoData struct {
	Country string
	Region  string
	City    string
}

func findGeoData(w http.ResponseWriter, r *http.Request) {
	var data geoData
	data.Country = r.Header.Get("X-AppEngine-Country")
	data.Region = r.Header.Get("X-AppEngine-Region")
	data.City = r.Header.Get("X-AppEngine-City")

	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
