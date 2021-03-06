import React from 'react'
import logo from '../../logo.svg'
import {BrowserRouter as Route, Switch, Link} from 'react-router-dom'
import Login from '../Login'
import Map from '../Map'
import Data from '../Data'
import Insights from '../Insights'
import Users from '../Users'

const selected = ['3.5px solid white', '#3F4049', 'relative', '-3px']

class Sidebar extends React.Component {
	constructor(props) {
    super(props)
    this.state = {
      map: '',
      data: '',
      insights: '',
      users: '',
      border: '',
      menuStyle: {
        opacity: 0.5,
        pointerEvents: 'none'
      },
      logoutDisplay: 'none',
    }
  }

  componentDidUpdate(prevProps) {
    if (prevProps.enabled !== this.props.enabled) {
      if (this.props.enabled) {
        this.setState({map: selected, border: 'Border', menuStyle: {}, logoutDisplay: ''})
      }
    }
  }

  render() {
    let logoLink = ''
    if (this.props.enabled) {
      logoLink = '/carte'
    } else {
      logoLink = '/connexion'
    }
    return (
      <>
        <Switch>
          <Route path='/carte' component={Map}/>
          <Route path='/donnees' component={Data}/>
          <Route path='/insights' component={Insights}/>
          <Route path='/voyageurs' component={Users}/>
          <Route path='/connexion' component={Login}/>
        </Switch>
        <div className='App-menu'>
          <Link to={logoLink}>
            <span onClick={() => {
              if (this.props.enabled) {
                this.setState({map: selected, data: '', insights: '', users: ''})
                this.props.change(<Map/>)
              }
            }} className='Home-link'>
                <img src={logo} className='App-logo' alt='Logo'/>
                <h1>RATP</h1>
            </span>
          </Link>
          <ul style={this.state.menuStyle} className='Menu-tabs'>
            <Link to='/carte'>
              <li style={{borderLeft: this.state.map[0], background: this.state.map[1]}} className='Map'>
                <span style={{position: this.state.map[2], left: this.state.map[3]}} onClick={() => {
                  if (this.props.enabled) {
                    this.setState({map: selected, data: '', insights: '', users: ''})
                    this.props.change(<Map/>)
                  }
                }}>
                Carte
                </span>
              </li>
            </Link>
            <Link to='/donnees'>
              <li style={{borderLeft: this.state.data[0], background: this.state.data[1]}} className='Data'>
                <span style={{position: this.state.data[2], left: this.state.data[3]}} onClick={() => {
                  if (this.props.enabled) {
                    this.setState({map: '', data: selected, insights: '', users: ''})
                    this.props.change(<Data/>)
                  }
                }}>
                  Données
                </span>
              </li>
            </Link>
            <Link to='/insights'>
              <li style={{borderLeft: this.state.insights[0], background: this.state.insights[1]}} className='Insights'>
                <span style={{position: this.state.insights[2], left: this.state.insights[3]}} onClick={() => {
                  if (this.props.enabled) {
                    this.setState({map: '', data: '', insights: selected, users: ''})
                    this.props.change(<Insights/>)
                  }
                }}>
                  Insights
                </span>
              </li>
            </Link>
            <Link to='/voyageurs'>
              <li style={{borderLeft: this.state.users[0], background: this.state.users[1]}} className={'Users ' + this.state.border}>
                <span style={{position: this.state.users[2], left: this.state.users[3]}} onClick={() => {
                  if (this.props.enabled) {
                    this.setState({map: '', data: '', insights: '', users: selected})
                    this.props.change(<Users/>)
                  }
                }}>
                  Voyageurs
                </span>
              </li>
            </Link>
            <Link to='/connexion'>
              <li style={{display: this.state.logoutDisplay}} className='Logout'>
                <span onClick={() => {
                  if (this.props.enabled) {
                    this.props.change(
                      <Login
                        login={() => this.props.change(<Map/>)}
                      />
                    )
                    this.setState({border: '', menuStyle: { opacity: 0.5, pointerEvents: 'none' }, logoutDisplay: 'none', map: '', data: '', insights: '', users: ''})
                  }
                }}>
                    Déconnexion
                  </span>
              </li>
            </Link>
          </ul>
        </div>
      </>
    )
  }
}

export default Sidebar
