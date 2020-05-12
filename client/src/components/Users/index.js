import React from 'react'
import Head from './Head'
import {fetchData} from '../Map/functions'
import redTrophy from '../../assets/red-trophy.png'
import yellowTrophy from '../../assets/yellow-trophy.png'

const emptyClient = {
  name: 'N/A',
  age: 'N/A',
  genre: 'N/A',
  code_postal: 'N/A',
  criteria: [
    {
      name: 'N/A',
      score: 'N/A'
    },
    {
      name: 'N/A',
      score: 'N/A'
    },
    {
      name: 'N/A',
      score: 'N/A'
    },
    {
      name: 'N/A',
      score: 'N/A'
    },
    {
      name: 'N/A',
      score: 'N/A'
    },
    {
      name: 'N/A',
      score: 'N/A'
    }
  ]
}

class Users extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      naDisplay: 'block',
      dataDisplay: 'none',
      clientInput: '',
      client: emptyClient,
      clientStats: {},
      globalAverageData: [
        {
          name: 'N/A',
          score: 'N/A'
        },
        {
          name: 'N/A',
          score: 'N/A'
        },
        {
          name: 'N/A',
          score: 'N/A'
        },
        {
          name: 'N/A',
          score: 'N/A'
        },
        {
          name: 'N/A',
          score: 'N/A'
        },
        {
          name: 'N/A',
          score: 'N/A'
        }
      ],
      globalAverageDataDisplay: 'block',
      userCriteriaAverages: false,
      userAverageDataDisplay: 'none',
    }
  }

  async componentDidMount() {
    const globalCriteria = await fetchData('maratp-clients-scoring')
    globalCriteria[0].forEach(obj => obj.score = parseFloat(obj.score).toLocaleString())
    this.setState({globalAverageData: globalCriteria[0]})
  }

	render() {
    const noData = (elem, i) => {
      let html = <></>
      if (elem === 'p') {
        html = <p style={{display: this.state.naDisplay}}>N/A</p>
      } else if (elem === 'b') {
        html = <b style={{display: this.state.naDisplay}}>N/A</b>
      } else if (elem === 'footer-p') {
        html = <></>
        if (this.state.userCriteriaAverages) {
          html = <p style={{display: this.state.userAverageDataDisplay}}>{this.state.client.criteria[i].name}</p>
        }
      } else if (elem === 'footer-b'){
        html = <></>
        if (this.state.userCriteriaAverages) {
          html = <b style={{display: this.state.userAverageDataDisplay}}>{this.state.client.criteria[i].score}</b>
        }
      }
      return html
    }
    return (
			<>
				<Head/>
        <div className='User-container'>
          <form onSubmit={async (event) => {
            event.preventDefault()
            if (this.state.clientInput !== '') {
              const clientData = await fetchData('maratp-client/' + this.state.clientInput)
              if (clientData[0].hasOwnProperty('genre')) {
                clientData[0].genre = clientData[0].genre.replace(clientData[0].genre.charAt(0), clientData[0].genre.charAt(0).toUpperCase())
              }
              if (clientData[0].hasOwnProperty('code_postal')) {
                clientData[0].code_postal = parseInt(clientData[0].code_postal).toLocaleString()
              }
              clientData[0].name = 'John Doe - ' + clientData[0].identifiant
              if (clientData[0].genre === 'Femme') {
                clientData[0].name = 'Jane Doe - ' + clientData[0].identifiant
              }
              const clientCriteria = await fetchData('maratp-client-scoring/' + clientData[0].identifiant)
              clientData[0].criteria = clientCriteria[0]
              this.setState({client: clientData[0], naDisplay: 'none', dataDisplay: 'block', clientStats: {
                nbJourneys: Math.floor(Math.random() * Math.floor(34)),
                typePref: 'Marche',
                caloriesLost: parseInt(Math.floor(Math.random() * Math.floor(5000))).toLocaleString(),
                kmTraveled: Math.floor(Math.random() * Math.floor(468)),
                carbonDioxydeSaved: Math.floor(Math.random() * Math.floor(1492))
              }})
            } else {
              this.setState({client: emptyClient, naDisplay: 'block', dataDisplay: 'none'})
            }
          }}>
            <label htmlFor='user'>Utilisateur</label>
            <div>
              <input onChange={(event) => {
                this.setState({clientInput: event.target.value})
              }} placeholder='ID, adresse e-mail...' type='text' name='user'/>
              <button type='submit'>
                <span className='Shuffle-icon'/>
              </button>
            </div>
          </form>
          <div className='User'>
            <div className='Profile'>
              <h3>Profil Voyageur</h3>
              <span>Connecté avec MaRATP</span>
              <span>Prénom, Nom :</span>
            <p>{this.state.client.name}</p>
              <span>Age :</span>
            <p>{this.state.client.age}</p>
              <span>Sexe :</span>
              <p>{this.state.client.genre}</p>
              <span>Code postal :</span>
              <p>{this.state.client.code_postal}</p>
            </div>
            <div className='Stats'>
              <h3>Statistiques</h3>
              <span>De l'utilisateur</span>
              <div>
                <span>Nombre de trajets :</span>
                {noData('p')}
                <p style={{display: this.state.dataDisplay}}>{this.state.clientStats.nbJourneys}</p>
                <span>Type de transport favori :</span>
                {noData('p')}
                <p style={{display: this.state.dataDisplay}}>{this.state.clientStats.typePref}</p>
                <span>Calories perdues :</span>
                {noData('p')}
                <p style={{display: this.state.dataDisplay}}>{this.state.clientStats.caloriesLost}</p>
                <span>Km parcourus :</span>
                {noData('p')}
                <p style={{display: this.state.dataDisplay}}>{this.state.clientStats.kmTraveled + 'km'}</p>
                <span>CO<sub>2</sub> économisé :</span>
                {noData('p')}
                <p style={{display: this.state.dataDisplay}}>{this.state.clientStats.carbonDioxydeSaved + 'g'}</p>
              </div>
            </div>
            <div className='Badges'>
              <h3>Badges</h3>
              <span>Acquis par l'utilisateur</span>
              {noData('p')}
              <div style={{display: this.state.dataDisplay}}>
                <img src={redTrophy} alt='Red Trophy'/>
                <span>level 5</span>
                <p>{this.state.clientStats.kmTraveled + 'km'}</p>
                <span>Distance réalisée</span>
                <img src={yellowTrophy} alt='Yellow Trophy'/>
                <span>level 3</span>
                <p>{this.state.clientStats.nbJourneys}</p>
                <span>Nombre de trajets</span>
              </div>
            </div>
            <div className='Relation'>
              <h3>Relation</h3>
              <span>De l'utilisateur</span>
              {noData('p')}
              <div style={{display: this.state.dataDisplay}}>
                  <span className='User-positive'/>
                <p>utilisateur fidèle à Citi+.</p>
                <span>Utilisation moyenne : 2 à 3 fois par jour</span>
              </div>
            </div>
          </div>
        </div>
        <div className='Criteria-ranking'>
          <div>
            <h3>Classement des critères</h3>
            <div>
              <span>Moyenne globale</span>
              <label className="Switch">
                <input onChange={() => {
                  if (this.state.userCriteriaAverages) {
                    this.setState({userCriteriaAverages: false, globalAverageDataDisplay: 'block', userAverageDataDisplay: 'none'})
                  } else {
                    this.setState({userCriteriaAverages: true, globalAverageDataDisplay: 'none', userAverageDataDisplay: 'block'})
                  }
                }} type="checkbox"/>
                <span className="Slider"></span>
              </label>
              <span>Moyenne de l'utilisateur</span>
            </div>
          </div>
          <span className='Head-item-subtitle'>Connecté avec MaRATP</span>
          <div>
            <div className='Criteria'>
              <div>
                <span>Numéro 1</span>
                {noData('footer-p', 0)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[0].name}</p>
                {noData('footer-b', 0)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[0].score}</b>
              </div>
              <div>
                <span>Numéro 2</span>
                {noData('footer-p', 1)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[1].name}</p>
                {noData('footer-b', 1)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[1].score}</b>
              </div>
              <div>
                <span>Numéro 3</span>
                {noData('footer-p', 2)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[2].name}</p>
                {noData('footer-b', 2)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[2].score}</b>
              </div>
              <div >
                <span>Numéro 4</span>
                {noData('footer-p', 3)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[3].name}</p>
                {noData('footer-b', 3)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[3].score}</b>
              </div>
              <div>
                <span>Numéro 5</span>
                {noData('footer-p', 4)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[4].name}</p>
                {noData('footer-b', 4)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[4].score}</b>
              </div>
              <div>
                <span>Numéro 6</span>
                {noData('footer-p', 5)}
                <p style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[5].name}</p>
                {noData('footer-b', 5)}
                <b style={{display: this.state.globalAverageDataDisplay}}>{this.state.globalAverageData[5].score}</b>
              </div>
            </div>
          </div>
        </div>
			</>
		)
  }
}

export default Users
