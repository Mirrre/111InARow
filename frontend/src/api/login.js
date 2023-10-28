import axios from '../http/http'

export const Login = (data) => axios.post('/user/login',data)
export const Register = (data) => axios.post('/user/register',data)
