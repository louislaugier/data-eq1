package maratp

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
)

type score struct {
	ScoreName  string  `json:"name"`
	ScoreValue float64 `json:"score"`
}

// ClientsScoringGET gets criteria ranking and scoring averages for all user journeys
func ClientsScoringGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rows, err := database.Db.Query("SELECT score_air, score_affluence, score_silence, score_vision, score_marche, score_nature FROM trajets;")
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
	defer rows.Close()
	var journeys []journey
	for rows.Next() {
		var journey journey
		rows.Scan(&journey.ScoreAir, &journey.ScoreAffluence, &journey.ScoreSilence, &journey.ScoreVision, &journey.ScoreMarche, &journey.ScoreNature)
		journeys = append(journeys, journey)
	}
	scores := map[string]float64{}
	for _, v := range journeys {
		scores["Air"] += float64(v.ScoreAir)
		scores["Affluence"] += float64(v.ScoreAffluence)
		scores["Silence"] += float64(v.ScoreSilence)
		scores["Vision"] += float64(v.ScoreVision)
		scores["Marche"] += float64(v.ScoreMarche)
		scores["Nature"] += float64(v.ScoreNature)
	}
	for k, v := range scores {
		scores[k] = v / float64(len(journeys))
	}
	sortedScores := map[float64]string{}
	keys := []float64{}
	for k, v := range scores {
		sortedScores[v] = k
		keys = append(keys, v)
	}
	sort.Float64s(keys)
	var scoring []*score
	for _, v := range keys {
		scoring = append(scoring, &score{
			ScoreName:  sortedScores[v],
			ScoreValue: v,
		})
	}
	for i, j := 0, len(scoring)-1; i < j; i, j = i+1, j-1 {
		scoring[i], scoring[j] = scoring[j], scoring[i]
	}
	res.Data = append(res.Data, scoring)
	res.Meta.Query = "Classement des critères et scoring des trajets des utilisateurs"
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// ClientScoringGET gets criteria ranking and scoring from user journeys by user ID
func ClientScoringGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	rows, err := database.Db.Query("SELECT score_air, score_affluence, score_silence, score_vision, score_marche, score_nature FROM trajets WHERE id_user = '" + ID + "';")
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
	defer rows.Close()
	var journey journey
	for rows.Next() {
		rows.Scan(&journey.ScoreAir, &journey.ScoreAffluence, &journey.ScoreSilence, &journey.ScoreVision, &journey.ScoreMarche, &journey.ScoreNature)
	}
	var scoring []score
	for i := 0; i < reflect.ValueOf(&journey).Elem().NumField(); i++ {
		if strings.Contains(reflect.ValueOf(&journey).Elem().Type().Field(i).Name, "Score") {
			var record score
			record.ScoreName = strings.TrimPrefix(reflect.ValueOf(&journey).Elem().Type().Field(i).Name, "Score")
			record.ScoreValue = float64(reflect.ValueOf(&journey).Elem().Field(i).Interface().(int))
			scoring = append(scoring, record)
		}
	}
	sort.Slice(scoring, func(p, q int) bool {
		return scoring[p].ScoreValue > scoring[q].ScoreValue
	})
	res.Data = append(res.Data, scoring)
	res.Meta.Query = "Classement des critères et scoring des trajets de l'utilisateur " + ID
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}
