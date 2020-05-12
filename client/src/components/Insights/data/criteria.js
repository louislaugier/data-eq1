export const criteriaSeries = [
  {
    name: 'Pollution de l\'air',
    data: [],
  },
  {
    name: 'Nuisances visuelles',
    data: []
  }
]

export const criteriaOptions = {
  colors: ['#3751FF','#FC1101'],
  fill: {
    type: 'gradient',
    gradient: {
      type: 'vertical',
      shadeIntensity: 1,
      opacityFrom: 0.7,
      opacityTo: 0.9,
      colorStops: [
        [
          {
            offset: 0,
            color: '#3751FF',
            opacity: 1
          },
          {
            offset: 100,
            color: '#4BC0AD',
            opacity: 1
          }
        ],
        [
          {
            offset: 0,
            color: '#FECE00',
            opacity: 1
          },
          {
            offset: 100,
            color: '#FC1101',
            opacity: 1
          }
        ]
      ]
    }
  },
  dataLabels: {
    enabled: false
  },
  title: {
    text: '(Indices moyens, ville de Paris)'
  },
  xaxis: {
    categories: [],
    axisTicks: {
      show: false
    }
  },
  yaxis: {
    // show: false,
    opposite: 'true'
  }
}
