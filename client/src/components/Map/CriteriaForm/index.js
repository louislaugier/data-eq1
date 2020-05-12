import React from 'react'

class CriteriaForm extends React.Component {
	constructor(props) {
    super(props)
    this.state = {
      espacesVerts: 'hidden',
      stationsExterieures: 'hidden',
      stationsVelib: 'hidden'
    }
  }

	render() {
    return (
      <div className='Criteria-select'>
        <h3>Choix des critères :</h3>
        <ul>
          <li>
            <div className='Box'>
              <span className='Checkbox'/>
              <input onClick={() => {
                if (this.state.espacesVerts === 'hidden'){
                  this.props.dataVisibility('espaces-verts', this.props.map,true)
                  this.setState({espacesVerts: 'visible'})
                } else {
                  this.props.dataVisibility('espaces-verts', this.props.map,false)
                  this.setState({espacesVerts: 'hidden'})
                }
              }} type='checkbox' name='espaces-verts'/>
            </div>
            <label htmlFor='espaces-verts'>Espaces verts</label>
          </li>
          <li>
            <div className='Box'>
              <span className='Checkbox'/>
              <input onClick={() => {
                if (this.state.stationsExterieures === 'hidden'){
                  this.props.dataVisibility('stations-exterieures', this.props.map,true)
                  this.setState({stationsExterieures: 'visible'})
                } else {
                  this.props.dataVisibility('stations-exterieures', this.props.map,false)
                  this.setState({stationsExterieures: 'hidden'})
                }
              }} type='checkbox' name='stations-exterieures'/>
            </div>
            <label htmlFor='stations-exterieures'>Stations extérieures</label>
          </li>
          <li>
            <div className='Box'>
              <span className='Checkbox'/>
              <input onClick={() => {
                if (this.state.stationsVelib === 'hidden'){
                  this.props.dataVisibility('stations-velib', this.props.map,true)
                  this.setState({stationsVelib: 'visible'})
                } else {
                  this.props.dataVisibility('stations-velib', this.props.map,false)
                  this.setState({stationsVelib: 'hidden'})
                }
              }} type='checkbox' name='stations-velib'/>
            </div>
            <label htmlFor='stations-velib'>Stations Vélib</label>
          </li>
        </ul>
      </div>
    )
  }
}

export default CriteriaForm
