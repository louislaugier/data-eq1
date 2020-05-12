package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	database "github.com/louislaugier/projet-data-4A-scripts"
)

type trajet struct {
	IdentifiantTrajet      string
	IdentifiantUtilisateur int
	PointDepart            string
	PointArrivee           string
	TypeTrajet             string
	ScoreAir               int
	ScoreSilence           int
	ScoreVision            int
	ScoreMarche            int
	ScoreNature            int
	NoteUtilisateur        int
}

func main() {
	db, err := sql.Open(database.Type, database.URI)
	if err != nil {
		fmt.Println("Error connecting to database", db)
	}
	f, err := os.Open("trajets.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}
	txn, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	trajets := []trajet{}
	for indexRecords, record := range records {
		if indexRecords > 0 {
			t := trajet{
				IdentifiantTrajet: record[0],
				PointDepart:       record[2],
				PointArrivee:      record[3],
				TypeTrajet:        record[4],
			}
			for indexFields, field := range record {
				fmt.Println(indexFields, field)
				switch indexFields {
				case 1:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.IdentifiantUtilisateur = i
				case 5:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.ScoreAir = i
				case 6:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.ScoreSilence = i
				case 7:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.ScoreVision = i
				case 8:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.ScoreMarche = i
				case 9:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.ScoreNature = i
				case 10:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					t.NoteUtilisateur = i
				}
				trajets = append(trajets, t)
				_, err := txn.Exec("INSERT INTO trajets (id_trajet, id_user, point_depart, point_arrivee, type_trajet, score_air, score_affluence, score_silence, score_vision, score_marche, score_nature, note_utilisateur) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);", t.IdentifiantTrajet, t.IdentifiantUtilisateur, t.PointDepart, t.PointArrivee, t.TypeTrajet, t.ScoreAir, t.ScoreSilence, t.ScoreVision, t.ScoreMarche, t.ScoreNature, t.NoteUtilisateur)
				if err != nil {
					fmt.Println("Error inserting:", err)
				}
			}
		}
	}
	fmt.Println("Nombre de lignes dans le CSV : ", len(records))
	fmt.Println("Nombre de stations : ", len(trajets))
	err = txn.Commit()
	if err != nil {
		fmt.Println("Error with transaction:", err)
	}
	db.Close()
}
