import React from "react";
import Product from "./Product";

export class LoggedIn extends React.Component {
    // If a user logs out we will remove their tokens and profile info
    // logout: function () {
    //   localStorage.removeItem('id_token');
    //   localStorage.removeItem('access_token');
    //   localStorage.removeItem('profile');
    //   location.reload();
    // },
    getInitialState() {
        return {
            products: [
                { "Id": 1, "Name": "Dororo", "Slug": "dororo", "Description": "Give me back my body" },
                { "Id": 2, "Name": "O Maiden", "Slug": "o-maiden", "Description": "Literature club becomes sex ed" },
                { "Id": 3, "Name": "Demon Slayer", "Slug": "demon-slayer", "Description": "Year of the basket loli" },
                { "Id": 4, "Name": "Boku No Hero", "Slug": "my-hero", "Description": "Imagine being useless until you're not" },
                { "Id": 5, "Name": "Haikyuu", "Slug": "haikyuu", "Description": "Volleyball is cool" },
                { "Id": 6, "Name": "Boruto", "Slug": "boruto", "Description": "Boruto's dad needs an anime" }
            ]
        }
    }

    // Once this components mounts, we will make a call to the API to get the product data
    componentDidMount() {
        this.serverRequest = $.get('http://localhost:3000/products', function (result) {
            this.setState({
                products: result,
            }); x
        }.bind(this));
    }

    render() {
        return (
            <div className="col-lg-12">
                <h2>ORAORA.IO</h2>
                <p>Below you'll find garbage information</p>
                <div className="row">

                    {this.state.products.map(function (product, i) {
                        return <Product key={i} product={product} />
                    })}

                </div>
            </div>
        );
    }
}