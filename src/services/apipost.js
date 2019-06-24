import axios from 'axios'

const apiheader = axios.create({ 
    baseURL: 'http://10.12.16.52:8080',
    headers: {
        'Content-Type': 'application/json'
    } })

export default apiheader