import React, { Component, Proptypes } from 'react'
import Title from 'react-title-component'
import { records } from '../api/handler.js'

export default React extends Component {
  static propTypes = {
    records: PropTypes.array
  };

  render() {
    return (
      <div>
        <Title render={prev => `${prev} | Records!`}/>
        <p>Records will go here</p>
      </div>
    )
  }

  renderRecords() {
    const { records } = this.props;

    return (
      <ul>
        {records.map(r =>
          <li key={r.id}>
            {r.name}
          </li>
        )}
      </ul>
    );
  }
}

export default Transmit.createContainer(Browse, {
  // For some reason you need this here or it doesn't render correctly
  initialVariables: {
  },

  fragments: {
    records() {
      return records().then(data => data);
    }
  }
});
