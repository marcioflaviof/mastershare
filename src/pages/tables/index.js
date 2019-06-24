import React, { Component } from 'react'
import api from '../../services/api'
import './styles.css'
import apiheader from '../../services/apipost';

export default class Main extends Component {
    state = {
        tables: [],
        
    }


    componentDidMount() {
        this.loadTables()

    }

    registerTables = async () => {
        await apiheader.post('/table/', {
            master: "Victinho",
            users: [
                {
                    name: "Otavio",
                    email: "otavio@hotamil.com",
                    pass: "123",
                    account: -456
                },
                {
                    name: "Daniel",
                    email: "daniel@hotamil.com",
                    pass: "123",
                    account: -123
                }
            ],
            products: [
                {
                    name: "Galinha Azul",
                    price: 42
                }
            ],
            account: 1
        }).then(res =>{
            console.log(res)
            console.log(res.data)
        })

    }

    loadTables = async () => {
        const response = await api.get("/table/")

        this.setState({tables: response.data})
    }


    render() {
        const {tables} = this.state

        return (
            <div className="table-list">
                {tables.map(table => (
                    <article key={table._id}>
                        <strong>Mestre: {table.Master}</strong>
                        <p id="destaque">Usuarios:</p>
                        {table.Users.map(user =>(
                            <p><strong>Nome: </strong>{user.name}: <br />
                            <strong>Email: </strong> {user.email}, <br />
                            <h3>A conta de {user.name} foi de: </h3> R$ {Number(user.bill).toFixed(2)}

                        </p>
                        ))}
                        <p id="destaque">Produtos:</p>
                        {table.Products.map(product => (
                            <p>{product.name}: {Number(product.price).toFixed(2)}</p>
                        ))}

                    </article>
                ))}
            </div>
        )
    }
}