/*eslint-env mocha*/
import expect from 'expect'
import App from './index'
import React from 'react'
import {shallow} from 'enzyme'

describe('App', () => {
  it('container', () => {
    const div = shallow(<App />).shallow()
    expect(div).toExist()
  })
})
