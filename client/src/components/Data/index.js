import React from 'react'
import Header from '../Header'
import Chart from 'react-apexcharts'
import criteria from './data/criteria'
import formatDate from './functions'
import {fetchData} from '../Map/functions'
import Tables from './Tables'

class Data extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
      arrondissementFilter: 'block',
      chartType: 'line',
      query: 'indices-qualite-air',
      criteria: 'Qualité de l\'air',
      arrondissement: '1er arrondissement',
      arr: '1',
      series: [],
      options: criteria.qualiteAir.options,
      sidebarData: 'Realtime',
      temperature: '',
      humidity: ''
    }
  }

  async addDataToChart(arrondissement, criteriaObject) {
    let abscissa = []
    let data = ''
    if (this.mounted) {
      await this.setState({query: criteriaObject.query, series: [], options: criteriaObject.options})
      switch (criteriaObject.name) {
        case 'nuisances-visuelles':
          abscissa.push('1er arrondissement')
          for (let i = 2; i < 21; i++) {
            abscissa.push(i + 'ème arrondissement')
          }
          this.setState({chartType: 'bar'})
          data = await fetchData(this.state.query)
          break
        default:
          this.setState({chartType: 'line'})
          for (let i = 6; i > -1; i--) {
            let date = new Date()
            date.setDate(date.getDate() - i)
            abscissa.push(formatDate(date))
          }
          data = await fetchData(this.state.query + '/' + arrondissement)
      }
      for (const [key, value] of Object.entries(criteriaObject.series)) {
        value.data = Object.entries(data[0])[key][1]
      }
      this.setState({options: {...this.state.options, xaxis: {categories: abscissa}}})
      if (criteriaObject.name === 'nuisances-visuelles') {
        criteriaObject.series[0].data.forEach((value, index) => {
          criteriaObject.series[0].data[index] = Math.round(100 * value) / 100
        })
      }
      const weather = await fetchData('meteo')
      this.setState({series: criteriaObject.series, temperature: weather[0].temperature_celsius, humidity: weather[0].humidity_percentage})
    }
  }

	componentDidMount() {
    this.mounted = true
    this.addDataToChart('1', criteria.qualiteAir)
  }

  componentWillUnmount() {
    this.mounted = false
  }

	render() {
    let subtitle = ''
    let list = ''
    let sidebarElems = ''
    if (this.state.arrondissement === 'Moyenne globale') {
      subtitle = <span className='Chart-sidebar-subtitle'>Moyenne de la ville de Paris</span>
    } else {
      subtitle = <span className='Chart-sidebar-subtitle'>{this.state.arrondissement}</span>
    }
    switch (this.state.criteria) {
      case 'Nuisances visuelles':
        sidebarElems = this.state.series.map((value, key) => (
          <li key={key} className='Chart-sidebar-item'>
            {value.name} moyenne
            <span>{Math.round(100 * value.data.reduce((a, b) => a + b, 0) / value.data.length) / 100}</span>
          </li>
        ))
        break
      default:
        sidebarElems = this.state.series.map((value, key) => (
          <li key={key} className='Chart-sidebar-item'>
            {value.name}
            <span>{value.data[6]}</span>
          </li>
        ))
    }
    list = <ul className='Chart-sidebar-list'>
      <h4>({this.state.options.title.text})</h4>
      {sidebarElems}
      <li className='Chart-sidebar-item'>
        <span className='Sidebar-temp'>
          Température
          <span>{this.state.temperature}°C</span>
        </span>
        <span className='Sidebar-hum'>
          Humidité
          <span>{this.state.humidity}%</span>
        </span>
      </li>
    </ul>
    return (
			<>
				<Header
					title='Données 2020'
					subtitle = {false}
          filterMenu = {true}
          arrondissementFilter = {this.state.arrondissementFilter}
          arrondissement = {this.state.arrondissement}
          handleCriteria = {(event) =>{
            this.setState({criteria: event.target.value})
            this.setState({arrondissementFilter: 'block'})
            switch (event.target.value) {
              case 'Nuisances sonores':
                this.setState({sidebarData: 'Données historiques', options: criteria.nuisancesSonores.options})
                this.addDataToChart(this.state.arr, criteria.nuisancesSonores)
                break
              case 'Nuisances visuelles':
                this.setState({arrondissementFilter: 'none', sidebarData: 'Realtime', options: criteria.nuisancesVisuelles.options})
                this.addDataToChart(this.state.arr, criteria.nuisancesVisuelles)
                break
              default:
                this.setState({sidebarData: 'Realtime', options: criteria.qualiteAir.options})
                this.addDataToChart(this.state.arr, criteria.qualiteAir)
            }
          }}
          handleArrondissement = {(event) => {
            let critObj = ''
            switch (this.state.criteria) {
              case 'Nuisances sonores':
                critObj = criteria.nuisancesSonores
                break
              case 'Nuisances visuelles':
                critObj = criteria.nuisancesVisuelles
                break
              default:
                critObj = criteria.qualiteAir
            }
            if (event.target.value !== 'Moyenne des arrondissements') {
              this.setState({arrondissement: event.target.value})
            }
            switch (event.target.value) {
              case 'all':
                this.setState({arr: 'all', arrondissement: 'Moyenne des arrondissements'})
                this.addDataToChart('all', critObj)
                break
              default:
                const arr = parseInt(event.target.value.charAt(0) + event.target.value.charAt(1) || '').toString()
                this.setState({arr: arr})
                this.addDataToChart(arr, critObj)
            }
          }}
				/>
        <div className = 'Chart-container'>
          <div className = 'Chart-content'>
            <h3>{this.state.criteria}</h3>
            <span>Ville de Paris</span>
            <span className='Subtitle-arrondissement'>{this.state.arrondissement || ''}</span>
            <select className='Timeframe'>
              <option value='7J'>1S</option>
              <option value='7J'>1M</option>
              <option value='7J'>6M</option>
              <option value='7J'>1A</option>
              <option value='7J'>5A</option>
            </select>
            <span className='Timeframe-icon'/>
            <Chart
              className = {'Chart '+this.state.query}
              height = {window.innerHeight / 2.5}
              options = {this.state.options}
              series = {this.state.series}
              type = {this.state.chartType}
            />
          </div>
          <div className='Chart-sidebar'>
            <h3 className={this.state.sidebarData}>{this.state.sidebarData}</h3>
            {subtitle}
            {list}
          </div>
        </div>
        <Tables/>
			</>
	  )
  }
}

export default Data
