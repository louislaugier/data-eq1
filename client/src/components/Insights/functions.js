import {fetchData} from '../Map/functions'

export async function getCriteriaValues(criteria) {
  let res = []
  let data = []
  switch (criteria) {
    case 'qualite-air':
      res = await fetchData('arrondissements-qualite-air')
      res[0].features.forEach(feature => {
        data.push(feature.properties.valeur_moy_indice)
      })
      return data
    case 'nuisances-sonores':
      res = await fetchData('zones-nuisances-sonores')
      res[0].features.forEach(feature => {
        data.push(feature.properties.decibels)
      })
      return data
    default:
      data = await fetchData('indices-elevation')
      data[0].Elevation.forEach((num,index) => data[0].Elevation[index] = Math.round(100 * num) / 100)
      return data[0].Elevation
  }
}
