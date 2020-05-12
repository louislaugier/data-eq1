import React from 'react'
import Header from '../Header'
import Loader from  'react-loader'
import mapboxgl from 'mapbox-gl/dist/mapbox-gl'
import 'mapbox-gl/dist/mapbox-gl.css'
import lines, {standardizeLineBranches} from './data/lines'
import heatZones from './data/heatZones'
import criteria from './data/criteria'
import {fetchData} from './functions'
import {parseGeoJSONFeatures, parseCoordinates, addTrainLinesToMap, addHeatZonesToMap, addCriteriaToMap, getPopupDescriptions, dataVisibility} from './functions'
import Legend from './Legend'
import HeatZoneForm from './HeatZoneForm'
import CriteriaForm from './CriteriaForm'

class Map extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      heatZone: 'Aucune',
      heatZoneGradientClassName:'',
      heatZoneUnit: '',
      espacesVerts: false,
      stationsExterieures: false,
      stationsVelib: false
    }
    this.dataVisibility = dataVisibility.bind(this)
  }

  async componentDidMount() {
    this.mounted = true
    for (let value of Object.values(heatZones)) {
      let data = await fetchData(value.query)
      value.data = data[0]
    }
    for (let value of Object.values(lines)) {
      value.coordinates = parseCoordinates(await fetchData(value.query))
    }
    const trainLines = standardizeLineBranches(lines)
    for (let value of Object.values(criteria)) {
      value.featureCollection = parseGeoJSONFeatures(await fetchData(value.query))
    }
    let map = ''
    if (this.mounted) {
      mapboxgl.accessToken = 'pk.eyJ1IjoibG91aXNsOTgiLCJhIjoiY2s3OWE2cHZpMDN0NjNlbnNwdjV1b2gzcSJ9.VjiswdWuJvyw_Eqxx4Vs2g'
      map = new mapboxgl.Map({
        container: 'Map',
        style: 'mapbox://styles/mapbox/light-v10',
        center: [2.3490508924482185, 48.858064440253486],
        zoom: 11.5,
        preserveDrawingBuffer: true,
        dragPan: true,
        scrollZoom: true
      })
      this.setState({map: map})
      map.on('load', async function() {
        addHeatZonesToMap(heatZones, this)
        addTrainLinesToMap(trainLines, this)
        addCriteriaToMap(criteria, this)
        this.setLayoutProperty('qualite-air', 'visibility', 'none')
        this.setLayoutProperty('nuisances-sonores', 'visibility', 'none')
        this.setLayoutProperty('nuisances-visuelles', 'visibility', 'none')
        this.setLayoutProperty('espaces-verts', 'visibility', 'none')
        this.setLayoutProperty('stations-exterieures', 'visibility', 'none')
        this.setLayoutProperty('stations-velib', 'visibility', 'none')
      })
      getPopupDescriptions(map)
    }
  }

  componentWillUnmount(){
    this.mounted = false
  }

  render() {
    return (
      <>

        <Header title = 'Carte de Paris 2020'
          subtitle = {true}
          sub1 = 'Carte'
          sub2 = 'RATP'
          filterMenu = {false}
        />
        <div className='Map-container'>
          <Loader loaded={this.mounted}/>
          <div id='Map'/>
          <Legend
            heatZone = {this.state.heatZone}
            heatZoneGradientClassName = {this.state.heatZoneGradientClassName}
            heatZoneUnit = {this.state.heatZoneUnit}
            espacesVerts = {this.state.espacesVerts}
            stationsExterieures = {this.state.stationsExterieures}
            stationsVelib = {this.state.stationsVelib}
          />
        </div>
        <div className='Toggle-menu'>
          <HeatZoneForm
            map = {this.state.map}
            dataVisibility = {this.dataVisibility}
          />
          <CriteriaForm
            map = {this.state.map}
            dataVisibility = {this.dataVisibility}
          />
        </div>

      </>
    )
  }
}

export default Map
