import React from 'react'
import Head from './Head'
import Charts from './Charts'
import Footer from './Footer'

class Insights extends React.Component {
  componentDidMount(){
    this.mounted = true
  }
  componentWillMount(){
    this.mounted = false
  }
	render() {
    	return (
			<>
				<Head/>
        <Charts/>
        <Footer/>
			</>
		)
  }
}

export default Insights
