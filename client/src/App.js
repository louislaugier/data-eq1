import React from 'react'
import './App.css'
import {Redirect} from 'react-router-dom'
import Sidebar from './components/Sidebar'
import Login from './components/Login'
import Map from './components/Map'

class App extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      current: <Login
        login={() => this.setState({current: <Map/>, sidebar: true})}
      />,
      sidebar: false
    }
  }

  render() {
    return (
      <>
        <Redirect to="/connexion" />
        <div className='App'>
          <Sidebar
            enabled={this.state.sidebar}
            change={(component) => {
              this.setState({current: component})
              if (component.type.name === 'Login') {
                this.setState({sidebar: false})
              } else {
                this.setState({sidebar: true})
              }
            }}
          />
          <div className='App-container'>
            {this.state.current}
          </div>
        </div>
      </>
    )
  }
}

export default App
