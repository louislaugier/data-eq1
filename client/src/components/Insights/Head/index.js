import React from 'react'
import Header from '../../Header'

class Head extends React.Component {
  render() {
    return (
			<>
				<Header
					title='Insights'
					subtitle={false}
					filterMenu={false}
				/>
        <div className='Insights-head'>
          <div className='Head-item'>
            <div>
              <div>
                <h3>Utilisateurs</h3>
                <span className="Head-item-subtitle">7 derniers jours</span>
              </div>
              <span>
                12% depuis la semaine dernière
              </span>
            </div>
            <b>{parseInt(17654).toLocaleString()}</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>Nombre de trajets</h3>
                <span className="Head-item-subtitle">30 derniers jours</span>
              </div>
              <span>
                8% depuis le mois dernier
              </span>
            </div>
            <b>{parseInt(1563123).toLocaleString()}</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>CO<sub>2</sub> économisé (g.)</h3>
                <span className="Head-item-subtitle">30 derniers jours</span>
              </div>
              <span>
                8% depuis le mois dernier
              </span>
            </div>
            <b>{parseInt(768652).toLocaleString()}</b>
          </div>
        </div>
			</>
		)
  }
}

export default Head
