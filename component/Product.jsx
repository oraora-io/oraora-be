import React from "react";

export class Product extends React.Component {
    // The upvote and downvote functions will send an API request to the backend
    // These calls too will send the access_token and be verified before a success response
    // is returned

    upvote() {
        var product = this.props.product;
        fetch('http://localhost:3000/products/' + product.Slug + '/feedback', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                vote: 1
            })
        });
    }

    downvote() {
        var product = this.props.product;
        fetch('http://localhost:3000/products/' + product.Slug + '/feedback', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                vote: -1
            })
        });
    }

    getInitialState() {
        return {
            voted: null
        }
    }

    render() {
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
            </div>
        );
    }
}