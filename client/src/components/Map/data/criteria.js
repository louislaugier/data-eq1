export default {
  espacesVerts: {
    query: 'espaces-verts',
    featureCollection: [],
    layer: {
      id: 'espaces-verts',
      type: 'fill',
      source: 'espaces-verts',
      paint: {
        'fill-color': '#0C9A00',
        'fill-opacity': 0.5
      }
    }
  },
  stationsVelib: {
    query: 'stations-velib',
    featureCollection: [],
    layer: {
      id: 'stations-velib',
      type: 'circle',
      source: 'stations-velib',
      'paint': {
        'circle-radius': 2,
        'circle-color': '#575884'
      }
    }
  },
  stationsMetroRER: {
    query: 'stations-ratp?type=train',
    featureCollection: [],
    layer: {
      id: 'stations-metro-RER',
      type: 'circle',
      source: 'stations-metro-RER',
      'paint': {
        'circle-radius': 4,
        'circle-color': '#788F97'
      }
    }
  },
  stationsExterieures: {
    query: 'stations-ratp?outside=true',
    featureCollection: [],
    layer: {
      id: 'stations-exterieures',
      type: 'circle',
      source: 'stations-exterieures',
      'paint': {
        'circle-radius': 5,
        'circle-color': 'white',
        'circle-stroke-width': 3,
        'circle-stroke-color': '#0D72CF'
      }
    }
  }
}
