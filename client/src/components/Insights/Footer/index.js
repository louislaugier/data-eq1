import React from 'react'

class Footer extends React.Component {
  constructor(props) {
		super(props)
		this.state = {
      button: 'Tout afficher',
      style: {}
    }
  }

  render() {
    return (
			<>
        <footer style={this.state.style} className='Insights-footer'>
          <div>
            <h3>Dernières informations :</h3>
            <button onClick={() => {
              if (this.state.button === 'Tout afficher'){
                this.setState({style: { height: '33vh', position: 'absolute', top: '64.5vh', width: '72vw'}})
                this.setState({button: 'Réduire'})
              } else {
                this.setState({style: {}})
                this.setState({button: 'Tout afficher'})
              }
            }} className='Insights-footer-expand'>{this.state.button}</button>
          </div>
          <div>
            <div className='Footer-item'>
              <span className='Footer-important'>+17% d’utilisateurs de mobilité positive le mois dernier</span>
              <span>de 385 578 à 563 123 utilisateurs</span>
              <span>+15% de femmes (50 / 75 ans)</span>
              <span>+19% d’hommes de (25 / 50 ans)</span>
            </div>
            <div className='Footer-item'>
              <span className='Footer-important'>-3% de zones bruyantes</span>
              <span>Les 50 / 75 ans font des trajets plus longs</span>
              <span>+15% de femmes (50 / 75 ans)</span>
              <span>+19% d’hommes de (25 / 50 ans)</span>
            </div>
            <div className='Footer-item'>
              <span className='Footer-important'>Augmentation de la pollution depuis 2 jours à Châtelet</span>
              <span>+17% de CO2 depuis 5 jours</span>
              <span>+5% de NO2 depuis 5 jours</span>
            </div>
          </div>
          <div>
            <div className='Footer-item'>
              <span className='Footer-important'>+5% d’utilisateurs de mobilité positive les 7 derniers jours</span>
              <span>de 476 398 à 563 123 utilisateurs</span>
              <span>+11% de femmes (19 / 25 ans)</span>
              <span>+10% d’hommes de (19 / 25 ans)</span>
            </div>
            <div className='Footer-item'>
              <span className='Footer-important'>Augmentation de la pollution depuis 16 jours à Auber</span>
              <span>+18% de CO2 depuis 11 jours</span>
              <span>+5% de NO2 depuis 8 jours</span>
            </div>
            <div className='Footer-item'>
              <span className='Footer-important'>+4% d'élévation moyenne dans le 14e arrondissement</span>
              <span>Les 35 / 50 ans font des trajets plus calmes</span>
              <span>+8% de femmes (35 / 50 ans)</span>
              <span>+13% d’hommes de (35 / 50 ans)</span>
            </div>
          </div>
        </footer>
			</>
		)
  }
}

export default Footer
