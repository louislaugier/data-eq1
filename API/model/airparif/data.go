package airparif

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"data-eq1/API/database"
	"data-eq1/API/model"

	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
)

type record struct {
	O3   float64
	NO2  float64
	PM10 float64
	PM25 float64
}

type data struct {
	O3   []float64
	NO2  []float64
	PM10 []float64
	PM25 []float64
}

// ReverseValues export for maratp package
func (d data) ReverseValues() {
	for i, j := 0, len(d.O3)-1; i < j; i, j = i+1, j-1 {
		d.O3[i], d.O3[j] = d.O3[j], d.O3[i]
	}
	for i, j := 0, len(d.NO2)-1; i < j; i, j = i+1, j-1 {
		d.NO2[i], d.NO2[j] = d.NO2[j], d.NO2[i]
	}
	for i, j := 0, len(d.PM10)-1; i < j; i, j = i+1, j-1 {
		d.PM10[i], d.PM10[j] = d.PM10[j], d.PM10[i]
	}
	for i, j := 0, len(d.PM25)-1; i < j; i, j = i+1, j-1 {
		d.PM25[i], d.PM25[j] = d.PM25[j], d.PM25[i]
	}
}

// ArrondissementsGET gets air quality indice for all arrondissements
func ArrondissementsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	file, err := os.Open("model/airparif/arrondissements/all.xml")
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

