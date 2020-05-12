package paris

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
	geojson "github.com/paulmach/go.geojson"
)

type greenArea struct {
	Identifiant                 int             `json:"identifiant_espace_vert,omitempty"`
	GeoShape                    string          `json:"geo_shape,omitempty"`
	Feature                     geojson.Feature `json:"feature,omitempty"`
	Nom                         string          `json:"nom_de_lespace_vert,omitempty"`
	Typologie                   string          `json:"typologie_d_espace_vert,omitempty"`
	Categorie                   string          `json:"categorie,omitempty"`
	AdresseNumero               int             `json:"adresse_numero,omitempty"`
	AdresseComplement           string          `json:"adresse_complement,omitempty"`
	AdresseTypeVoie             string          `json:"adresse_type_voie,omitempty"`
	AdresseLibelleVoie          string          `json:"adresse_libelle_voie,omitempty"`
	CodePostal                  int             `json:"code_postal,omitempty"`
	SurfaceCalculee             int             `json:"surface_calculee,omitempty"`
	SuperficieTotaleReelle      int             `json:"superficie_totale_reelle,omitempty"`
	SurfaceHorticole            int             `json:"surface_horticole,omitempty"`
	PresenceCloture             string          `json:"presence_cloture,omitempty"`
	Perimetre                   float64         `json:"perimetre,omitempty"`
	AnneeOuverture              int             `json:"annee_ouverture,omitempty"`
	AnneeRenovation             int             `json:"annee_renovation,omitempty"`
	AncienNom                   string          `json:"ancien_nom,omitempty"`
	AnneeChangementNom          int             `json:"annee_changement_nom,omitempty"`
	NombreEntites               int             `json:"nombre_d_entites,omitempty"`
	Ouverture24                 string          `json:"ouverture_24h_sur24,omitempty"`
	IdentifiantDivision         int             `json:"identifiant_division,omitempty"`
	IdentifiantAtelierHorticole int             `json:"identifiant_atelier_horticole,omitempty"`
	Ida3dEnb                    string          `json:"ida3d_enb,omitempty"`
	SiteVilles                  string          `json:"site_villes,omitempty"`
	IdentifiantEquipement       string          `json:"identifiant_equipement,omitempty"`
	Competence                  string          `json:"competence,omitempty"`
	URLPlan                     string          `json:"url_plan,omitempty"`
	GeoPoint                    string          `json:"geo_point,omitempty"`
}

type query struct {
	Limit           string `json:"limit"`
	Offset          string `json:"offset"`
	Typologie       string `json:"typologie,omitempty"`
	Categorie       string `json:"categorie,omitempty"`
	CodePostal      string `json:"code_postal,omitempty"`
	AnneeOuverture  string `json:"annee_ouverture,omitempty"`
	AnneeRenovation string `json:"annee_renovation,omitempty"`
}

func (e *greenArea) stringToGeoJSONFeature() {
	e.Feature.Geometry = &geojson.Geometry{}
	e.Feature.Properties = map[string]interface{}{
		"nom":                    e.Nom,
		"typologie":              e.Typologie,
		"categorie":              e.Categorie,
		"adresse":                strconv.Itoa(e.AdresseNumero) + " " + e.AdresseComplement + " " + e.AdresseTypeVoie + " " + e.AdresseLibelleVoie + " " + strconv.Itoa(e.CodePostal),
		"superficieTotaleReelle": e.SuperficieTotaleReelle,
		"presenceCloture":        e.PresenceCloture,
		"perimetre":              e.Perimetre,
		"anneeOuverture":         e.AnneeOuverture,
		"anneeRenovation":        e.AnneeRenovation,
		"ancienNom":              e.AncienNom,
		"anneeChangementNom":     e.AnneeChangementNom,
		"nombreEntites":          e.NombreEntites,
		"ouverture24":            e.Ouverture24,
	}
	if strings.Contains(e.GeoShape, "Polygon") == true {
		e.Feature.Type = "Polygon"
		e.Feature.Geometry.Type = "Polygon"
		json.NewDecoder(strings.NewReader(strings.TrimRight(e.GeoShape[35:], "}"))).Decode(&e.Feature.Geometry.Polygon)
	}
	if strings.Contains(e.GeoShape, "MultiPolygon") == true {
		e.Feature.Type = "MultiPolygon"
		e.Feature.Geometry.Type = "MultiPolygon"
		json.NewDecoder(strings.NewReader(strings.TrimRight(e.GeoShape[40:], "}"))).Decode(&e.Feature.Geometry.MultiPolygon)
	}
	e.GeoShape = ""
}

