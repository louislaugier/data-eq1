import React from 'react'
import Header from '../../Header'

class Head extends React.Component {
  render() {
    return (
			<>
				<Header
					title='Voyageurs'
					subtitle={false}
          filterMenu={false}
				/>
        <div className='Users-head'>
          <div className='Head-item'>
            <div>
              <div>
                <h3>Nombre de trajets</h3>
                <span className="Head-item-subtitle">Moyenne par utilisateur</span>
              </div>
              <span>
                12% depuis la semaine dernière
              </span>
            </div>
            <b>24</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>Km parcourus</h3>
                <span className="Head-item-subtitle">Moyenne par utilisateur</span>
              </div>
              <span>
                8% depuis la semaine dernière
              </span>
            </div>
            <b>304</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>CO<sub>2</sub> économisé (g)</h3>
                <span className="Head-item-subtitle">Moyenne par utilisateur</span>
              </div>
              <span>
                8% depuis la semaine dernière
              </span>
            </div>
            <b>2040</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>Km parcourus</h3>
                <span className="Head-item-subtitle">Moyenne par trajet</span>
              </div>
              <span>
                16% depuis la semaine dernière
              </span>
            </div>
            <b>11</b>
          </div>
          <div className='Head-item'>
            <div>
              <div>
                <h3>CO<sub>2</sub> économisé (g)</h3>
                <span className="Head-item-subtitle">Moyenne par trajet</span>
              </div>
              <span>
                8% depuis la semaine dernière
              </span>
            </div>
            <b>403</b>
          </div>
        </div>
      </>
		)
  }
}

export default Head
