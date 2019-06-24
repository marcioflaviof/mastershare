import React from 'react'
import {BrowserRouter, Switch, Route} from 'react-router-dom'

import Tables from './pages/tables'
import Login from './pages/login'
import Products from './pages/products'
import User from './pages/user'


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route path="/tables" component={Tables} />
            <Route path="/login/" component={Login} />
            <Route path="/products" component={Products} />
            <Route path="/user/" component={User} />

        </Switch>
    </BrowserRouter>
)

export default Routes