import axios from 'axios'

class Api {
  constructor() {

  }

  async root () {
    return axios.get('/api').then(response => (response.data))
  }

  async services () {
    return axios.get('/api/services').then(response => (response.data))
  }

  async service (id) {
    return axios.get('/api/services/'+id).then(response => (response.data))
  }

  async groups () {
    return axios.get('/api/groups').then(response => (response.data))
  }

  async group (id) {
    return axios.get('/api/groups/'+id).then(response => (response.data))
  }

}
const api = new Api()
export default api
