var App = React.createClass({
    componentWillMount: function () {
    },
    render: function () {

        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            // Should be <Home /> but neglecting login function for now
            return (<LoggedIn />);
        }
    }
});

var Home = React.createClass({
    render: function () {
        return (
            <div className="container">
                <div className="col-xs-12 jumbotron text-center">
                    <h1>Oraora</h1>
                    <p>Welcome to Oraora</p>
                    <a className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>);
    }
});

var LoggedIn = React.createClass({
    // If a user logs out we will remove their tokens and profile info
    // logout: function () {
    //     localStorage.removeItem('id_token');
    //     localStorage.removeItem('access_token');
    //     localStorage.removeItem('profile');
    //     location.reload();
    // },
    getInitialState: function () {
        return {
            products: []
        }
    },
    // Once this components mounts, we will make a call to the API to get the product data
    componentDidMount: function () {
        this.serverRequest = $.get('http://localhost:3000/products', function (result) {
            this.setState({
                products: result,
            });
        }.bind(this));
    },

    render: function () {
        console.log("Are you here?");
        return (
            <div className="col-lg-12">
                {/* <span className="pull-right"><a onClick={this.logout}>Log out</a></span> */}
                <h2>Oraora</h2>
                <p>Welcome to Oraora</p>
                <div className="row">

                    {this.state.products.map(function (product, i) {
                        return <Product key={i} product={product} />
                    })}

                </div>
            </div>);

        
    },
});

var Product = React.createClass({
    // The upvote and downvote functions will send an API request to the backend
    // These calls too will send the access_token and be verified before a success response
    // is returned
    upvote: function () {
        var product = this.props.product;
        this.serverRequest = $.post('http://localhost:3000/products/' + product.Slug + '/feedback', { vote: 1 }, function (result) {
            this.setState({ voted: "Upvoted" })
        }.bind(this));
    },
    downvote: function () {
        var product = this.props.product;
        this.serverRequest = $.post('http://localhost:3000/products/' + product.Slug + '/feedback', { vote: -1 }, function (result) {
            this.setState({ voted: "Downvoted" })
        }.bind(this));
    },
    getInitialState: function () {
        return {
            voted: null
        }
    },
    render: function () {
        return (
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading">{this.props.product.Name} <span className="pull-right">{this.state.voted}</span></div>
                    <div className="panel-body">
                        {this.props.product.Description}
                    </div>
                    <div className="panel-footer">
                        <a onClick={this.upvote} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbs-up"></span>
                        </a>
                        <a onClick={this.downvote} className="btn btn-default pull-right">
                            <span className="glyphicon glyphicon-thumbs-down"></span>
                        </a>
                    </div>
                </div>
            </div>);
    }
})

ReactDOM.render(<App />, document.getElementById('app'));