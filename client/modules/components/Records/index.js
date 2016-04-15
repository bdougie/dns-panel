import React, {Component, PropTypes} from 'react'
import Title from 'react-title-component'
import {resolve} from 'react-resolver'
import axios from 'axios'
import styles from './styles.css'

@resolve('records', () => {
  return axios.get('http://localhost:3130/records')
  .then(({data}) => data)
  .catch((err) => console.log(err))
})

export default class Records extends Component {
  static propTypes = {
    records: PropTypes.array
  }

  _handleClick() {
    console.log('clicked')
  }

  render() {
    return (
      <div className={styles.container}>
        <Title render={prev => `${prev} | Records!`}/>
        <p className={styles.explanation} >You DNS Records are listed below</p>
        <div className={styles.records}>
          {this.renderRecords()}
        </div>
      </div>
    )
  }

  renderRecords() {
    const {records} = this.props
    const buttonColor = {
      edit: {color: '#48BBEC'},
      remove: {color: 'red'},
    }

    return (
      <div className={styles.record}>
        {records.map(r =>
          <div className={styles.values} key={r.id}>
            <div style={buttonColor.edit} onClick={this._handleClick} className={styles.button}> Edit</div>
            <div style={buttonColor.remove} onClick={this._handleClick} className={styles.button}> Remove </div>
            <div className={styles.item}>{r.domain}</div>
            <div className={styles.item}>{r.name}</div>
            <div className={styles.item}>{r.address}</div>
          </div>
        )}
      </div>
    )
  }
}

