export default {
  qualiteAir: {
    name: 'qualite-air',
    query: 'indices-qualite-air',
    series: [
      {
        name: 'O3',
        data: [],
      },
      {
        name: 'NO2',
        data: []
      },
      {
        name: 'PM10',
        data: []
      },
      {
        name: 'PM25',
        data: []
      }
    ],
    options: {
      chart: {
        zoom: {
          enabled: false
        },
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        curve: 'smooth'
      },
      legend: {
        tooltipHoverFormatter: function(val, opts) {
          return val + ' - ' + opts.w.globals.series[opts.seriesIndex][opts.dataPointIndex] + ''
        }
      },
      markers: {
        size: 0,
        hover: {
          sizeOffset: 6
        }
      },
      xaxis: {
        categories: [
        ],
      },
      yaxis: {
        // show: false,
        opposite: 'true'
      },
      title: {
        text: 'Concentration en μg/m³'
      },
      tooltip: {
        y: [
          {
            title: {
              formatter: function (val) {
                return val
              }
            }
          },
          {
            title: {
              formatter: function (val) {
                return val
              }
            }
          },
          {
            title: {
              formatter: function (val) {
                return val
              }
            }
          },
          {
            title: {
              formatter: function (val) {
                return val
              }
            }
          }
        ]
      },
      grid: {
        borderColor: '#f1f1f1',
      }
    }
  },
  nuisancesSonores: {
    name: 'nuisances-sonores',
    query: 'indices-nuisances-sonores',
    series: [
      {
      name: 'décibels',
      data: []
      }
    ],
    options: {
      chart: {
        zoom: {
          enabled: false
        },
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        curve: 'smooth'
      },
      legend: {
        tooltipHoverFormatter: (val, opts) => {
          return val + ' - ' + opts.w.globals.series[opts.seriesIndex][opts.dataPointIndex] + ''
        }
      },
      markers: {
        size: 0,
        hover: {
          sizeOffset: 6
        }
      },
      xaxis: {
        categories: [],
      },
      yaxis: {
        // show: false,
        opposite: 'true'
      },
      title: {
        text: 'Niveau sonore moyen en décibels'
      },
      tooltip: {
        y: [
          {
            title: {
              formatter: (val) => {
                return val
              }
            }
          }
        ]
      },
      grid: {
        borderColor: '#f1f1f1',
      }
    }
  },
  nuisancesVisuelles: {
    name: 'nuisances-visuelles',
    query: 'indices-elevation',
    series: [
      {
      name: 'hauteur (m.)',
      data: []
      }
    ],
    options: {
      chart: {
        zoom: {
          enabled: false
        },
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        curve: 'smooth'
      },
      legend: {
        tooltipHoverFormatter: function(val, opts) {
          return val + ' - ' + opts.w.globals.series[opts.seriesIndex][opts.dataPointIndex] + ''
        }
      },
      markers: {
        size: 0,
        hover: {
          sizeOffset: 6
        }
      },
      xaxis: {
        categories: [],
      },
      yaxis: {
        // show: false,
        opposite: 'true'
      },
      title: {
        text: 'Hauteur moyenne du paysage en m.'
      },
      tooltip: {
        y: [
          {
            title: {
              formatter: function (val) {
                return val
              }
            }
          }
        ]
      },
      grid: {
        borderColor: '#f1f1f1',
      }
    }
  }
}
