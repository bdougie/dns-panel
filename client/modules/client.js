import {render} from 'react-transmit'
import { Router, browserHistory } from 'react-router'
import routes from '../modules/routes'

const app = document.getElementById('app')

render(Router, { history: browserHistory, routes: routes }, app)
