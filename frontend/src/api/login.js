import axios from '../http/http'

export const Login = (data) => axios.post('/users/login',data)
export const Register = (data) => axios.post('/users/register',data)
