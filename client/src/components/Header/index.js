import React from 'react'

class Header extends React.Component {

  getItems() {
    if (this.props.subtitle === true && this.props.filterMenu === false) {
      return (
        <>
          <span style={{marginBottom:'3.5vh', display:'inline-block'}}>
            {this.props.sub1}
          </span>
          <span> {this.props.sub2}</span>
        </>
      )
    }
    const arrondissements = []
    arrondissements.push(<option value={'1er arrondissement'} key={1}>1er arrondissement</option>)
    for (let i = 1; i < 21; i++) {
      arrondissements.push(<option value={i + 'ème arrondissement'} key={i}>{i}ème arrondissement</option>)
    }
    delete arrondissements[1]
    if (this.props.subtitle === false && this.props.filterMenu === true) {
      return (
        <form className='Filter-menu' style={{marginTop: '2vh'}}>
          <div style={{display: this.props.arrondissementFilter}} className='Filter-arrondissement'>
            <span>Localisation :</span>
            <select value={this.props.arrondissement} onChange={this.props.handleArrondissement}>
              <option value='all' key='all'>Moyenne des arrondissements</option>
              {arrondissements}
            </select>
            <span className='Dropdown-icon'/>
          </div>
          <div className='Filter-criteria'>
            <span>Critère :</span>
              <div className='Filter-criteria-select'>
                <select onChange={this.props.handleCriteria}>
                  <option value='Qualité de l&quot;air'>Qualité de l'air</option>
                  <option value='Nuisances sonores'>Nuisances sonores</option>
                  <option value='Nuisances visuelles'>Nuisances visuelles</option>
                </select>
                <span className='Dropdown-icon'/>
              </div>
          </div>
        </form>
      )
    }
    return ''
  }

	render() {
    return (
      <>
        <header className='App-header'>
          <h2>
              {this.props.title}
          </h2>
          {this.getItems()}
        </header>
      </>
    )
  }
}

export default Header
