import React, {Component, PropTypes} from 'react'
import Title from 'react-title-component'
import {resolve} from 'react-resolver'
import axios from 'axios'

@resolve('records', () => {
  return axios.get('http://localhost:3130/records')
  .then(({data}) => data)
  .catch((err) => console.log(err))
})

export default class Records extends Component {
  static propTypes = {
    records: PropTypes.array
  }

  render() {
    return (
      <div>
        <Title render={prev => `${prev} | Records!`}/>
        <p>Records are list below</p>
        <div className='records'>
          {this.renderRecords()}
        </div>
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

