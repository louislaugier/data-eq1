CREATE TABLE espaces_verts (
    identifiant INTEGER NOT NULL,
    nom VARCHAR,
    typologie VARCHAR,
    categorie VARCHAR,
    adresse_numero INTEGER,
    AdresseComplement VARCHAR,
    AdresseTypeVoie VARCHAR,
    AdresseLibelleVoie VARCHAR,
    CodePostal INTEGER,
    SurfaceCalculee INTEGER,
    SuperficieTotaleReelle INTEGER,
    SurfaceHorticole INTEGER,
    PresenceCloture VARCHAR,
    Perimetre FLOAT,
    AnneeOuverture INTEGER,
    AnneeRenovation INTEGER,
    AncienNom VARCHAR,
    AnneeChangementNom INTEGER,
    NombreEntites INTEGER,
    Ouverture24 VARCHAR,
    IdentifiantDivision INTEGER,
    IdentifiantAtelierHorticole INTEGER,
    Ida3dEnb VARCHAR,
    SiteVilles VARCHAR,
    IdentifiantEquipement VARCHAR,
    Competence VARCHAR,
    GeoShape VARCHAR NOT NULL,
    URLPlan VARCHAR,
    GeoPoINTEGER VARCHAR
);

-- Run main.go then execute the following

UPDATE espaces_verts
    SET geo_shape = REPLACE(geo_shape, '"', '''');