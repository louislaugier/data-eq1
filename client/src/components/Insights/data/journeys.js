export const userData = {
  totals: {
    total: 0,
    women: 0,
    men: 0
  },
  zero_to_eighteen: {
    women: 0,
    men: 0
  },
  nineteen_to_twentyfive: {
    women: 0,
    men: 0
  },
  twentysix_to_thirtyfour: {
    women: 0,
    men: 0
  },
  thirtyfive_to_fifty: {
    women: 0,
    men: 0
  },
  fiftyone_to_sixtyfive: {
    women: 0,
    men: 0
  },
  sixtysix_and_over: {
    women: 0,
    men: 0
  }
}

export const journeySeries = [
  {
    name: 'Hommes',
    data: [],
  },
  {
    name: 'Femmes',
    data: [],
  },
]

export const journeyOptions = {
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
        // [
        //   {
        //     offset: 0,
        //     color: '#FF9737',
        //     opacity: 1
        //   },
        //   {
        //     offset: 100,
        //     color: '#C04B4B',
        //     opacity: 1
        //   }
        // ]
      ]
    }
  },
  title: {
    text: 'Derniers 30 jours'
  },
  // grid: {
  //   show: false,
  // },
  dataLabels: {
    enabled: false
  },
  xaxis: {
    categories: ['0-18 ans', '19-25 ans', '26-34 ans', '35-50 ans', '51-65 ans', '66 ans et+'],
    axisTicks: {
      show: false
    }
  },
  yaxis: {
    // show: false,
    opposite: 'true'
  }
}
