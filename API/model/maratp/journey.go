package maratp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
)

type journey struct {
	ID              string `json:"identifiant_trajet,omitempty"`
	UserID          int    `json:"identifiant_utilisateur,omitempty"`
	PointDepart     string `json:"point_depart,omitempty"`
	PointArrivee    string `json:"point_arrivee,omitempty"`
	TypeTrajet      string `json:"type_trajet,omitempty"`
	ScoreAir        int    `json:"score_air"`
	ScoreAffluence  int    `json:"score_affluence"`
	ScoreSilence    int    `json:"score_silence"`
	ScoreVision     int    `json:"score_vision"`
	ScoreMarche     int    `json:"score_marche"`
	ScoreNature     int    `json:"score_nature"`
	NoteUtilisateur int    `json:"note_utilisateur,omitempty"`
	ImpactPollution int    `json:"impact_pollution,omitempty"`
	ImpactAffluence int    `json:"impact_affluence,omitempty"`
	ImpactSon       int    `json:"impact_son,omitempty"`
}

type people struct {
	Women int `json:"women"`
	Men   int `json:"men"`
}

type data struct {
	Totals struct {
		Total      int `json:"total,omitempty"`
		TotalWomen int `json:"women,omitempty"`
		TotalMen   int `json:"men,omitempty"`
	} `json:"totals"`
	ZeroToEighteen        people `json:"zero_to_eighteen"`
	NineteenToTwentyFive  people `json:"nineteen_to_twentyfive"`
	TwentySixToThirtyFour people `json:"twentysix_to_thirtyfour"`
	ThirtyFiveToFifty     people `json:"thirtyfive_to_fifty"`
	FiftyOneToSixtyFive   people `json:"fiftyone_to_sixtyfive"`
	SixtySixPlus          people `json:"sixtysix_and_over"`
}

// JourneysGET gets all journeys
func JourneysGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rows, err := database.Db.Query("SELECT id_trajet, id_user, point_depart, point_arrivee, type_trajet, score_air, score_affluence, score_silence, score_vision, score_marche, score_nature, note_utilisateur FROM trajets;")
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
		var journey journey
		rows.Scan(&journey.ID, &journey.UserID, &journey.PointDepart, &journey.PointArrivee, &journey.TypeTrajet, &journey.ScoreAir, &journey.ScoreAffluence, &journey.ScoreSilence, &journey.ScoreVision, &journey.ScoreMarche, &journey.ScoreNature, &journey.NoteUtilisateur)
		res.Data = append(res.Data, journey)
	}
	defer rows.Close()
	res.Meta.Query = "Trajets des utilisateurs MaRATP"
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// JourneyTypeGET gets the number of users and their age and gender for a journey type
func JourneyTypeGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	t, _ := mux.Vars(r)["type"]
	rows, err := database.Db.Query("SELECT id_user FROM trajets WHERE type_trajet = '" + t + "';")
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
	var userIDs []int
	for rows.Next() {
		var ID int
		rows.Scan(&ID)
		userIDs = append(userIDs, ID)
	}
	var clients []client
	for _, value := range userIDs {
		rows, err = database.Db.Query("SELECT age, genre FROM clients WHERE identifiant = '" + strconv.Itoa(value) + "';")
		var res model.Response
		if err != nil {
			res.StatusCode = 500
			res.Message = "Internal Server Error"
			res.Error = "Database Error"
			log.Println("Error with transaction:", err)
		}
		for rows.Next() {
			var client client
			rows.Scan(&client.Age, &client.Genre)
			clients = append(clients, client)
		}
		rows.Close()
	}
	var d data
	d.Totals.Total = len(clients)
	for _, c := range clients {
		if c.Genre == "femme" {
			d.Totals.TotalWomen++
		} else if c.Genre == "homme" {
			d.Totals.TotalMen++
		}
		age, _ := strconv.Atoi(c.Age)
		switch {
		case age > 0 && age < 19:
			if c.Genre == "femme" {
				d.ZeroToEighteen.Women++
			} else if c.Genre == "homme" {
				d.ZeroToEighteen.Men++
			}
		case age > 18 && age < 26:
			if c.Genre == "femme" {
				d.NineteenToTwentyFive.Women++
			} else if c.Genre == "homme" {
				d.NineteenToTwentyFive.Men++
			}
		case age > 25 && age < 35:
			if c.Genre == "femme" {
				d.TwentySixToThirtyFour.Women++
			} else if c.Genre == "homme" {
				d.TwentySixToThirtyFour.Men++
			}
		case age > 36 && age < 51:
			if c.Genre == "femme" {
				d.ThirtyFiveToFifty.Women++
			} else if c.Genre == "homme" {
				d.ThirtyFiveToFifty.Men++
			}
		case age > 50 && age < 66:
			if c.Genre == "femme" {
				d.FiftyOneToSixtyFive.Women++
			} else if c.Genre == "homme" {
				d.FiftyOneToSixtyFive.Men++
			}
		default:
			if c.Genre == "femme" {
				d.SixtySixPlus.Women++
			} else if c.Genre == "homme" {
				d.SixtySixPlus.Men++
			}
		}
	}
	// add random fake data for a realistic front-end
	d.ZeroToEighteen.Women += 250*d.ZeroToEighteen.Women + 2541
	d.ZeroToEighteen.Men += 250*d.ZeroToEighteen.Men + 2622
	d.NineteenToTwentyFive.Women += 250*d.NineteenToTwentyFive.Women + 4158
	d.NineteenToTwentyFive.Men += 250*d.NineteenToTwentyFive.Men + 3482
	d.TwentySixToThirtyFour.Women += 250*d.TwentySixToThirtyFour.Women + 2555
	d.TwentySixToThirtyFour.Men += 250*d.TwentySixToThirtyFour.Men + 1651
	d.ThirtyFiveToFifty.Women += 250*d.ThirtyFiveToFifty.Women + 1710
	d.ThirtyFiveToFifty.Men += 250*d.ThirtyFiveToFifty.Men + 1875
	d.FiftyOneToSixtyFive.Women += 250*d.FiftyOneToSixtyFive.Women + 1520
	d.FiftyOneToSixtyFive.Men += 250*d.FiftyOneToSixtyFive.Men + 2590
	d.SixtySixPlus.Women += 250*d.SixtySixPlus.Women + 3469
	d.SixtySixPlus.Men += 250*d.SixtySixPlus.Men + 2686
	d.Totals.TotalWomen = (d.ZeroToEighteen.Women + d.NineteenToTwentyFive.Women + d.TwentySixToThirtyFour.Women + d.ThirtyFiveToFifty.Women + d.FiftyOneToSixtyFive.Women + d.SixtySixPlus.Women)
	d.Totals.TotalMen = (d.ZeroToEighteen.Men + d.NineteenToTwentyFive.Men + d.TwentySixToThirtyFour.Men + d.ThirtyFiveToFifty.Men + d.FiftyOneToSixtyFive.Men + d.SixtySixPlus.Men)
	d.Totals.Total = (d.Totals.TotalWomen + d.Totals.TotalMen)
	res.Data = append(res.Data, d)
	res.Meta.Query = "Utilisateurs des trajets avec l'option : " + t
	json.NewEncoder(w).Encode(res)
}
