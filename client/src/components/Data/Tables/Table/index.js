import React from 'react'
import MaterialTable from 'material-table'
import {fetchData} from '../../../Map/functions'

class Table extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
      columns: [],
      data: []
    }
  }

  async componentDidMount() {
    this.mounted = true
    const columns = await fetchData('colonnes-table/' + this.props.name)
    columns.forEach(columnName => {
      if (this.mounted) {
        this.setState({columns: [...this.state.columns, {title: columnName, field: columnName}]})
      }
    })
    let rows = []
    const data = await fetchData('rangees-table/' + this.props.name)
    data.forEach(record => {
      const row = {}
      this.state.columns.forEach((column, index) => {
        row[column.title] = record[index]
      })
      rows.push(row)
    })
    if (this.mounted) {
      this.setState({data: rows})
    }
  }

  componentWillUnmount() {
    this.mounted = false
  }

  render() {
    return (
      <>
        <div className='Table'>
          <MaterialTable
            title={this.props.name}
            columns={this.state.columns}
            data={this.state.data}
            editable={{
              onRowAdd: (newData) =>
                new Promise((resolve) => {
                  setTimeout(() => {
                    resolve()
                    this.setState((prevState) => {
                      const data = [...prevState.data]
                      data.push(newData)
                      return { ...prevState, data }
                    })
                  }, 600)
                }),
              onRowUpdate: (prevData, newData) =>
                new Promise((resolve) => {
                  setTimeout(() => {
                    resolve()
                    if (prevData) {
                      this.setState((prevState) => {
                        const data = [...prevState.data]
                        data[data.indexOf(prevData)] = newData
                        return { ...prevState, data }
                      })
                    }
                  }, 600)
                }),
              onRowDelete: (prevData) =>
                new Promise((resolve) => {
                  setTimeout(() => {
                    resolve()
                    this.setState((prevState) => {
                      const data = [...prevState.data]
                      data.splice(data.indexOf(prevData), 1)
                      return { ...prevState, data }
                    })
                  }, 600)
                })
            }}
          />
        </div>
      </>
    )
  }
}

export default Table



