package elevation

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"data-eq1/API/model"

	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
)

type googleMapsElevationJSONData struct {
	Results []googleMapsElevation `json:"results"`
	Status  string                `json:"status"`
}

type googleMapsElevation struct {
	Elevation  float64     `json:"elevation"`
	Location   interface{} `json:"location"`
	Resolution float64     `json:"resolution"`
}

type data struct {
	Elevation [20]interface{}
}

func arrondissementElevation(arrondissementNumber int) (geojson.Feature, error) {
	var coords string
	switch arrondissementNumber {
	case 1:
		coords = "48.863277,2.335506"
	case 2:
		coords = "48.868636,2.343533"
	case 3:
		coords = "48.863147,2.360585"
	case 4:
		coords = "48.854791,2.357551"
	case 5:
		coords = "48.844652,2.351118"
	case 6:
		coords = "48.849347,2.334118"
	case 7:
		coords = "48.858203,2.312966"
	case 8:
		coords = "48.873479,2.314293"
	case 9:
		coords = "48.877998,2.339028"
	case 10:
		coords = "48.876907,2.361253"
	case 11:
		coords = "48.859504,2.380259"
	case 12:
		coords = "48.840961,2.391264"
	case 13:
		coords = "48.829500,2.363722"
	case 14:
		coords = "48.830312,2.327306"
	case 15:
		coords = "48.841094,2.292716"
	case 16:
		coords = "48.850301,2.269222"
	case 17:
		coords = "48.887935,2.309186"
	case 18:
		coords = "48.893369,2.348956"
	case 19:
		coords = "48.887684,2.384420"
	case 20:
		coords = "48.863855,2.403118"
	}
	var f geojson.Feature
	result, err := http.Get("https://maps.googleapis.com/maps/api/elevation/json?locations=" + coords + "&key=" + APIkey)
	if err == nil {
		f.Geometry = &geojson.Geometry{}
		f.Type = "Point"
		f.Geometry.Type = "Point"
		x, _ := strconv.ParseFloat(strings.Split(coords, ",")[0], 64)
		y, _ := strconv.ParseFloat(strings.Split(coords, ",")[1], 64)
		f.Geometry.Point = []float64{y, x}
		readResult, _ := ioutil.ReadAll(result.Body)
		defer result.Body.Close()
		var g googleMapsElevationJSONData
		json.Unmarshal(readResult, &g)
		f.Properties = map[string]interface{}{
			"elevation": g.Results[0].Elevation,
		}
		return f, nil
	}
	return f, err
}

// DataGET gets elevation data for all arrondissements
func DataGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	var d data
	for i := 1; i < 21; i++ {
		f, err := arrondissementElevation(i)
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Google Maps API Error"
			log.Println("Error with HTTP request:", err)
		}
		d.Elevation[i-1] = f.Properties["elevation"]
	}
	res.Data = append(res.Data, d)
	res.Meta.Query = "Elévation moyenne du paysage des arrondissements de Paris en mètres (données en temps réel)"
	res.Meta.ResultCount = len(d.Elevation)
	json.NewEncoder(w).Encode(res)
}

// ArrondissementGET gets elevation data for an arrondissement
func ArrondissementGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	arr, _ := mux.Vars(r)["arrondissement"]
	if arr != "all" {
		a, _ := strconv.Atoi(arr)
		f, err := arrondissementElevation(a)
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Google Maps API Error"
			log.Println("Error with HTTP request:", err)
		}
		res.Data = append(res.Data, f.Properties["elevation"])
	} else {
		result, _ := http.Get("http://localhost:80/api/indices-elevation")
		readResult, _ := ioutil.ReadAll(result.Body)
		var data struct {
			Data []struct {
				Elevation []float64 `json:"Elevation"`
			} `json:"data"`
		}
		json.Unmarshal(readResult, &data)
		var total float64 = 0
		for _, value := range data.Data[0].Elevation {
			total += value
		}
		res.Data = append(res.Data, total/float64(len(data.Data[0].Elevation)))
		res.Meta.Query = "Elévation moyenne du paysage de la ville de Paris en mètres (données en temps réel)"
	}
	if arr == "1" && arr != "all" {
		res.Meta.Query = "Elévation moyenne du paysage du " + arr + "er arrondissement de Paris en mètres (données en temps réel)"
	} else if arr != "1" && arr != "all" {
		res.Meta.Query = "Elévation moyenne du paysage du " + arr + "ème arrondissement de Paris en mètres (données en temps réel)"
	}
	res.Meta.ResultCount = 1
	json.NewEncoder(w).Encode(res)
}
