import axios from 'axios'

const qs = require('querystring')
const tokenKey = "statping_user";

class Api {
  constructor() {

  }

  async core() {
    return axios.get('/api').then(response => (response.data))
  }

  async core_save(obj) {
    return axios.post('/api/core', obj).then(response => (response.data))
  }

  async setup_save(data) {
    return axios.post('/api/setup', qs.stringify(data)).then(response => (response.data))
  }

  async services() {
    return axios.get('/api/services').then(response => (response.data))
  }

  async service(id) {
    return axios.get('/api/services/' + id).then(response => (response.data))
  }

  async service_create(data) {
    return axios.post('/api/services', data).then(response => (response.data))
  }

  async service_update(data) {
    return axios.post('/api/services/' + data.id, data).then(response => (response.data))
  }

  async service_hits(id, start, end, group, fill=true) {
    return axios.get('/api/services/' + id + '/hits_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
  }

    async service_ping(id, start, end, group, fill=true) {
        return axios.get('/api/services/' + id + '/ping_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
    }

    async service_failures_data(id, start, end, group, fill=true) {
        return axios.get('/api/services/' + id + '/failure_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
    }

  async service_heatmap(id, start, end, group) {
    return axios.get('/api/services/' + id + '/heatmap').then(response => (response.data))
  }

  async service_failures(id, start, end, limit = 999, offset = 0) {
    return axios.get('/api/services/' + id + '/failures?start=' + start + '&end=' + end + '&limit=' + limit+ '&offset=' + offset).then(response => (response.data))
  }

  async service_delete(id) {
    return axios.delete('/api/services/' + id).then(response => (response.data))
  }

  async services_reorder(data) {
    return axios.post('/api/reorder/services', data).then(response => (response.data))
  }

  async groups() {
    return axios.get('/api/groups').then(response => (response.data))
  }

  async groups_reorder(data) {
      window.console.log('/api/reorder/groups', data)
    return axios.post('/api/reorder/groups', data).then(response => (response.data))
  }

  async group_delete(id) {
    return axios.delete('/api/groups/' + id).then(response => (response.data))
  }

  async group_create(data) {
    return axios.post('/api/groups', data).then(response => (response.data))
  }

  async group_update(data) {
    return axios.post('/api/groups/' + data.id, data).then(response => (response.data))
  }

  async users() {
    return axios.get('/api/users').then(response => (response.data))
  }

  async user_create(data) {
    return axios.post('/api/users', data).then(response => (response.data))
  }

  async user_update(data) {
    return axios.post('/api/users/' + data.id, data).then(response => (response.data))
  }

  async user_delete(id) {
    return axios.delete('/api/users/' + id).then(response => (response.data))
  }

  async incident_updates(incident) {
    return axios.post('/api/incidents/'+incident.id+'/updates', data).then(response => (response.data))
  }

    async incident_update_create(incident, data) {
        return axios.post('/api/incidents/'+incident.id+'/updates', data).then(response => (response.data))
    }

    async incidents_service(service) {
        return axios.get('/api/services/'+service.id+'/incidents').then(response => (response.data))
    }

    async incident_create(service, data) {
        return axios.post('/api/services/'+service.id+'/incidents', data).then(response => (response.data))
    }

    async incident_delete(incident) {
        return axios.delete('/api/incidents/'+incident.id).then(response => (response.data))
    }

  async messages() {
    return axios.get('/api/messages').then(response => (response.data))
  }

  async message_create(data) {
    return axios.post('/api/messages', data).then(response => (response.data))
  }

  async message_update(data) {
    return axios.post('/api/messages/' + data.id, data).then(response => (response.data))
  }

  async message_delete(id) {
    return axios.delete('/api/messages/' + id).then(response => (response.data))
  }

  async group(id) {
    return axios.get('/api/groups/' + id).then(response => (response.data))
  }

  async notifiers() {
    return axios.get('/api/notifiers').then(response => (response.data))
  }

  async notifier_save(data) {
    return axios.post('/api/notifier/' + data.method, data).then(response => (response.data))
  }

  async notifier_test(data) {
    return axios.post('/api/notifier/' + data.method + '/test', data).then(response => (response.data))
  }

  async renewApiKeys() {
    return axios.get('/api/renew').then(response => (response.data))
  }

  async cache() {
    return axios.get('/api/cache').then(response => (response.data))
  }

  async clearCache() {
    return axios.get('/api/clear_cache').then(response => (response.data))
  }

  async logs() {
    return axios.get('/api/logs').then(response => (response.data))
  }

  async logs_last() {
    return axios.get('/api/logs/last').then(response => (response.data))
  }

  async theme() {
    return axios.get('/api/theme').then(response => (response.data))
  }

  async theme_generate(create = true) {
    if (create) {
      return axios.get('/api/theme/create').then(response => (response.data))
    } else {
      return axios.delete('/api/theme').then(response => (response.data))
    }
  }

  async theme_save(data) {
    return axios.post('/api/theme', data).then(response => (response.data))
  }

  async login(username, password) {
    const f = {username: username, password: password}
    return axios.post('/api/login', qs.stringify(f))
        .then(response => (response.data))
  }

  async logout() {
    await axios.get('/api/logout').then(response => (response.data))
    return localStorage.removeItem(tokenKey)
  }

  saveToken(username, token, admin) {
    const user = {username: username, token: token, admin: admin}
    localStorage.setItem(tokenKey, JSON.stringify(user));
    return user
  }

  async scss_base() {
    return await axios({
      url: '/scss/base.scss',
      method: 'GET',
      responseType: 'blob'
    }).then((response) => {
      const reader = new window.FileReader();
      return reader.readAsText(response.data)
    })
  }

  token() {
    const tk = localStorage.getItem(tokenKey)
    if (!tk) {
      return {};
    }
    return JSON.parse(tk);
  }

  authToken() {
    let user = JSON.parse(localStorage.getItem(tokenKey));
    if (user && user.token) {
      return {'Authorization': 'Bearer ' + user.token};
    } else {
      return {};
    }
  }

  async allActions(...all) {
    await axios.all([all])
  }

}
const api = new Api()
export default api
