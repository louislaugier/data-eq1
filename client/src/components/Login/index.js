import React from 'react'
import {Link} from 'react-router-dom'
import logo from '../../login-logo.png'
import Header from '../Header'


class Login extends React.Component {
  render() {
    return (
      <>
        <Header title = 'Connexion au Dashboard'
          subtitle = {false}
          filterMenu = {false}
        />
        <form className='Login-form'>
          <div className='Login-items'>
            <div>
              <label>Adresse e-mail</label>
              <input className='Input-login' type='text' placeholder='exemple@exemple.fr'/>
            </div>
            <div>
              <label>Adresse e-mail</label>
              <input className='Input-login' type='password' placeholder='*********'/>
              <span onClick={() => alert('Merci de contacter votre administrateur système.')} href='https://ratp.fr'>Mot de passe oublié ?</span>
            </div>
            <img src={logo} className='Login-logo' alt='Logo'/>
          </div>
          <Link to='/carte'>
            <button type='button' onClick={this.props.login}>Connexion</button>
          </Link>
        </form>
      </>
    )
  }
}

export default Login
