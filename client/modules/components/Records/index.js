import React, {Component, PropTypes} from 'react'
import Title from 'react-title-component'
import http from 'http'
import {resolve} from 'react-resolver'

@resolve('records', () => {
  // records are returning empty
  return http.get('http://localhost:3130/records')
})

export default class Records extends Component {
  static propTypes = {
    records: PropTypes.object
  }

  render() {
    return (
      <div>
        <Title render={prev => `${prev} | Records!`}/>
        <p>Records will go here</p>
        {this.renderRecords}
      </div>
    )
  }

  renderRecords() {
    const {records} = this.props

    return (
      <ul>
        {records.map(r =>
          <li key={r.id}>
            {r.name}
          </li>
        )}
      </ul>
    )
  }
}

