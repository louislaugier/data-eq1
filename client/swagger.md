This file was converted from Swagger YAML to MD, omitting swagger meta data, tag names and models.  

# Citi+
Documentation for Citi+ REST API.

## Base URL
http://localhost:80/api/vi/

## Version: 1.0.0

**Contact information:**  
l_laugier@etu-webschoolfactory.fr  

### /zones-qualite-air

#### GET
##### Summary:

Get a GeoJSON feature collection of chosen areas with air quality data as properties

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /indices-qualite-air/{arrondissement}

#### GET
##### Summary:

Get air quality data by arrondissement number for the last 7 days

##### Description:

Use "all" for an average of all arrondissements.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| arrondissement | path | example: 15 | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /arrondissements-qualite-air

#### GET
##### Summary:

Get a GeoJSON feature collection of areas in each arrondissement with air quality data as properties

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /zones-nuisances-sonores

#### GET
##### Summary:

Get a GeoJSON feature collection of areas in each arrondissement with sound nuisance data as properties

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /indices-nuisances-sonores/{arrondissement}

#### GET
##### Summary:

Get sound nuisance data by arrondissement number for the last 7 recorded days

##### Description:

Use "all" for an average of all arrondissements.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| arrondissement | path | example: 15 | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /tables-bdd

#### GET
##### Summary:

Get all table names from database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string |
| offset | query |  | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /colonnes-table/{table}

#### GET
##### Summary:

Get column names by table name

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| table | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /rangees-table/{table}

#### GET
##### Summary:

Get string records by table name

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| table | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /zones-elevation

#### GET
##### Summary:

Get a GeoJSON feature collection of chosen areas with landscape elevation data as properties

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /indices-elevation

#### GET
##### Summary:

Get landscape elevation data for all arrondissements

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /indice-elevation/{arrondissement}

#### GET
##### Summary:

Get landscape elevation data by arrondissement number

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| arrondissement | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-clients

#### GET
##### Summary:

Get all clients

##### Description:

test

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string |
| offset | query |  | No | string |
| genre | query |  | No | string |
| titre_transport | query | example: Navigo Mois | No | string |
| frequence_transport | query | example: 1 fois par semaine | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

#### POST
##### Summary:

Create a new client

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-client/{id}

#### GET
##### Summary:

Get a client by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

#### PUT
##### Summary:

Update an existing client

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

#### DELETE
##### Summary:

Delete an existing client

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-trajets

#### GET
##### Summary:

Get all journeys

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-trajets/{type}

#### GET
##### Summary:

Get number of clients, their age and their gender for a journey type

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| type | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-clients-scoring

#### GET
##### Summary:

Get criteria ranking and scoring averages for all user journeys

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /maratp-client-scoring/{id}

#### GET
##### Summary:

Get criteria ranking and scoring from user journeys by user ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /meteo

#### GET
##### Summary:

Get weather and humidity percentage in realtime

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /espaces-verts

#### GET
##### Summary:

Get all green areas of Paris

##### Description:

test

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string |
| offset | query |  | No | string |
| typologie | query | example: Décorations sur la voie publique | No | string |
| categorie | query | example: Jardiniere | No | string |
| code_postal | query | example: 75015 | No | string |
| annee_ouverture | query | example: 1978 | No | string |
| annee_renovation | query | example: 1995 | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /espace-vert/{id}

#### GET
##### Summary:

Get a green area by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /stations-ratp

#### GET
##### Summary:

Get all RATP stations

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string |
| offset | query |  | No | string |
| outside | query |  | No | string |
| type | query |  | No | string |
| line | query |  | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /station-ratp/{id}

#### GET
##### Summary:

Get a RATP station by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /stations-velib

#### GET
##### Summary:

Get all Vélib stations

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query |  | No | string |
| offset | query |  | No | string |
| borne | query |  | No | string |
| retour | query |  | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |

### /station-velib/{id}

#### GET
##### Summary:

Get a Vélib station by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | Database error |
