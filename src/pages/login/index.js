import React, { Component } from 'react'
import {AsyncStorage} from 'react'
import apiheader from '../../services/apipost';

export default class Login extends Component {
    state = {
        loggedInUser: null,
    }

    componentDidMount() {
        this.signIn()
    }

    signIn = async () => {
        const response = await apiheader.post('/user/login', {
            email: "aaa@hotamil.com",
            pass: "123",
        }).then(err => console.log(err))

        const { user, token } = response.data

        await AsyncStorage({ "Codeapi:token":token, "CodeApi:user":JSON.stringify(user) })

        this.setState({loggedInUser: user})

        
    }
    

    render () {
        return (
            <div>
                <button onClick={this.signIn} title="Entrar">Entrar</button>

            </div>



        )
    }

}