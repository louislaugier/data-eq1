import React from 'react'

class Legend extends React.Component {

	render() {
    let heatZoneLegend = ''
    if (this.props.heatZone !== 'Aucune') {
      heatZoneLegend =
        <>
          <h3 className={this.props.heatZoneGradientClassName}>{this.props.heatZone} :</h3>
          <span className={'Help-popup ' + this.props.heatZoneGradientClassName}>?</span>
          <div className={'Gradient ' + this.props.heatZoneGradientClassName}/>
          <div className='Gradient-measures'>
            <span>0</span>
            <span>>100</span>
            <span className='Legend-unit'>({this.props.heatZoneUnit})</span>
          </div>
        </>
    } else {
      heatZoneLegend = <p>Aucune zone de chaleur sélectionnée.</p>
    }
    let criteriaLegend = ''
    let espacesVerts = ''
    if (this.props.espacesVerts === true) {
      espacesVerts = <p className='Espaces-verts'>Espaces verts</p>
    }
    let stationsExterieures = ''
    if (this.props.stationsExterieures === true) {
      stationsExterieures = <p className='Stations-exterieures'>Stations extérieures</p>
    }
    let stationsVelib = ''
    if (this.props.stationsVelib === true) {
      stationsVelib = <p className='Stations-velib'>Stations Vélib</p>
    }
    if (this.props.espacesVerts === false && this.props.stationsVelib === false && this.props.stationsExterieures === false) {
      criteriaLegend = <p>Aucun critère sélectionné.</p>
    } else {
      criteriaLegend =
      <>
        {espacesVerts}
        {stationsExterieures}
        {stationsVelib}
      </>
    }
    return (
      <div className='Map-legend'>
        <div className='Legend-heatzone'>
          {heatZoneLegend}
        </div>
        <div className='Legend-criteria'>
          <h3>Critères sélectionnés :</h3>
          <div className='Criteria-items'>
            {criteriaLegend}
          </div>
        </div>
      </div>
	  )
  }
}

export default Legend
