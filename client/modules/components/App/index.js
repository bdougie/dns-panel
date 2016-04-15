import React from 'react'
import {IndexLink, Link} from 'react-router'
import Title from 'react-title-component'
import styles from './styles.css'

export default React.createClass({
  render() {
    return (
      <div className={styles.container}>
        <div className='container'>
          <nav className={styles.navbar}>
            <IndexLink className={styles.navItem} to='/'>Home</IndexLink>
            <Link className={styles.navItem} to='/records'>View Records</Link>
          </nav>
          <Title render='DNS Panel'/>
          <div className={styles.page}>
            <div className={styles.hero}>
              <h1>Generic DNS Panel</h1>
            </div>
            <div className={styles.children}>
              {this.props.children}
            </div>
          </div>
        </div>
      </div>
    )
  }
})

