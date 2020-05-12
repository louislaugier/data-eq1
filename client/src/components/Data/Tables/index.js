import React from 'react'
import {fetchData} from '../../Map/functions'
import Table from './Table'

class Tables extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
      tables: [],
      overlayDisplay: 'none',
      newTableDisplay: 'none',
      newTableColumnCount: 2,
      editTableDisplay: 'none',
      table: '',
    }
  }

	async componentDidMount() {
    this.mounted = true
    const tables = await fetchData('tables-bdd')
    if (this.mounted) {
      this.setState({tables: tables})
    }
  }

  componentWillUnmount() {
    this.mounted = false
  }

	render() {
    const tables = this.state.tables.map((table) =>
      <li key={table} className='Table-select'>
        <span className='Table-select-open'>
          {table}
        </span>
        <div className='Table-links'>
          <span onClick={() => {
            this.setState({overlayDisplay: 'block', editTableDisplay: 'block', table: <Table name={table}/>})
          }} className='Table-link-edit'>
            Modifier la table
          </span>
          <span className='Table-link-download'>Télécharger CSV</span>
        </div>
      </li>
    )
    return (
			<>
        <div style={{display: this.state.overlayDisplay}} className='Overlay'>
        </div>
				<div className='Database-tables'>
          <div className='Database-head'>
            <div className='Database-title'>
              <h3>Base de données</h3>
              <span>Liste des tables utilisées</span>
            </div>
            <button onClick={() => this.setState({overlayDisplay: 'block', newTableDisplay: 'block'})} className='New-table'>Ajouter une table</button>
          </div>
          {this.state.closeTable}
          <ul className='Table-list'>
            {tables}
          </ul>
          {this.state.table}
        </div>
        <div style={{display: this.state.newTableDisplay}} className='Create-table'>
          <span onClick={() => {
              this.setState({overlayDisplay: 'none', newTableDisplay: 'none', newTableColumnCount: 2})
          }} className='Close'/>
          <h3>Créer une nouvelle table</h3>
          <form>
              <label>Nom de la table</label>
              <input type='text'/>
              <button>Importer un fichier CSV</button>
              <hr/>
              {
                Array.apply(null, {length: this.state.newTableColumnCount}).map((_, index) => {
                  return (
                    <div key={index + 1}>
                      <div className='Name-group'>
                        <label>Nom colonne {index + 1}</label>
                        <input type='text'/>
                      </div>
                      <div>
                        <label>Type</label>
                        <select>
                          <option>BOOL</option>
                          <option>INT</option>
                          <option>FLOAT</option>
                          <option>VARCHAR</option>
                          <option>TIMESTAMP</option>
                        </select>
                      </div>
                      <div>
                        <label>Nullable</label>
                        <select>
                          <option>oui</option>
                          <option>non</option>
                        </select>
                      </div>
                      <div>
                        <label>Valeur par défaut</label>
                        <input type='text'/>
                      </div>
                    </div>
                  )
                })
              }
              <button onClick={(event) => {
                event.preventDefault()
                this.setState({newTableColumnCount: this.state.newTableColumnCount + 1})
              }} className='New-col'>Ajouter une colonne</button>
              <div>
                <button onClick={(event) => {
                  event.preventDefault()
                  this.setState({overlayDisplay: 'none', newTableDisplay: 'none', newTableColumnCount: 2})
                }}>Retour</button>
                <button onClick={(event) => {
                  event.preventDefault()
                  this.setState({overlayDisplay: 'none', newTableDisplay: 'none', newTableColumnCount: 2})
                }} type='submit'>Valider</button>
              </div>
          </form>
        </div>
        <div style={{display: this.state.editTableDisplay}} className='Edit-table'>
          <span onClick={() => {
              this.setState({overlayDisplay: 'none', editTableDisplay: 'none', table: ''})
          }} className='Close'/>
          {this.state.table}
        </div>
			</>
	  )
  }
}

export default Tables

