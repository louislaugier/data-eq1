package velib

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"data-eq1/API/database"
	"data-eq1/API/model"

	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
)

type station struct {
	IdentifiantStation          int             `json:"identifiant_station,omitempty"`
	Feature                     geojson.Feature `json:"feature,omitempty"`
	NomStation                  string          `json:"nom_station,omitempty"`
	StationFonctionnment        string          `json:"station_en_fonctionnement,omitempty"`
	CapaciteStation             int             `json:"capacite_station,omitempty"`
	NombreBornettesLibres       int             `json:"nombre_bornettes_libres,omitempty"`
	NombreTotalVelosDisponibles int             `json:"nombre_total_velos_disponibles,omitempty"`
	VelosMecaniquesDisponibles  int             `json:"velos_mecaniques_disponibles,omitempty"`
	VelosElectriquesDisponibles int             `json:"velos_electriques_disponibles,omitempty"`
	BornePaiementDispo          string          `json:"borne_paiement_dispo,omitempty"`
	RetourVelibPossible         string          `json:"retour_velib_possible,omitempty"`
	DerniereActualisation       time.Time       `json:"derniere_actualisation,omitempty"`
	Coordonnees                 string          `json:"coordonnees,omitempty"`
	NomCommunesEquipees         string          `json:"nom_communes_equipees,omitempty"`
	CodeInseeCommunesEquipees   int             `json:"code_insee_communes_equipees,omitempty"`
}

type query struct {
	Limit               string `json:"limit"`
	Offset              string `json:"offset"`
	BornePaiementDispo  string `json:"borne,omitempty"`
	RetourVelibPossible string `json:"retour,omitempty"`
}

func (s *station) stringToGeoJSONFeature() {
	s.Feature.Geometry = &geojson.Geometry{}
	s.Feature.Properties = map[string]interface{}{
		"identifiantStation":          s.IdentifiantStation,
		"nomStation":                  s.NomStation,
		"stationFonctionnement":       s.StationFonctionnment,
		"capaciteStation":             s.CapaciteStation,
		"nombreBornettesLibres":       s.NombreBornettesLibres,
		"nombreTotalVelosDisponibles": s.NombreTotalVelosDisponibles,
		"velosMecaniquesDisponibles":  s.VelosMecaniquesDisponibles,
		"velosElectriquesDisponibles": s.VelosElectriquesDisponibles,
		"bornePaiementDispo":          s.BornePaiementDispo,
		"retourVelibPossible":         s.RetourVelibPossible,
		"derniereActualisation":       s.DerniereActualisation,
		"nomCommunesEquipees":         s.NomCommunesEquipees,
		"codeInseeCommunesEquipees":   s.CodeInseeCommunesEquipees,
	}
	s.Feature.Type = "Point"
	s.Feature.Geometry.Type = "Point"
	lng, _ := strconv.ParseFloat(strings.Split(s.Coordonnees, ",")[0], 64)
	lat, _ := strconv.ParseFloat(strings.Split(s.Coordonnees, ",")[1], 64)
	s.Feature.Geometry.Point = []float64{lat, lng}
}

func (q *query) standardizeSQLQuery(values url.Values) string {
	q.Limit = "1398"
	q.Offset = "0"
	if values.Get("limit") != "" {
		q.Limit = values.Get("limit")
	}
	if values.Get("offset") != "" {
		q.Offset = values.Get("offset")
	}
	var borne string
	if values.Get("borne") != "" {
		borne = "borne_paiement_dispo = '" + values.Get("borne") + "' AND "
		q.BornePaiementDispo = values.Get("borne")
	}
	var retour string
	if values.Get("retour") != "" {
		retour = "retour_velib_possible = '" + values.Get("retour") + "' AND "
		q.RetourVelibPossible = values.Get("retour")
	}

	if q.BornePaiementDispo != "" || q.RetourVelibPossible != "" {
		return strings.TrimSuffix("WHERE "+borne+retour, " AND ")
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
	rows, err := database.Db.Query("SELECT identifiant_station, nom_station, station_en_fonctionnement, capacite_station, nb_bornettes_libres, nb_total_velos_dispo, velos_mecaniques_dispo, velos_electriques_dispo, borne_paiement_dispo, retour_velib_possible, derniere_actualisation, coordonnees, nom_communes_equipees, code_insee_communes_equipees FROM stations_velib " + s + " LIMIT " + q.Limit + " OFFSET " + q.Offset + ";")
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
	for rows.Next() {
		var station station
		rows.Scan(&station.IdentifiantStation, &station.NomStation, &station.StationFonctionnment, &station.CapaciteStation, &station.NombreBornettesLibres, &station.NombreTotalVelosDisponibles, &station.VelosMecaniquesDisponibles, &station.VelosElectriquesDisponibles, &station.BornePaiementDispo, &station.RetourVelibPossible, &station.DerniereActualisation, &station.Coordonnees, &station.NomCommunesEquipees, &station.CodeInseeCommunesEquipees)
		station.stringToGeoJSONFeature()
		res.Data = append(res.Data, station)
	}
	defer rows.Close()
	res.Meta.Query = q
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// StationGET gets a station by id
func StationGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	rows, err := database.Db.Query("SELECT identifiant_station, nom_station, station_en_fonctionnement, capacite_station, nb_bornettes_libres, nb_total_velos_dispo, velos_mecaniques_dispo, velos_electriques_dispo, borne_paiement_dispo, retour_velib_possible, derniere_actualisation, coordonnees, nom_communes_equipees, code_insee_communes_equipees FROM stations_velib WHERE identifiant_station = '" + ID + "' LIMIT 1;")
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
		rows.Scan(&station.IdentifiantStation, &station.NomStation, &station.StationFonctionnment, &station.CapaciteStation, &station.NombreBornettesLibres, &station.NombreTotalVelosDisponibles, &station.VelosMecaniquesDisponibles, &station.VelosElectriquesDisponibles, &station.BornePaiementDispo, &station.RetourVelibPossible, &station.DerniereActualisation, &station.Coordonnees, &station.NomCommunesEquipees, &station.CodeInseeCommunesEquipees)
	}
	defer rows.Close()
	if station.IdentifiantStation != 0 {
		res.Data = append(res.Data, station)
	}
	res.Meta.Query = ID
	res.Meta.ResultCount = len(res.Data)
	if res.Meta.ResultCount == 0 {
		res.Error = "Non-existing ID: " + ID
		res.Message = "Velib station not found"
	}
	json.NewEncoder(w).Encode(res)
}
