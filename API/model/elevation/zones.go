package elevation

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/louislaugier/projet-data-4A-go/model"
	geojson "github.com/paulmach/go.geojson"
)

// ZonesGET gets elevation data for each arrondissement
func ZonesGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	var fc geojson.FeatureCollection
	for i := 1; i < 21; i++ {
		f, err := arrondissementElevation(i)
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Google Maps API Error"
			log.Println("Error with HTTP request:", err)
		}
		fc.Features = append(fc.Features, &f)
	}
	res.Data = append(res.Data, fc)
	res.Meta.Query = "Elévation moyenne du paysage des arrondissements de Paris en mètres (données en temps réel)"
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}
