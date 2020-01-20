import axios from 'axios'

const qs = require('querystring')
const tokenKey = "statping_user";

class Api {
  constructor() {

  }

  async core () {
    return axios.get('/api').then(response => (response.data))
  }

  async services () {
    return axios.get('/api/services').then(response => (response.data))
  }

  async service (id) {
    return axios.get('/api/services/'+id).then(response => (response.data))
  }

  async service_delete (id) {
    return axios.delete('/api/services/'+id).then(response => (response.data))
  }

  async groups () {
    return axios.get('/api/groups').then(response => (response.data))
  }

  async group_delete (id) {
    return axios.delete('/api/groups/'+id).then(response => (response.data))
  }

  async group_create (data) {
    return axios.post('/api/groups', data).then(response => (response.data))
  }

  async users () {
    return axios.get('/api/users').then(response => (response.data))
  }

  async user_create (id) {
    return axios.post('/api/users').then(response => (response.data))
  }

  async user_delete (id) {
    return axios.delete('/api/users/'+id).then(response => (response.data))
  }

  async messages () {
    return axios.get('/api/messages').then(response => (response.data))
  }

  async group (id) {
    return axios.get('/api/groups/'+id).then(response => (response.data))
  }

  async notifiers () {
    return axios.get('/api/notifiers').then(response => (response.data))
  }

    async renewApiKeys () {
        return axios.get('/api/renew').then(response => (response.data))
    }

  async login (username, password) {
    const f = {username: username, password: password}
    return axios.post('/api/login', qs.stringify(f))
      .then(response => (response.data))
  }

  async logout () {
    await axios.get('/api/logout').then(response => (response.data))
    return localStorage.removeItem(tokenKey)
  }

  saveToken (username, token) {
    const user = {username: username, token: token}
    localStorage.setItem(tokenKey, JSON.stringify(user));
    return user
  }

  token () {
    return JSON.parse(localStorage.getItem(tokenKey));
  }

  authToken () {
    let user = JSON.parse(localStorage.getItem(tokenKey));
    if (user && user.token) {
      return { 'Authorization': 'Bearer ' + user.token };
    } else {
      return {};
    }
  }

}
const api = new Api()
export default api
