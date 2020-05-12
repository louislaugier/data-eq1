package router

import (
	"net/http"

	"data-eq1/API/model/airparif"
	"data-eq1/API/model/bruitparif"
	"data-eq1/API/model/database"
	"data-eq1/API/model/elevation"
	"data-eq1/API/model/maratp"
	"data-eq1/API/model/paris"
	"data-eq1/API/model/ratp"
	"data-eq1/API/model/velib"

	"github.com/gorilla/mux"
)

// Init initiates the router
func Init() *mux.Router {
	r := mux.NewRouter()
	a := "/api/v1"
	// Swagger docs
	r.PathPrefix(a + "/documentation").Handler(http.StripPrefix(a+"/documentation", http.FileServer(http.Dir("./docs"))))
	// Airparif routes
	r.HandleFunc(a+"/zones-qualite-air", airparif.ZonesGET).Methods("GET")
	r.HandleFunc(a+"/indices-qualite-air/{arrondissement}", airparif.DataGET).Methods("GET")
	r.HandleFunc(a+"/arrondissements-qualite-air", airparif.ArrondissementsGET).Methods("GET")
	// Bruitparif routes
	r.HandleFunc(a+"/zones-nuisances-sonores", bruitparif.ZonesGET).Methods("GET")
	r.HandleFunc(a+"/indices-nuisances-sonores/{arrondissement}", bruitparif.DataGET).Methods("GET")
	// Database routes
	r.HandleFunc(a+"/tables-bdd", database.TablesGET).Methods("GET")
	r.HandleFunc(a+"/colonnes-table/{table}", database.ColsGET).Methods("GET")
	r.HandleFunc(a+"/rangees-table/{table}", database.RowsGET).Methods("GET")
	// Elevation routes
	r.HandleFunc(a+"/zones-elevation", elevation.ZonesGET).Methods("GET")
	r.HandleFunc(a+"/indices-elevation", elevation.DataGET).Methods("GET")
	r.HandleFunc(a+"/indice-elevation/{arrondissement}", elevation.ArrondissementGET).Methods("GET")
	// MaRATP routes
	r.HandleFunc(a+"/maratp-clients", maratp.ClientsGET).Methods("GET")
	r.HandleFunc(a+"/maratp-client/{id}", maratp.ClientGET).Methods("GET")
	r.HandleFunc(a+"/maratp-clients", maratp.ClientCreate).Methods("POST")
	r.HandleFunc(a+"/maratp-client/{id}", maratp.ClientUpdate).Methods("PUT")
	r.HandleFunc(a+"/maratp-client/{id}", maratp.ClientDelete).Methods("DELETE")
	r.HandleFunc(a+"/maratp-trajets", maratp.JourneysGET).Methods("GET")
	r.HandleFunc(a+"/maratp-trajets/{type}", maratp.JourneyTypeGET).Methods("GET")
	r.HandleFunc(a+"/maratp-clients-scoring", maratp.ClientsScoringGET).Methods("GET")
	r.HandleFunc(a+"/maratp-client-scoring/{id}", maratp.ClientScoringGET).Methods("GET")
	// Paris routes
	r.HandleFunc(a+"/espaces-verts", paris.GreenAreasGET).Methods("GET")
	r.HandleFunc(a+"/espace-vert/{id}", paris.GreenAreaGET).Methods("GET")
	r.HandleFunc(a+"/meteo", paris.WeatherGET).Methods("GET")
	// RATP routes
	r.HandleFunc(a+"/stations-ratp", ratp.StationsGET).Methods("GET")
	r.HandleFunc(a+"/station-ratp/{id}", ratp.StationGET).Methods("GET")
	// Velib routes
	r.HandleFunc(a+"/stations-velib", velib.StationsGET).Methods("GET")
	r.HandleFunc(a+"/station-velib/{id}", velib.StationGET).Methods("GET")
	return r
}
