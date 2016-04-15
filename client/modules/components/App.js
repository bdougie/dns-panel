import React from 'react'
import {IndexLink, Link} from 'react-router'
import Title from 'react-title-component'

export default React.createClass({
  render() {
    return (
      <div>
        <Title render='Awesome App'/>
        <h1>Special DNS Panel</h1>
        <ul>
          <li><IndexLink to='/'>Home</IndexLink></li>
          <li><Link to='/records'>View Records</Link></li>
        </ul>
        {this.props.children}
      </div>
    )
  }
})

