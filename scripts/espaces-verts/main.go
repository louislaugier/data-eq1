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

type espaceVert struct {
	Identifiant                 int
	Nom                         string
	Typologie                   string
	Categorie                   string
	AdresseNumero               int
	AdresseComplement           string
	AdresseTypeVoie             string
	AdresseLibelleVoie          string
	CodePostal                  int
	SurfaceCalculee             int
	SuperficieTotaleReelle      int
	SurfaceHorticole            int
	PresenceCloture             string
	Perimetre                   float64
	AnneeOuverture              int
	AnneeRenovation             int
	AncienNom                   string
	AnneeChangementNom          int
	NombreEntites               int
	Ouverture24                 string
	IdentifiantDivision         int
	IdentifiantAtelierHorticole int
	Ida3dEnb                    string
	SiteVilles                  string
	IdentifiantEquipement       string
	Competence                  string
	GeoShape                    string
	URLPlan                     string
	GeoPoint                    string
}

func main() {
	db, err := sql.Open(database.Type, database.URI)
	if err != nil {
		fmt.Println("Error connecting to database", db)
	}
	f, err := os.Open("espaces_verts.csv")
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
	espacesVerts := []espaceVert{}
	for indexRecords, record := range records {
		if indexRecords > 0 {
			e := espaceVert{
				Nom:                   record[1],
				Typologie:             record[2],
				Categorie:             record[3],
				AdresseComplement:     record[5],
				AdresseTypeVoie:       record[6],
				AdresseLibelleVoie:    record[7],
				PresenceCloture:       record[12],
				AncienNom:             record[16],
				Ouverture24:           record[19],
				Ida3dEnb:              record[22],
				SiteVilles:            record[23],
				IdentifiantEquipement: record[24],
				Competence:            record[25],
				GeoShape:              record[26],
				URLPlan:               record[27],
				GeoPoint:              record[28],
			}
			for indexFields, field := range record {
				switch indexFields {
				case 0:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.Identifiant = i
				case 4:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.AdresseNumero = i
				case 8:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.CodePostal = i
				case 9:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.SurfaceCalculee = i
				case 10:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.SuperficieTotaleReelle = i
				case 11:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.SurfaceHorticole = i
				case 13:
					i, err := strconv.ParseFloat(field, 64)
					if err != nil {
						fmt.Println("Error converting to float64")
					}
					e.Perimetre = i
				case 14:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.AnneeOuverture = i
				case 15:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.AnneeRenovation = i
				case 17:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.AnneeChangementNom = i
				case 18:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.NombreEntites = i
				case 20:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.IdentifiantDivision = i
				case 21:
					i, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Error converting to int")
					}
					e.IdentifiantAtelierHorticole = i
				}
				espacesVerts = append(espacesVerts, e)
				_, err := txn.Exec("INSERT INTO espaces_verts (identifiant_espace_vert, nom, typologie, categorie, adresse_numero, adresse_complement, adresse_type_voie, adresse_libelle_voie, code_postal, surface_calculee, superficie_totale_reelle, surface_horticole, presence_cloture, perimetre, annee_ouverture, annee_renovation, ancien_nom, annee_changement_nom, nombre_entites, ouverture_24h_24h, id_division, id_atelier_horticole, ida3d_enb, site_villes, id_eqpt, competence, geo_shape, url_plan, geo_point) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29);", e.Identifiant, e.Nom, e.Typologie, e.Categorie, e.AdresseNumero, e.AdresseComplement, e.AdresseTypeVoie, e.AdresseLibelleVoie, e.CodePostal, e.SurfaceCalculee, e.SuperficieTotaleReelle, e.SurfaceHorticole, e.PresenceCloture, e.Perimetre, e.AnneeOuverture, e.AnneeRenovation, e.AncienNom, e.AnneeChangementNom, e.NombreEntites, e.Ouverture24, e.IdentifiantDivision, e.IdentifiantAtelierHorticole, e.Ida3dEnb, e.SiteVilles, e.IdentifiantEquipement, e.Competence, e.GeoShape, e.URLPlan, e.GeoPoint)
				if err != nil {
					fmt.Println("Error inserting:", err)
				}
			}
		}
	}
	fmt.Println("Nombre de lignes dans le CSV : ", len(records))
	fmt.Println("Nombre d'espaces verts : ", len(espacesVerts))
	err = txn.Commit()
	if err != nil {
		fmt.Println("Error with transaction:", err)
	}
	db.Close()
}