// DataGET saves airparif data for current day and returns data from last 7 days for 1 arrondissement or averages for all of Paris
func DataGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	arr, _ := mux.Vars(r)["arrondissement"]
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	var d data
	if arr != "all" {
		file, err := os.Open("model/airparif/arrondissements/" + arr + ".xml")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		xml, _ := ioutil.ReadAll(file)
		resp, err := http.Post("http://magellan.airparif.asso.fr/geoserver/WPS/wps?authkey="+APIkey, "text/xml", bytes.NewReader((xml)))
		if err != nil {
			res.StatusCode = 500
			res.Message = "External Server Error"
			res.Error = "Airparif API Error"
			log.Println("Error with HTTP request:", err)
			log.Println(err)
		}
		defer resp.Body.Close()
		var fc geojson.FeatureCollection
		json.NewDecoder(resp.Body).Decode(&fc)
		if err == nil {
			rows, err := database.Db.Query("SELECT date FROM qualite_air WHERE arrondissement = " + arr + " ORDER BY date DESC LIMIT 1;")
			if err != nil {
				res.StatusCode = 500
				res.Message = "Internal Server Error"
				res.Error = "Database Error"
				log.Println("Error with transaction:", err)
				log.Println(err)
			}
			var date time.Time
			for rows.Next() {
				rows.Scan(&date)
			}
			if date.Format("2006-01-02") != time.Now().Format("2006-01-02") {
				txn, err := database.Db.Begin()
				if err != nil {
					log.Println(err)
				}
				txn.Exec("INSERT INTO qualite_air (arrondissement, date, o3, no2, pm10, pm25) VALUES ($1, $2, $3, $4, $5, $6);", arr, time.Now().Format("2006-01-02"), fc.Features[0].Properties["valeur_o3"], fc.Features[0].Properties["valeur_no2"], fc.Features[0].Properties["valeur_pm10"], fc.Features[0].Properties["valeur_pm25"])
				err = txn.Commit()
				if err != nil {
					res.StatusCode = 500
					res.Message = "Internal Server Error"
					res.Error = "Database Error"
					log.Println("Error with transaction:", err)
				}
			}
		}
		rows, err := database.Db.Query("SELECT o3, no2, pm10, pm25 FROM qualite_air WHERE arrondissement = " + arr + " ORDER BY date DESC LIMIT 7;")
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Database Error"
			log.Println("Error with transaction:", err)
		}
		for rows.Next() {
			var r record
			rows.Scan(&r.O3, &r.NO2, &r.PM10, &r.PM25)
			d.O3 = append(d.O3, r.O3)
			d.NO2 = append(d.NO2, r.NO2)
			d.PM10 = append(d.PM10, r.PM10)
			d.PM25 = append(d.PM25, r.PM25)
		}
		if arr == "1" {
			res.Meta.Query = "Moyennes en temps réel de la concentration en μg/m3 d'ozone, de dioxyde d'azote et de particules en suspension (PM10/25) dans le " + arr + "er arrondissement pour les 7 derniers jours"
		} else {
			res.Meta.Query = "Moyennes en temps réel de la concentration en μg/m3 d'ozone, de dioxyde d'azote et de particules en suspension (PM10/25) dans le " + arr + "ème arrondissement pour les 7 derniers jours"
		}
		defer rows.Close()
	} else {
		rows, err := database.Db.Query("SELECT o3, no2, pm10, pm25 FROM qualite_air ORDER BY date DESC, arrondissement LIMIT 140;")
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Database Error"
			log.Println("Error with transaction:", err)
		}
		var tempData data
		for rows.Next() {
			var r record
			rows.Scan(&r.O3, &r.NO2, &r.PM10, &r.PM25)
			tempData.O3 = append(tempData.O3, r.O3)
			tempData.NO2 = append(tempData.NO2, r.NO2)
			tempData.PM10 = append(tempData.PM10, r.PM10)
			tempData.PM25 = append(tempData.PM25, r.PM25)
		}
		for i := 1; i < 8; i++ {
			d.O3 = append(d.O3, math.Round(((tempData.O3[i*20-1]+tempData.O3[(i*20-1)-1]+tempData.O3[(i*20-1)-2]+tempData.O3[(i*20-1)-3]+tempData.O3[(i*20-1)-4]+tempData.O3[(i*20-1)-5]+tempData.O3[(i*20-1)-6]+tempData.O3[(i*20-1)-7]+tempData.O3[(i*20-1)-8]+tempData.O3[(i*20-1)-9]+tempData.O3[(i*20-1)-10]+tempData.O3[(i*20-1)-11]+tempData.O3[(i*20-1)-12]+tempData.O3[(i*20-1)-13]+tempData.O3[(i*20-1)-14]+tempData.O3[(i*20-1)-15]+tempData.O3[(i*20-1)-16]+tempData.O3[(i*20-1)-17]+tempData.O3[(i*20-1)-18]+tempData.O3[(i*20-1)-19])/20)*100)/100)
			d.NO2 = append(d.NO2, math.Round(((tempData.NO2[i*20-1]+tempData.NO2[(i*20-1)-1]+tempData.NO2[(i*20-1)-2]+tempData.NO2[(i*20-1)-3]+tempData.NO2[(i*20-1)-4]+tempData.NO2[(i*20-1)-5]+tempData.NO2[(i*20-1)-6]+tempData.NO2[(i*20-1)-7]+tempData.NO2[(i*20-1)-8]+tempData.NO2[(i*20-1)-9]+tempData.NO2[(i*20-1)-10]+tempData.NO2[(i*20-1)-11]+tempData.NO2[(i*20-1)-12]+tempData.NO2[(i*20-1)-13]+tempData.NO2[(i*20-1)-14]+tempData.NO2[(i*20-1)-15]+tempData.NO2[(i*20-1)-16]+tempData.NO2[(i*20-1)-17]+tempData.NO2[(i*20-1)-18]+tempData.NO2[(i*20-1)-19])/20)*100)/100)
			d.PM10 = append(d.PM10, math.Round(((tempData.PM10[i*20-1]+tempData.PM10[(i*20-1)-1]+tempData.PM10[(i*20-1)-2]+tempData.PM10[(i*20-1)-3]+tempData.PM10[(i*20-1)-4]+tempData.PM10[(i*20-1)-5]+tempData.PM10[(i*20-1)-6]+tempData.PM10[(i*20-1)-7]+tempData.PM10[(i*20-1)-8]+tempData.PM10[(i*20-1)-9]+tempData.PM10[(i*20-1)-10]+tempData.PM10[(i*20-1)-11]+tempData.PM10[(i*20-1)-12]+tempData.PM10[(i*20-1)-13]+tempData.PM10[(i*20-1)-14]+tempData.PM10[(i*20-1)-15]+tempData.PM10[(i*20-1)-16]+tempData.PM10[(i*20-1)-17]+tempData.PM10[(i*20-1)-18]+tempData.PM10[(i*20-1)-19])/20)*100)/100)
			d.PM25 = append(d.PM25, math.Round(((tempData.PM25[i*20-1]+tempData.PM25[(i*20-1)-1]+tempData.PM25[(i*20-1)-2]+tempData.PM25[(i*20-1)-3]+tempData.PM25[(i*20-1)-4]+tempData.PM25[(i*20-1)-5]+tempData.PM25[(i*20-1)-6]+tempData.PM25[(i*20-1)-7]+tempData.PM25[(i*20-1)-8]+tempData.PM25[(i*20-1)-9]+tempData.PM25[(i*20-1)-10]+tempData.PM25[(i*20-1)-11]+tempData.PM25[(i*20-1)-12]+tempData.PM25[(i*20-1)-13]+tempData.PM25[(i*20-1)-14]+tempData.PM25[(i*20-1)-15]+tempData.PM25[(i*20-1)-16]+tempData.PM25[(i*20-1)-17]+tempData.PM25[(i*20-1)-18]+tempData.PM25[(i*20-1)-19])/20)*100)/100)
		}
		defer rows.Close()
		res.Meta.Query = "Moyennes en temps réel de la concentration en μg/m3 d'ozone, de dioxyde d'azote et de particules en suspension (PM10/25) dans tous les arrondissements pour les 7 derniers jours"
	}
	d.ReverseValues()
	res.Data = append(res.Data, d)
	json.NewEncoder(w).Encode(res)
}
