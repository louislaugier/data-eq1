package airparif

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/louislaugier/projet-data-4A-go/model"
	geojson "github.com/paulmach/go.geojson"
)

// ZonesGET gets all zones
func ZonesGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	file, err := os.Open("model/airparif/zones-request.xml")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	xml, _ := ioutil.ReadAll(file)
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	resp, err := http.Post("http://magellan.airparif.asso.fr/geoserver/WPS/wps?authkey="+APIkey, "text/xml", bytes.NewReader((xml)))
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Airparif Server Error"
		log.Println("Error with Airparif HTTP request:", err)
	}
	defer resp.Body.Close()
	var fc geojson.FeatureCollection
	json.NewDecoder(resp.Body).Decode(&fc)
	res.Data = append(res.Data, fc)
	res.Meta.Query = "Positions géographiques des zones de pollution à Paris (données en temps réel)"
	res.Meta.ResultCount = len(fc.Features)
	json.NewEncoder(w).Encode(res)
}
