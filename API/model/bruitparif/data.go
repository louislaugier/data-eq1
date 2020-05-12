package bruitparif

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
)

type data struct {
	Decibels []float64
}

func (d data) reverseValues() {
	for i, j := 0, len(d.Decibels)-1; i < j; i, j = i+1, j-1 {
		d.Decibels[i], d.Decibels[j] = d.Decibels[j], d.Decibels[i]
	}
}

// DataGET gets bruitparif data from last 7 recorded days for 1 arrondissement or averages for all of Paris
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
		rows, err := database.Db.Query("SELECT decibels FROM nuisances_sonores WHERE arrondissement = " + arr + " ORDER BY date DESC LIMIT 7;")
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Database Error"
			log.Println("Error with transaction:", err)
		}
		for rows.Next() {
			var decibels float64
			rows.Scan(&decibels)
			d.Decibels = append(d.Decibels, decibels)
		}
		if arr == "1" {
			res.Meta.Query = "Moyennes historiques des nuisances sonores dans le " + arr + "er arrondissement en décibels"
		} else {
			res.Meta.Query = "Moyennes historiques des nuisances sonores dans le " + arr + "ème arrondissement en décibels"
		}
		defer rows.Close()
	} else {
		rows, err := database.Db.Query("SELECT decibels FROM nuisances_sonores ORDER BY date DESC, arrondissement;")
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Database Error"
			log.Println("Error with transaction:", err)
		}
		var tempData data
		for rows.Next() {
			var decibels float64
			rows.Scan(&decibels)
			tempData.Decibels = append(tempData.Decibels, decibels)
		}
		for i := 1; i < 8; i++ {
			d.Decibels = append(d.Decibels, math.Round(((tempData.Decibels[i*20-1]+tempData.Decibels[(i*20-1)-1]+tempData.Decibels[(i*20-1)-2]+tempData.Decibels[(i*20-1)-3]+tempData.Decibels[(i*20-1)-4]+tempData.Decibels[(i*20-1)-5]+tempData.Decibels[(i*20-1)-6]+tempData.Decibels[(i*20-1)-7]+tempData.Decibels[(i*20-1)-8]+tempData.Decibels[(i*20-1)-9]+tempData.Decibels[(i*20-1)-10]+tempData.Decibels[(i*20-1)-11]+tempData.Decibels[(i*20-1)-12]+tempData.Decibels[(i*20-1)-13]+tempData.Decibels[(i*20-1)-14]+tempData.Decibels[(i*20-1)-15]+tempData.Decibels[(i*20-1)-16]+tempData.Decibels[(i*20-1)-17]+tempData.Decibels[(i*20-1)-18]+tempData.Decibels[(i*20-1)-19])/20)*100)/100)
		}
		defer rows.Close()
		res.Meta.Query = "Moyennes historiques des nuisances sonores dans tous les arrondissements en décibels"
	}
	d.reverseValues()
	res.Data = append(res.Data, d)
	json.NewEncoder(w).Encode(res)
}
