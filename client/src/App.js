import React from 'react'
import './App.css'
import Sidebar from './components/Sidebar'
import Login from './components/Login'
import Map from './components/Map'
import {BrowserRouter as Router, Route, Switch, Link} from 'react-router-dom'

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
      <Router>
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
      </Router>
    )
  }
}

export default App
