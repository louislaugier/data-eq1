package maratp

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
)

type client struct {
	ID                 int    `json:"identifiant,omitempty"`
	Age                string `json:"age,omitempty"`
	Genre              string `json:"genre,omitempty"`
	CodePostal         string `json:"code_postal,omitempty"`
	Annciennete        int    `json:"anciennete,omitempty"`
	AbonneAlerting     bool   `json:"abonne_alerting,omitempty"`
	Alertes            string `json:"alertes,omitempty"`
	TitreTransport     string `json:"titre_transport,omitempty"`
	FrequenceTransport string `json:"frequence_transport,omitempty"`
	FavorisHoraires    string `json:"favoris_horaires,omitempty"`
	FavorisAdresses    string `json:"favoris_adresses,omitempty"`
}

type query struct {
	Limit              string `json:"limit"`
	Offset             string `json:"offset"`
	Genre              string `json:"genre,omitempty"`
	TitreTransport     string `json:"titre_transport,omitempty"`
	FrequenceTransport string `json:"frequence_transport,omitempty"`
}

func (q *query) standardizeSQLQuery(values url.Values) string {
	q.Limit = "1000"
	q.Offset = "0"
	if values.Get("limit") != "" {
		q.Limit = values.Get("limit")
	}
	if values.Get("offset") != "" {
		q.Offset = values.Get("offset")
	}
	var genre string
	if values.Get("genre") != "" {
		genre = "genre = '" + values.Get("genre") + "' AND "
		q.Genre = values.Get("genre")
	}
	var titreTransport string
	if values.Get("titre_transport") != "" {
		titreTransport = "titre_transport = '" + values.Get("titre_transport") + "' AND "
		q.TitreTransport = values.Get("titre_transport")
	}
	var frequenceTransport string
	if values.Get("code_postal") != "" {
		frequenceTransport = "frequence_transport = '" + values.Get("frequence_transport") + "' AND "
		q.FrequenceTransport = values.Get("frequence_transport")
	}
	if q.Genre != "" || q.TitreTransport != "" || q.FrequenceTransport != "" {
		return strings.TrimSuffix("WHERE "+genre+titreTransport+frequenceTransport, " AND ")
	}
	return ""
}

// ClientsGET gets all clients
func ClientsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var q query
	u := r.URL.Query()
	s := q.standardizeSQLQuery(u)
	rows, err := database.Db.Query("SELECT identifiant, age, genre, code_postal, anciennete, abonne_alerting, alertes, titre_transport, frequence_transport, favoris_horaires, favoris_adresses FROM clients " + s + " LIMIT " + q.Limit + " OFFSET " + q.Offset + ";")
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
		var client client
		rows.Scan(&client.ID, &client.Age, &client.Genre, &client.CodePostal, &client.Annciennete, &client.AbonneAlerting, &client.Alertes, &client.TitreTransport, &client.FrequenceTransport, &client.FavorisHoraires, &client.FavorisAdresses)
		res.Data = append(res.Data, client)
	}
	defer rows.Close()
	res.Meta.Query = q
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// ClientGET gets a client by id
func ClientGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	rows, err := database.Db.Query("SELECT identifiant, age, genre, code_postal, anciennete, abonne_alerting, alertes, titre_transport, frequence_transport, favoris_horaires, favoris_adresses FROM clients WHERE identifiant = '" + ID + "' LIMIT 1;")
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
	var client client
	for rows.Next() {
		rows.Scan(&client.ID, &client.Age, &client.Genre, &client.CodePostal, &client.Annciennete, &client.AbonneAlerting, &client.Alertes, &client.TitreTransport, &client.FrequenceTransport, &client.FavorisHoraires, &client.FavorisAdresses)
	}
	defer rows.Close()
	if client.ID != 0 {
		res.Data = append(res.Data, client)
	}
	res.Meta.Query = ID
	res.Meta.ResultCount = len(res.Data)
	if res.Meta.ResultCount == 0 {
		res.Error = "Non-existing ID: " + ID
		res.Message = "User not found"
	}
	json.NewEncoder(w).Encode(res)
}

// ClientCreate creates a client
func ClientCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var client client
	_ = json.NewDecoder(r.Body).Decode(&client)
	txn, err := database.Db.Begin()
	if err != nil {
		log.Println(err)
	}
	txn.Exec("INSERT INTO clients (identifiant, age, genre, code_postal, anciennete, abonne_alerting, alertes, titre_transport, frequence_transport, favoris_horaires, favoris_adresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);", client.ID, client.Age, client.Genre, client.CodePostal, client.Annciennete, client.AbonneAlerting, client.Alertes, client.TitreTransport, client.FrequenceTransport, client.FavorisHoraires, client.FavorisAdresses)
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "Client " + strconv.Itoa(client.ID) + " created",
	}
	err = txn.Commit()
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	res.Data = append(res.Data, client)
	res.Meta.Query = "POST " + strconv.Itoa(client.ID)
	json.NewEncoder(w).Encode(res)
}

// ClientUpdate updates a client
func ClientUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	var client client
	_ = json.NewDecoder(r.Body).Decode(&client)
	txn, err := database.Db.Begin()
	if err != nil {
		log.Println(err)
	}
	txn.Exec("UPDATE clients SET identifiant = $1, age = $2, genre = $3, code_postal = $4, anciennete = $5, abonne_alerting = $6, alertes = $7, titre_transport = $8, frequence_transport = $9, favoris_horaires = $10, favoris_adresses = $11 WHERE identifiant = "+ID+";", client.ID, client.Age, client.Genre, client.CodePostal, client.Annciennete, client.AbonneAlerting, client.Alertes, client.TitreTransport, client.FrequenceTransport, client.FavorisHoraires, client.FavorisAdresses)
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "Client " + ID + " updated",
	}
	err = txn.Commit()
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	res.Data = append(res.Data, client)
	res.Meta.Query = "PUT " + ID
	json.NewEncoder(w).Encode(res)
}

// ClientDelete deletes a client
func ClientDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	txn, err := database.Db.Begin()
	if err != nil {
		log.Println(err)
	}
	txn.Exec("DELETE FROM clients WHERE identifiant = " + ID + ";")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "Client " + ID + " deleted",
	}
	err = txn.Commit()
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	res.Meta.Query = "DELETE " + ID
	json.NewEncoder(w).Encode(res)
}
