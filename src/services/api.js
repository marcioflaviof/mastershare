import axios from 'axios'

const api = axios.create({ baseURL: 'http://10.12.16.52:8080' })


export default api