import mapboxgl from 'mapbox-gl/dist/mapbox-gl'

export async function fetchData(query) {
  const data = []
  await fetch('http://localhost:80/api/v1/' + query)
    .then(res => res.json())
    .then((response) => {
      response.data.forEach(d => data.push(d))
    })
    .catch(error => console.log('Error:', error))
  return data
}

export function parseGeoJSONFeatures(data) {
  const features = []
  for (const index in data) {
    [Object.values(data[index])].forEach(record => features.push(record[1]))
  }
  return features
}

export function parseCoordinates(data) {
  const coordinates = []
  for (const index in data) {
    [Object.values(data[index])].forEach(record => coordinates.push(record[1].geometry.coordinates))
  }
  return coordinates
}

export function addHeatZonesToMap(heatZones, map){
  for (const value of Object.values(heatZones)) {
    map.addSource(value.layer.id, {
      type: 'geojson',
      data: value.data
    })
    map.addLayer(value.layer)
  }
}

export function addTrainLinesToMap(trainLines, map) {
  for (const value of Object.values(trainLines)) {
    map.addSource(value.name, {
      'type': 'geojson',
      'data': {
        'type': 'Feature',
        'properties': {
          ligne: value.name,
          nombreStations: value.nbStations || value.coordinates.length
        },
        'geometry': {
          'type': 'LineString',
          'coordinates': value.coordinates
        }
      }
    })
    map.addLayer({
      'id': value.name,
      'type': 'line',
      'source': value.name,
      'layout': {
        'line-join': 'round',
        'line-cap': 'round'
      },
      'paint': {
        'line-color': value.color,
        'line-width': 3
      }
    })
    const popup = new mapboxgl.Popup({
      closeButton: false,
      closeOnClick: false,
      className: 'Map-popup'
    })
    map.on('mouseenter', value.name, function(mousePos) {
      map.getCanvas().style.cursor = 'pointer'
      popup.setLngLat([mousePos.lngLat.lng, mousePos.lngLat.lat])
        .setHTML('Ligne : ' + mousePos.features[0].properties.ligne + '<br>Nombre de stations : ' + mousePos.features[0].properties.nombreStations)
        .addTo(map)
    })
    map.on('mouseleave', value.name, function() {
      map.getCanvas().style.cursor = ''
      popup.remove()
    })
  }
}

export function addCriteriaToMap(data, map) {
  for (const value of Object.values(data)) {
    map.addSource(value.layer.id, {
      type: 'geojson',
      data: {
        type: 'FeatureCollection',
        features: value.featureCollection
      }
    })
    map.addLayer(value.layer)
  }
}

export function getPopupDescriptions(map) {
  const popup = new mapboxgl.Popup({
    closeButton: false,
    closeOnClick: false,
    className: 'Map-popup'
  })
  map.on('mouseenter', 'stations-metro-RER', function(mousePos) {
    map.getCanvas().style.cursor = 'pointer'
    popup.setLngLat([mousePos.lngLat.lng, mousePos.lngLat.lat])
      .setHTML('Nom : ' + mousePos.features[0].properties.nom + '<br>Ligne : ' + mousePos.features[0].properties.ligne + '<br>Type de transport : ' + mousePos.features[0].properties.typeTransport + '<br>Adresse : ' + mousePos.features[0].properties.adresse)
      .addTo(map)
  })
  map.on('mouseleave', 'stations-metro-RER', function() {
    map.getCanvas().style.cursor = ''
    popup.remove()
  })
  map.on('mouseenter', 'espaces-verts', function(mousePos) {
    map.getCanvas().style.cursor = 'pointer'
    popup.setLngLat([mousePos.lngLat.lng, mousePos.lngLat.lat])
      .setHTML('Nom : ' + mousePos.features[0].properties.nom + '<br>Catégorie : ' + mousePos.features[0].properties.categorie + '<br>Typologie : ' + mousePos.features[0].properties.typologie + '<br>Année d\'ouverture : ' + mousePos.features[0].properties.anneeOuverture + '<br>Ouverture 24h sur 24 : ' + mousePos.features[0].properties.ouverture24 + '<br>Périmètre : ' + mousePos.features[0].properties.perimetre + '<br>Présence cloture : ' + mousePos.features[0].properties.presenceCloture)
      .addTo(map)
  })
  map.on('mouseleave', 'espaces-verts', function() {
    map.getCanvas().style.cursor = ''
    popup.remove()
  })
  map.on('mouseenter', 'stations-velib', function(mousePos) {
    map.getCanvas().style.cursor = 'pointer'
    popup.setLngLat([mousePos.lngLat.lng, mousePos.lngLat.lat])
      .setHTML('Nom : ' + mousePos.features[0].properties.nomStation + '<br>Dernière actualisation : ' + mousePos.features[0].properties.derniereActualisation + '<br>Bornettes libres : ' + mousePos.features[0].properties.nombreBornettesLibres + '<br>Capacité de la station : ' + mousePos.features[0].properties.capaciteStation + '<br>Vélos électriques disponibles : ' + mousePos.features[0].properties.velosElectriquesDisponibles + '<br>Vélos mécaniques disponibles : ' + mousePos.features[0].properties.velosMecaniquesDisponibles + '<br>Station en fonctionnement : ' + mousePos.features[0].properties.stationFonctionnement + '<br>Retour du Vélib possible : ' + mousePos.features[0].properties.retourVelibPossible)
      .addTo(map)
  })
  map.on('mouseleave', 'stations-velib', function() {
    map.getCanvas().style.cursor = ''
    popup.remove()
  })
}


export function dataVisibility(data, map, bool) {
  if (bool === true && data !== 'aucune') {
    map.setLayoutProperty(data, 'visibility', 'visible')
  } else if (bool === false){
    map.setLayoutProperty(data, 'visibility', 'none')
  }
  if (data === 'espaces-verts') {
    this.setState({espacesVerts: bool})
  } else if (data === 'stations-exterieures') {
    this.setState({stationsExterieures: bool})
  } else if (data === 'stations-velib') {
    this.setState({stationsVelib: bool})
  } else if (data === 'aucune') {
    this.setState({heatZone: 'Aucune'})
    map.setLayoutProperty('qualite-air', 'visibility', 'none')
    map.setLayoutProperty('nuisances-sonores', 'visibility', 'none')
    map.setLayoutProperty('nuisances-visuelles', 'visibility', 'none')
  } else if (data === 'qualite-air') {
    this.setState({heatZone: 'Qualité de l\'air', heatZoneGradientClassName: 'Qualite-air', heatZoneUnit: 'indice'})
    map.setLayoutProperty('nuisances-sonores', 'visibility', 'none')
    map.setLayoutProperty('nuisances-visuelles', 'visibility', 'none')
  } else if (data === 'nuisances-sonores') {
    this.setState({heatZone: 'Nuisances sonores', heatZoneGradientClassName: 'Nuisances-sonores', heatZoneUnit: 'dBA'})
    map.setLayoutProperty('qualite-air', 'visibility', 'none')
    map.setLayoutProperty('nuisances-visuelles', 'visibility', 'none')
  } else if (data === 'nuisances-visuelles') {
    this.setState({heatZone: 'Nuisances visuelles', heatZoneGradientClassName: 'Nuisances-visuelles', heatZoneUnit: 'm.'})
    map.setLayoutProperty('qualite-air', 'visibility', 'none')
    map.setLayoutProperty('nuisances-sonores', 'visibility', 'none')
  }
}

