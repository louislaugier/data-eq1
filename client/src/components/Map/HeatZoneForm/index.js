import React from 'react'

class HeatZoneForm extends React.Component {
	constructor(props) {
    super(props)
    this.state = {
      heatZone: 'Aucune'
    }
  }

	render() {
    return (
      <div className='Heatzone-select'>
          <h3>Choix de la zone de chaleur :</h3>
          <ul>
              <li>
                  <div className='Dot'>
                    <span className='Radio'/>
                    <input onClick={() => {
                      this.props.dataVisibility('aucune', this.props.map,true)
                      this.setState({heatZone: 'Aucune'})
                    }} defaultChecked type='radio' name='heatzone'/>
                  </div>
                  <label>Aucune</label>
              </li>
              <li>
                  <div className='Dot'>
                    <span className='Radio'/>
                    <input onClick={() => {
                      this.props.dataVisibility('qualite-air', this.props.map,true)
                      this.setState({heatZone: 'Qualité de l\'air'})
                    }} type='radio' name='heatzone'/>
                  </div>
                  <label>Qualité de l'air</label>
              </li>
              <li>
                  <div className='Dot'>
                    <span className='Radio'/>
                    <input onClick={() => {
                      this.props.dataVisibility('nuisances-sonores', this.props.map,true)
                      this.setState({heatZone: 'Nuisances sonores'})
                    }} type='radio' name='heatzone'/>
                  </div>
                  <label>Nuisances sonores</label>
              </li>
              <li>
                  <div className='Dot'>
                    <span className='Radio'/>
                    <input onClick={() => {
                      this.props.dataVisibility('nuisances-visuelles', this.props.map,true)
                      this.setState({heatZone: 'Nuisances visuelles'})
                    }} type='radio' name='heatzone'/>
                  </div>
                  <label>Nuisances visuelles</label>
              </li>
          </ul>
      </div>
    )
  }
}

export default HeatZoneForm
