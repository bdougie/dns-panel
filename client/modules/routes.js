import '../modules/styles.css'
import React from 'react'
import { Route, IndexRoute, Redirect } from 'react-router'
import App from './components/App'
import Home from './components/Home'
import NoMatch from './components/NoMatch'
import Dragon from './components/Dragon'

export default (
  <Route>
    <Route path="/" component={App}>
      <IndexRoute component={Home}/>
      <Route path="records" component={Dragon}/>
    </Route>
    <Redirect from="/not-dragon" to="/home"/>
    <Route path="*" status={404} component={NoMatch}/>
  </Route>
)
