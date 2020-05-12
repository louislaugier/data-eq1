export default {
  qualiteAir: {
    query: 'zones-qualite-air',
    data: {},
    layer: {
      id: 'qualite-air',
      type: 'heatmap',
      source: 'qualite-air',
      minzoom: 11,
      maxzoom: 13,
      paint: {
        'heatmap-weight': {
          'property': 'valeur_moy_indice',
          'stops': [
            [1, 0.5],
            [2, 1],
            [3, 1.5],
            [4, 2],
            [5, 2.5],
            [6, 3],
          ]
        },
        'heatmap-intensity': 0.2,
        'heatmap-color': [
          'interpolate',
          ['linear'],
          ['heatmap-density'],
          0,'rgba(255, 255, 255, 0)',
          0.2,'#FEE5D9',
          0.4,'#FCBBA1',
          0.6,'#FC9272',
          0.8,'#FB6A4A',
          0.9,'#DE2D26',
          1,'#A50F15'
        ],
        'heatmap-radius': 60
      }
    }
  },
  nuisancesSonores: {
    query: 'zones-nuisances-sonores',
    data: {},
    layer: {
      id: 'nuisances-sonores',
      type: 'heatmap',
      source: 'nuisances-sonores',
      minzoom: 11,
      maxzoom: 13,
      paint: {
        'heatmap-weight': {
          'property': 'decibels',
          'stops': [
            [1, 0.5],
            [2, 1],
            [3, 1.5],
            [4, 2],
            [5, 2.5],
            [6, 3],
          ]
        },
        'heatmap-intensity': 1,
        'heatmap-color': [
          'interpolate',
          ['linear'],
          ['heatmap-density'],
          0,'rgba(255, 255, 255, 0)',
          0.2,'#F2F0F7',
          0.4,'#DADAEB',
          0.6,'#BCBDDC',
          0.8,'#9E9AC8',
          0.9,'#756BB1',
          1,'#54278F'
        ],
        'heatmap-radius': 120
      }
    }
  },
  nuisancesVisuelles: {
    query: 'zones-elevation',
    data: {},
    layer: {
      id: 'nuisances-visuelles',
      type: 'heatmap',
      source: 'nuisances-visuelles',
      minzoom: 11,
      maxzoom: 13,
      paint: {
        'heatmap-weight': {
          'property': 'elevation',
          'stops': [
            [1, 0.5],
            [2, 1],
            [3, 1.5],
            [4, 2],
            [5, 2.5],
            [6, 3],
          ]
        },
        'heatmap-intensity': 1,
        'heatmap-color': [
          'interpolate',
          ['linear'],
          ['heatmap-density'],
          0,'rgba(255, 255, 255, 0)',
          0.2,'#EFF3FF',
          0.4,'#C6DBEF',
          0.6,'#9ECAE1',
          0.8,'#6BAED6',
          0.9,'#3182BD',
          1,'#08519C'
        ],
        'heatmap-radius': 120
      }
    }
  }
}


