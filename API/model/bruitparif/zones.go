package bruitparif

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
	geojson "github.com/paulmach/go.geojson"
)

// ZonesGET gets all sound nuisance zones
func ZonesGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	rows, err := database.Db.Query("SELECT decibels, coordinates FROM nuisances_sonores WHERE date = 'day';")
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	var fc geojson.FeatureCollection
	for rows.Next() {
		var decibels float64
		var coords string
		rows.Scan(&decibels, &coords)
		x, _ := strconv.ParseFloat(strings.Split(coords, ", ")[0], 64)
		y, _ := strconv.ParseFloat(strings.Split(coords, ", ")[1], 64)
		var f geojson.Feature
		f.Type = "Point"
		f.Geometry = &geojson.Geometry{}
		f.Geometry.Type = "Point"
		f.Geometry.Point = []float64{y, x}
		f.Properties = map[string]interface{}{
			"decibels": decibels,
		}
		fc.Features = append(fc.Features, &f)
	}
	res.Data = append(res.Data, fc)
	res.Meta.Query = "Positions géographiques des zones de nuisances sonores à Paris (données historiques)"
	res.Meta.ResultCount = len(fc.Features)
	json.NewEncoder(w).Encode(res)
}
