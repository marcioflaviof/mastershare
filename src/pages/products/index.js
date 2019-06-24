import React, { Component } from 'react'
import api from '../../services/api'
import apiheader from '../../services/apipost'
import './styles.css'

export default class Main extends Component {
    constructor(){
        super()
        this.state = {
            products: [],
            name: "",
            price: "",
        }    
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    componentDidMount() {
        this.loadProducts()
    }

    handleChange(event) {
        const {name, value} = event.target
        this.setState ({
            [name] : value
        })

    }

    handleSubmit(event) {   
        this.registerProducts(this.state.name, this.state.price)        
        event.preventDefault()
        window.location.reload()
    }

    registerProducts = async (name, price) => {
        await apiheader.post("/product/", {
            name: String(name),
            price: parseFloat(price)
        }).then(res =>{
            console.log(res)
            console.log(res.data)
        })

    }    

    loadProducts = async () => {
        const response = await api.get("/product/all/")
        

        this.setState({products: response.data})
    }



    render () {
        const {products} = this.state
        return (
            <div className="product-list">
                <form ref="form" onSubmit={this.handleSubmit}>
                    <input
                    name="name"
                    value={this.state.name}
                    onChange={this.handleChange}
                    placeholder="Name"
                    />
                    <br />
                    <input
                    name="price"
                    value={this.state.price}
                    onChange={this.handleChange}
                    placeholder="Price"
                    />
                    <br />
                    <button type="submit">Submit</button>



                </form>
                <p>Your product: {this.state.name} {this.state.price}</p>
                {products.map(product => (
                    <article key={product.price}>
                        <p>Nome: {product.name}</p>
                        <p>Preco: {product.price}</p>
                    </article>

                ))}

            </div>


        )
    }
}
