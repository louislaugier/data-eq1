import React from 'react'
import Loader from 'react-loader'
import Chart from 'react-apexcharts'
import {fetchData} from '../../Map/functions'
import {userData, journeySeries, journeyOptions} from '../data/journeys'
import {criteriaSeries, criteriaOptions} from '../data/criteria'
import {getCriteriaValues} from '../functions'

class Charts extends React.Component {
	constructor(props) {
		super(props)
    this.state = {
      mounted: false,
      data: userData,
      journeySeries: [],
      journeyOptions: {},
      criteriaSeries: [],
      criteriaOptions: {}
    }
  }

  async addDataToJourneyTypeChart(type) {
    const data = await fetchData('maratp-trajets/' + type)
    this.setState({data: data[0]})
      this.setState({journeySeries: [
        {
          name: 'Hommes',
          data: [this.state.data.zero_to_eighteen.men, this.state.data.nineteen_to_twentyfive.men, this.state.data.twentysix_to_thirtyfour.men, this.state.data.thirtyfive_to_fifty.men, this.state.data.fiftyone_to_sixtyfive.men, this.state.data.sixtysix_and_over.men],
        },
        {
          name: 'Femmes',
          data: [this.state.data.zero_to_eighteen.women, this.state.data.nineteen_to_twentyfive.women, this.state.data.twentysix_to_thirtyfour.women, this.state.data.thirtyfive_to_fifty.women, this.state.data.fiftyone_to_sixtyfive.women, this.state.data.sixtysix_and_over.women],
        },
      ]})
    this.state.mounted = true
  }

  handleCriteriaCompare = async (event) => {
    if (event.target.value.charAt(0) === 'x') {
      switch (event.target.value) {
        case 'x-nuisances-sonores':
          this.setState({criteriaSeries: [{
            name: 'Nuisances sonores',
            data: await getCriteriaValues('nuisances-sonores'),
          }, {...this.state.criteriaSeries[1]}]})
          break
        case 'x-nuisances-visuelles':
          this.setState({criteriaSeries: [{
            name: 'Nuisances visuelles',
            data: await getCriteriaValues('nuisances-visuelles'),
          }, {...this.state.criteriaSeries[1]}]})
          break
        default:
          this.setState({criteriaSeries: [{
            name: 'Qualité de l\'air',
            data: await getCriteriaValues('qualite-air'),
          }, {...this.state.criteriaSeries[1]}]})
      }
    } else {
      switch (event.target.value) {
        case 'y-nuisances-sonores':
          this.setState({criteriaSeries: [{...this.state.criteriaSeries[0]}, {
            name: 'Nuisances sonores',
            data: await getCriteriaValues('nuisances-sonores'),
          }]})
          break
        case 'y-nuisances-visuelles':
          this.setState({criteriaSeries: [{...this.state.criteriaSeries[0]}, {
            name: 'Nuisances visuelles',
            data: await getCriteriaValues('nuisances-visuelles'),
          }]})
          break
        default:
          this.setState({criteriaSeries: [{...this.state.criteriaSeries[0]}, {
            name: 'Qualité de l\'air',
            data: await getCriteriaValues('qualite-air'),
          }]})
      }
    }
  }

  async componentDidMount(){
    await this.setState({data: userData, journeySeries: journeySeries, journeyOptions: journeyOptions, criteriaSeries: criteriaSeries, criteriaOptions: criteriaOptions})
    const abcissa = []
    abcissa.push('1er arrondissement')
    for (let i = 2; i < 21; i++) {
      abcissa.push(i + 'ème arrondissement')
    }
    if (this.mounted) {
      this.setState({criteriaOptions: {...this.state.criteriaOptions, xaxis: { categories: abcissa }}})
      this.addDataToJourneyTypeChart('actif')
      this.setState({criteriaSeries: [
        {
          name: this.state.criteriaSeries[0].name,
          data: await getCriteriaValues('qualite-air'),
        },
        {
          name: this.state.criteriaSeries[1].name,
          data: await getCriteriaValues('nuisances-visuelles'),
        }
      ]})
    }
  }

  componentWillUnmount() {
    this.state.mounted = false
  }

	render() {
    return (
      <div className='Insights-charts'>
        <div className='Users-chart'>
          <select onChange={(event) => {
            this.addDataToJourneyTypeChart(event.target.value)
          }} autosize='true' className='Journey-type'>
            <option value='actif'>Utilisateurs des trajets avec l'option actif</option>
            <option value='sante'>Utilisateurs des trajets avec l'option santé</option>
            <option value='calme'>Utilisateurs des trajets avec l'option calme</option>
          </select>
          <span className='Dropdown-icon'/>
          <Chart
            className = 'Chart journeys'
            height = {window.innerHeight / 2.35}
            options = {this.state.journeyOptions}
            series = {this.state.journeySeries}
            type = 'bar'
          />
          <div className='Users-total'>
            <div>
              <h3>Utilisateurs</h3>
              <span>12%</span>
            </div>
            <div>
              <div>
                <span>Total :</span>
                <span>{this.state.data.totals.total || <Loader loaded={this.mounted}/>}</span>
              </div>
              <div>
                <span>Femmes :</span>
                <span>{parseInt(this.state.data.totals.women).toLocaleString()}</span>
              </div>
              <div>
                <span>Hommes :</span>
                <span>{parseInt(this.state.data.totals.men).toLocaleString()}</span>
              </div>
            </div>
          </div>
        </div>
        <div className='Criteria-chart'>
          <div className='Criteria-chart-head'>
            <select onChange={this.handleCriteriaCompare} autosize='true' className='Select-x'>
              <option value='x-qualite-air'>Pollution de l'air</option>
              <option value='x-nuisances-sonores'>Nuisances sonores</option>
              <option value='x-nuisances-visuelles'>Nuisances visuelles</option>
            </select>
            <span className='Dropdown-icon'/>
            <span className='Slash'> / </span>
            <select onChange={this.handleCriteriaCompare} className='Select-y'>
            <option value='y-nuisances-visuelles'>Nuisances visuelles</option>
              <option value='y-nuisances-sonores'>Nuisances sonores</option>
              <option value='y-qualite-air'>Pollution de l'air</option>
            </select>
            <span className='Dropdown-icon'/>
          </div>
          <Chart
            className = 'Chart criteria'
            height = {window.innerHeight / 2.05}
            options = {this.state.criteriaOptions}
            series = {this.state.criteriaSeries}
            type = 'bar'
          />
        </div>
      </div>
    )
  }
 }

export default Charts
