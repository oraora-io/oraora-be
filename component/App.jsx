import React from "react";

class App extends React.Component {
  state = {
    products: null
  };

  render() {
    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<LoggedIn />);
    }
  }
}

ReactDOM.render(<App />, document.getElementById('app'));