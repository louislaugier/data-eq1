package ratp

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
	geojson "github.com/paulmach/go.geojson"
)

type station struct {
	ID      int             `json:"identifiant,omitempty"`
	Feature geojson.Feature `json:"feature,omitempty"`
	// InfluxIndice  int             `json:"influx_indice"`
	Name          string `json:"name,omitempty"`
	Address       string `json:"address,omitempty"`
	Coordinates   string `json:"coordinates,omitempty"`
	TrainOutside  bool   `json:"train_outside,omitempty"`
	TransportType string `json:"transport_type,omitempty"`
	Line          string `json:"line,omitempty"`
	LineIndex     int    `json:"line_index,omitempty"`
}

type query struct {
	Limit   string `json:"limit"`
	Offset  string `json:"offset"`
	Outside string `json:"outside"`
	Type    string `json:"type"`
	Line    string `json:"line"`
}

// func scrapeGoogleMapsInflux(station station) int {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("www.google.com"),
// 		colly.IgnoreRobotsTxt(),
// 	)
// 	c.OnHTML(".section-result", func(e *colly.HTMLElement) {
// 		log.Println(e.ChildText("h3.section-result-title>span"))
// 	})
// 	c.OnError(func(r *colly.Response, err error) {
// 		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})
// 	c.Visit("https://www.google.com/maps/search/" + station.Name + "+station+de+" + station.TransportType)
// 	return 0
// }

func (s *station) stringToGeoJSONFeature() {
	s.Feature.Geometry = &geojson.Geometry{}
	s.Feature.Properties = map[string]interface{}{
		"identifiant":       s.ID,
		"nom":               s.Name,
		"adresse":           s.Address,
		"stationExterieure": s.TrainOutside,
		"typeTransport":     s.TransportType,
		"ligne":             s.Line,
	}
	s.Feature.Type = "Feature"
	s.Feature.Geometry.Type = "Point"
	lng, _ := strconv.ParseFloat(strings.Split(s.Coordinates, ",")[0], 64)
	lat, _ := strconv.ParseFloat(strings.Split(s.Coordinates, ",")[1], 64)
	s.Feature.Geometry.Point = []float64{lat, lng}
}

func (q *query) standardizeSQLQuery(values url.Values) string {
	q.Limit = "25338"
	if values.Get("type") == "RER" {
		q.Limit = "94"
	}
	if values.Get("type") == "métro" {
		q.Limit = "387"
	}
	if values.Get("type") == "bus" {
		q.Limit = "24381"
	}
	if values.Get("limit") != "" {
		q.Limit = values.Get("limit")
	}
	q.Offset = "0"
	if values.Get("offset") != "" {
		q.Offset = values.Get("offset")
	}
	var outside string
	if values.Get("outside") != "" {
		outside = "train_outside = '" + values.Get("outside") + "' AND "
		q.Outside = values.Get("outside")
	}
	var transportType string
	if values.Get("type") != "" {
		transportType = "transport_type = '" + values.Get("type") + "' AND "
		q.Type = values.Get("type")
	}
	if values.Get("type") == "train" {
		q.Limit = "481"
		transportType = "transport_type IN ('métro','RER') AND "
	}
	var line string
	if values.Get("line") != "" {
		line = "line = '" + values.Get("line") + "' AND "
		q.Line = values.Get("line")
	}
	if q.Outside != "" || q.Type != "" || q.Line != "" {
		return strings.TrimSuffix("WHERE "+outside+transportType+line, " AND ")
	}
	return ""
}

// StationsGET gets all stations
func StationsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var q query
	u := r.URL.Query()
	s := q.standardizeSQLQuery(u)
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	rows, err := database.Db.Query("SELECT * FROM (SELECT DISTINCT ON (coordinates, line) id, name, address, coordinates, train_outside, transport_type, line, line_index FROM ratp_stations " + s + " LIMIT " + q.Limit + " OFFSET " + q.Offset + ") AS stations ORDER BY line_index;")
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	for rows.Next() {
		var station station
		rows.Scan(&station.ID, &station.Name, &station.Address, &station.Coordinates, &station.TrainOutside, &station.TransportType, &station.Line, &station.LineIndex)
		station.stringToGeoJSONFeature()
		res.Data = append(res.Data, station)
	}
	defer rows.Close()
	res.Meta.Query = q
	res.Meta.ResultCount = len(res.Data)
	// station := &station{
	// 	Name:          "Pyramides",
	// 	TransportType: "métro",
	// }
	// station.InfluxIndice = scrapeGoogleMapsInflux(*station)
	json.NewEncoder(w).Encode(res)
}

// StationGET gets a station by id
func StationGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	rows, err := database.Db.Query("SELECT id, name, address, coordinates, train_outside, transport_type, line, line_index FROM ratp_stations WHERE id = '" + ID + "' LIMIT 1;")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	var station station
	for rows.Next() {
		rows.Scan(&station.ID, &station.Name, &station.Address, &station.Coordinates, &station.TrainOutside, &station.TransportType, &station.Line, &station.LineIndex)
	}
	defer rows.Close()
	if station.ID != 0 {
		res.Data = append(res.Data, station)
	}
	res.Meta.Query = ID
	res.Meta.ResultCount = len(res.Data)
	if res.Meta.ResultCount == 0 {
		res.Error = "Non-existing ID: " + ID
		res.Message = "Station not found"
	}
	json.NewEncoder(w).Encode(res)
}
