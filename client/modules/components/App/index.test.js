/*eslint-env mocha*/
import expect from 'expect'
import App from './index'
import React from 'react'
import {render} from 'react-dom'

describe('App', () => {
  it('component mounts', () => {
    const div = document.createElement('div')
    document.body.appendChild(div)
    render(<App/>, div)
    expect(div.innerHTML).toExist
  })
})