func (q *query) standardizeSQLQuery(values url.Values) string {
	q.Limit = "1934"
	q.Offset = "0"
	if values.Get("limit") != "" {
		q.Limit = values.Get("limit")
	}
	if values.Get("offset") != "" {
		q.Offset = values.Get("offset")
	}
	var typologie string
	if values.Get("typologie") != "" {
		typologie = "typologie = '" + values.Get("typologie") + "' AND "
		q.Typologie = values.Get("typologie")
	}
	var categorie string
	if values.Get("categorie") != "" {
		categorie = "categorie = '" + values.Get("categorie") + "' AND "
		q.Categorie = values.Get("categorie")
	}
	var codePostal string
	if values.Get("code_postal") != "" {
		codePostal = "code_postal = '" + values.Get("code_postal") + "' AND "
		q.CodePostal = values.Get("code_postal")
	}
	var anneeOuverture string
	if values.Get("annee_ouverture") != "" {
		anneeOuverture = "annee_ouverture = '" + values.Get("annee_ouverture") + "' AND "
		q.AnneeOuverture = values.Get("annee_ouverture")
	}
	var anneeRenovation string
	if values.Get("annee_renovation") != "" {
		anneeRenovation = "annee_renovation = '" + values.Get("annee_renovation") + "' AND "
		q.AnneeRenovation = values.Get("annee_renovation")
	}
	if q.Typologie != "" || q.Categorie != "" || q.CodePostal != "" || q.AnneeOuverture != "" || q.AnneeRenovation != "" {
		return strings.TrimSuffix("WHERE "+typologie+categorie+codePostal+anneeOuverture+anneeRenovation, " AND ")
	}
	return ""
}

// GreenAreasGET gets all green areas
func GreenAreasGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var q query
	u := r.URL.Query()
	s := q.standardizeSQLQuery(u)
	rows, err := database.Db.Query("SELECT identifiant_espace_vert, geo_shape, nom, typologie, categorie, adresse_numero, adresse_complement, adresse_type_voie, adresse_libelle_voie, code_postal, surface_calculee, superficie_totale_reelle, surface_horticole, presence_cloture, perimetre, annee_ouverture, annee_renovation, ancien_nom, annee_changement_nom, nombre_entites, ouverture_24h_24h, id_division, id_atelier_horticole, ida3d_enb, site_villes, id_eqpt, competence, url_plan, geo_point FROM espaces_verts " + s + " LIMIT " + q.Limit + " OFFSET " + q.Offset + ";")
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
		var greenArea greenArea
		rows.Scan(&greenArea.Identifiant, &greenArea.GeoShape, &greenArea.Nom, &greenArea.Typologie, &greenArea.Categorie, &greenArea.AdresseNumero, &greenArea.AdresseComplement, &greenArea.AdresseTypeVoie, &greenArea.AdresseLibelleVoie, &greenArea.CodePostal, &greenArea.SurfaceCalculee, &greenArea.SuperficieTotaleReelle, &greenArea.SurfaceHorticole, &greenArea.PresenceCloture, &greenArea.Perimetre, &greenArea.AnneeOuverture, &greenArea.AnneeRenovation, &greenArea.AncienNom, &greenArea.AnneeChangementNom, &greenArea.NombreEntites, &greenArea.Ouverture24, &greenArea.IdentifiantDivision, &greenArea.IdentifiantAtelierHorticole, &greenArea.Ida3dEnb, &greenArea.SiteVilles, &greenArea.IdentifiantEquipement, &greenArea.Competence, &greenArea.URLPlan, &greenArea.GeoPoint)
		greenArea.stringToGeoJSONFeature()
		res.Data = append(res.Data, greenArea)
	}
	defer rows.Close()
	res.Meta.Query = q
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// GreenAreaGET gets a green area by id
func GreenAreaGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID, _ := mux.Vars(r)["id"]
	rows, err := database.Db.Query("SELECT identifiant_espace_vert, geo_shape, nom, typologie, categorie, adresse_numero, adresse_complement, adresse_type_voie, adresse_libelle_voie, code_postal, surface_calculee, superficie_totale_reelle, surface_horticole, presence_cloture, perimetre, annee_ouverture, annee_renovation, ancien_nom, annee_changement_nom, nombre_entites, ouverture_24h_24h, id_division, id_atelier_horticole, ida3d_enb, site_villes, id_eqpt, competence, url_plan, geo_point FROM espaces_verts WHERE identifiant_espace_vert = '" + ID + "' LIMIT 1;")
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
	var greenArea greenArea
	for rows.Next() {
		rows.Scan(&greenArea.Identifiant, &greenArea.GeoShape, &greenArea.Nom, &greenArea.Typologie, &greenArea.Categorie, &greenArea.AdresseNumero, &greenArea.AdresseComplement, &greenArea.AdresseTypeVoie, &greenArea.AdresseLibelleVoie, &greenArea.CodePostal, &greenArea.SurfaceCalculee, &greenArea.SuperficieTotaleReelle, &greenArea.SurfaceHorticole, &greenArea.PresenceCloture, &greenArea.Perimetre, &greenArea.AnneeOuverture, &greenArea.AnneeRenovation, &greenArea.AncienNom, &greenArea.AnneeChangementNom, &greenArea.NombreEntites, &greenArea.Ouverture24, &greenArea.IdentifiantDivision, &greenArea.IdentifiantAtelierHorticole, &greenArea.Ida3dEnb, &greenArea.SiteVilles, &greenArea.IdentifiantEquipement, &greenArea.Competence, &greenArea.URLPlan, &greenArea.GeoPoint)
	}
	defer rows.Close()
	if greenArea.Identifiant != 0 {
		res.Data = append(res.Data, greenArea)
	}
	res.Meta.Query = ID
	res.Meta.ResultCount = len(res.Data)
	if res.Meta.ResultCount == 0 {
		res.Error = "Non-existing ID: " + ID
		res.Message = "Not found"
	}
	json.NewEncoder(w).Encode(res)
}
