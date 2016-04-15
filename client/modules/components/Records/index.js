import React, {Component, PropTypes} from 'react'
import Title from 'react-title-component'
import {records} from '../../api/handler.js'
import {createContainer} from 'react-transmit'

class Records extends Component {
  static propTypes = {
    records: PropTypes.array
  }

  render() {
    return (
      <div>
        <Title render={prev => `${prev} | Records!`}/>
        <p>Records will go here</p>
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

export default createContainer(Records, {
  initialVariables: {},
  fragments: {
    records: () => records().then(res => res),
  },
})
