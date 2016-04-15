import React from 'react'
import Title from 'react-title-component'
import styles from './styles.css'
import {Link} from 'react-router'

export default React.createClass({
  render() {
    return (
      <div className={styles.home}>
        <Title render={prev => `${prev} | Home`}/>
        <p className={styles.comment}>"An experience you shall never forget"</p>
        <Link className={styles.link} to='/records'>Click to View the Records</Link>
      </div>
    )
  }
})
