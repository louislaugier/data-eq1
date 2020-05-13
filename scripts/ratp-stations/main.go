package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	database "data-eq1/scripts"

	_ "github.com/lib/pq"
)

type station struct {
	ID          int
	Name        string
	Address     string
	Coordinates string
}

func main() {
	db, err := sql.Open(database.Type, database.URI)
	if err != nil {
		fmt.Println("Error connecting to database", db)
	}
	f, err := os.Open("ratp-stations.csv")
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
	stations := []station{}
	for indexRecords, record := range records {
		if indexRecords > 0 {
			s := station{
				Name:        record[1],
				Address:     record[2],
				Coordinates: record[3],
			}
			for indexFields, field := range record {
				fmt.Println(indexFields, field)
				switch indexFields {
				case 0:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					s.ID = i
				}
				stations = append(stations, s)
				_, err := txn.Exec("INSERT INTO ratp_stations (id, name, address, coordinates) VALUES ($1, $2, $3, $4);", s.ID, s.Name, s.Address, s.Coordinates)
				if err != nil {
					fmt.Println("Error inserting:", err)
				}
			}
		}
	}
	fmt.Println("Nombre de lignes dans le CSV : ", len(records))
	fmt.Println("Nombre de stations : ", len(stations))
	err = txn.Commit()
	if err != nil {
		fmt.Println("Error with transaction:", err)
	}
	db.Close()
}
