/*eslint-env mocha*/
import expect from 'expect'
import Home from './index'
import React from 'react'
import {render} from 'react-dom'

describe('Home', () => {
  it('component mounts', () => {
    const div = document.createElement('div')
    document.body.appendChild(div)
    render(<Home/>, div)
    expect(div.innerHTML).toExist
  })
})

