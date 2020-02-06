import axios from 'axios'

const qs = require('querystring')
const tokenKey = "statping_user";

class Api {
  constructor() {

  }

  async core () {
    return axios.get('/api').then(response => (response.data))
  }

  async core_save (obj) {
    return axios.post('/api/core', obj).then(response => (response.data))
  }

  async setup_save (data) {
    return axios.post('/api/setup', qs.stringify(data)).then(response => (response.data))
  }

  async services () {
    return axios.get('/api/services').then(response => (response.data))
  }

  async service (id) {
    return axios.get('/api/services/'+id).then(response => (response.data))
  }

  async service_create (data) {
    return axios.post('/api/services', data).then(response => (response.data))
  }

  async service_update (data) {
    return axios.post('/api/services/'+data.id, data).then(response => (response.data))
  }

  async service_hits (id, start, end, group) {
    return axios.get('/api/services/'+id+'/data?start=' + start + '&end=' + end + '&group=' + group).then(response => (response.data))
  }

  async service_failures (id, start, end) {
    return axios.get('/api/services/'+id+'/failures?start=' + start + '&end=' + end).then(response => (response.data))
  }

  async service_delete (id) {
    return axios.delete('/api/services/'+id).then(response => (response.data))
  }

  async services_reorder (data) {
    return axios.post('/api/reorder/services', data).then(response => (response.data))
  }

  async groups () {
    return axios.get('/api/groups').then(response => (response.data))
  }

    async groups_reorder (data) {
        return axios.post('/api/reorder/groups', data).then(response => (response.data))
    }

  async group_delete (id) {
    return axios.delete('/api/groups/'+id).then(response => (response.data))
  }

  async group_create (data) {
    return axios.post('/api/groups', data).then(response => (response.data))
  }

  async group_update (data) {
    return axios.post('/api/groups/'+data.id, data).then(response => (response.data))
  }

  async users () {
    return axios.get('/api/users').then(response => (response.data))
  }

  async user_create (data) {
    return axios.post('/api/users', data).then(response => (response.data))
  }

  async user_update (data) {
    return axios.post('/api/users/'+data.id, data).then(response => (response.data))
  }

  async user_delete (id) {
    return axios.delete('/api/users/'+id).then(response => (response.data))
  }

  async messages () {
    return axios.get('/api/messages').then(response => (response.data))
  }

  async message_create (data) {
    return axios.post('/api/messages', data).then(response => (response.data))
  }

  async message_update (data) {
    return axios.post('/api/messages/'+data.id, data).then(response => (response.data))
  }

  async message_delete (id) {
    return axios.delete('/api/messages/'+id).then(response => (response.data))
  }

  async group (id) {
    return axios.get('/api/groups/'+id).then(response => (response.data))
  }

  async notifiers () {
    return axios.get('/api/notifiers').then(response => (response.data))
  }

  async notifier_save (data) {
    return axios.post('/api/notifier/'+data.method, data).then(response => (response.data))
  }

  async notifier_test (data) {
    return axios.post('/api/notifier/'+data.method+'/test', data).then(response => (response.data))
  }

    async integrations () {
        return axios.get('/api/integrations').then(response => (response.data))
    }

    async integration (name) {
        return axios.get('/api/integrations/'+name).then(response => (response.data))
    }

    async integration_save (data) {
        return axios.post('/api/integrations/'+data.name, data).then(response => (response.data))
    }

    async renewApiKeys () {
        return axios.get('/api/renew').then(response => (response.data))
    }

    async cache () {
        return axios.get('/api/cache').then(response => (response.data))
    }

    async clearCache () {
        return axios.get('/api/clear_cache').then(response => (response.data))
    }

    async logs () {
        return axios.get('/api/logs').then(response => (response.data))
    }

    async logs_last () {
        return axios.get('/api/logs/last').then(response => (response.data))
    }

    async theme () {
        return axios.get('/api/theme').then(response => (response.data))
    }

    async theme_generate (create=true) {
      if (create) {
          return axios.get('/api/theme/create').then(response => (response.data))
      } else {
          return axios.delete('/api/theme').then(response => (response.data))
      }
    }

    async theme_save (data) {
        return axios.post('/api/theme', data).then(response => (response.data))
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

  async scss_base () {
    return await axios({
      url: '/scss/base.scss',
      method: 'GET',
      responseType: 'blob'
    }).then((response) => {
      const reader = new window.FileReader();
      return reader.readAsText(response.data)
    })
  }

  token () {
      const tk = localStorage.getItem(tokenKey)
      if (!tk) {
          return {};
      }
    return JSON.parse(tk);
  }

  authToken () {
    let user = JSON.parse(localStorage.getItem(tokenKey));
    if (user && user.token) {
      return { 'Authorization': 'Bearer ' + user.token };
    } else {
      return {};
    }
  }

  async allActions (...all) {
    await axios.all([all])
  }

}
const api = new Api()
export default api
