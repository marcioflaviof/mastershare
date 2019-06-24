import React, { Component } from 'react'
import api from '../../services/api'
import './styles.css'
import apiheader from '../../services/apipost'

export default class Main extends Component {
    constructor(){
        super()
            this.state = {
                user: [],
                name: "",
                emailget: "",
                emailpost: "",
                pass: "",
        }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleSubmitPost = this.handleSubmitPost.bind(this)
    }        



    handleChange(event) {
        const {name, value} = event.target
        this.setState ({
            [name] : value
        })
    }

    handleSubmit(event){
        this.loadUsers(this.state.emailget)
        event.preventDefault()
    }

    handleSubmitPost(event) {
        let user = this.state
        this.registerUsers(user.name, user.emailpost, user.pass)
        event.preventDefault()
        window.location.reload()
    }

    registerUsers = async(username, email, pass) => {
        await apiheader.post(`/user/`, {
            name: username,
            email: email,
            pass: pass
        }).catch(err => console.log(err))
    }

    loadUsers = async (email) => {
        const response = await api.get(`/user/?email=${email}`)


        this.setState({user: response.data})
    }

    render () {
        const {user} = this.state
        return (
            <div className="user-list">
                <form onSubmit={this.handleSubmit}>
                    <h1>Para procurar: </h1>
                    <br />
                    <input
                    name="emailget"
                    value={this.state.emailget}
                    onChange={this.handleChange}
                    placeholder="Digite seu Email"
                    />
                    <button type="submit">Submit</button>
                </form>
                <form onSubmit={this.handleSubmitPost}>
                    <h1>Para cadastrar usuario: </h1>
                    <br />
                    <input
                    name="name"
                    value={this.state.name}
                    onChange={this.handleChange}
                    placeholder="Digite seu nome"
                    />
                    <br />
                    <input
                    name="emailpost"
                    value={this.state.emailpost}
                    onChange={this.handleChange}
                    placeholder="Digite seu email"
                    />
                    <br />
                    <input
                    name="pass"
                    value={this.state.pass}
                    onChange={this.handleChange}
                    placeholder="Digite sua senha"
                    />
                    <br />
                    <button type="submit">Submit</button>
                </form>

                <article key={" "}>
                    <p>Nome: {user.name}</p>
                    <p>Email: {user.email}</p>
                </article>

                

            </div>


        )
    }
}