import axios from 'axios'

const qs = require('querystring');
axios.defaults.withCredentials = true

const tokenKey = "statping_auth";

class Api {
  constructor() {
    this.version = "0.91.0";
    this.commit = "b7ecf0c31b0c75c394061d2f6457a925e4440f1e";
  }

  async oauth() {
    const oauth = axios.get('api/oauth').then(response => (response.data))
    return oauth
  }

  async core() {
    const core = axios.get('api').then(response => (response.data))
    if (core.allow_reports) {
      await this.sentry_init()
    }
    return core
  }

  async core_save(obj) {
    return axios.post('api/core', obj).then(response => (response.data))
  }

  async oauth_save(obj) {
    return axios.post('api/oauth', obj).then(response => (response.data))
  }

  async setup_save(data) {
    return axios.post('api/setup', qs.stringify(data)).then(response => (response.data))
  }

  async services() {
    return axios.get('api/services').then(response => (response.data))
  }

  async service(id) {
    return axios.get('api/services/' + id).then(response => (response.data))
  }

  async service_create(data) {
    return axios.post('api/services', data).then(response => (response.data))
  }

  async service_update(data) {
    return axios.post('api/services/' + data.id, data).then(response => (response.data))
  }

  async service_hits(id, start, end, group, fill = true) {
    return axios.get('api/services/' + id + '/hits_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
  }

  async service_ping(id, start, end, group, fill = true) {
    return axios.get('api/services/' + id + '/ping_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
  }

  async service_failures_data(id, start, end, group, fill = true) {
    return axios.get('api/services/' + id + '/failure_data?start=' + start + '&end=' + end + '&group=' + group + '&fill=' + fill).then(response => (response.data))
  }

  async service_uptime(id, start, end) {
    return axios.get('api/services/' + id + '/uptime_data?start=' + start + '&end=' + end).then(response => (response.data))
  }

  async service_heatmap(id, start, end, group) {
    return axios.get('api/services/' + id + '/heatmap').then(response => (response.data))
  }

  async service_failures(id, start, end, limit = 999, offset = 0) {
    return axios.get('api/services/' + id + '/failures?start=' + start + '&end=' + end + '&limit=' + limit + '&offset=' + offset).then(response => (response.data))
  }

  async service_failures_delete(service) {
    return axios.delete('api/services/' + service.id + '/failures').then(response => (response.data))
  }

  async service_delete(id) {
    return axios.delete('api/services/' + id).then(response => (response.data))
  }

  async services_reorder(data) {
    return axios.post('api/reorder/services', data).then(response => (response.data))
  }

  async checkins() {
    return axios.get('api/checkins').then(response => (response.data))
  }

  async groups() {
    return axios.get('api/groups').then(response => (response.data))
  }

  async groups_reorder(data) {
    return axios.post('api/reorder/groups', data).then(response => (response.data))
  }

  async group_delete(id) {
    return axios.delete('api/groups/' + id).then(response => (response.data))
  }

  async group_create(data) {
    return axios.post('api/groups', data).then(response => (response.data))
  }

  async group_update(data) {
    return axios.post('api/groups/' + data.id, data).then(response => (response.data))
  }

  async users() {
    return axios.get('api/users').then(response => (response.data))
  }

  async user_create(data) {
    return axios.post('api/users', data).then(response => (response.data))
  }

  async user_update(data) {
    return axios.post('api/users/' + data.id, data).then(response => (response.data))
  }

  async user_delete(id) {
    return axios.delete('api/users/' + id).then(response => (response.data))
  }

  async incident_updates(incident) {
    return axios.get('api/incidents/' + incident.id + '/updates').then(response => (response.data))
  }

  async incident_update_create(update) {
    return axios.post('api/incidents/' + update.incident + '/updates', update).then(response => (response.data))
  }

  async incident_update_delete(update) {
    return axios.delete('api/incidents/' + update.incident + '/updates/' + update.id).then(response => (response.data))
  }

  async incidents_service(id) {
    return axios.get('api/services/' + id + '/incidents').then(response => (response.data))
  }

  async incident_create(service_id, data) {
    return axios.post('api/services/' + service_id + '/incidents', data).then(response => (response.data))
  }

  async incident_delete(incident) {
    return axios.delete('api/incidents/' + incident.id).then(response => (response.data))
  }

  async checkin(api) {
    return axios.get('api/checkins/' + api).then(response => (response.data))
  }

  async checkin_create(data) {
    return axios.post('api/checkins', data).then(response => (response.data))
  }

  async checkin_delete(checkin) {
    return axios.delete('api/checkins/' + checkin.api_key).then(response => (response.data))
  }

  async messages() {
    return axios.get('api/messages').then(response => (response.data))
  }

  async message_create(data) {
    return axios.post('api/messages', data).then(response => (response.data))
  }

  async message_update(data) {
    return axios.post('api/messages/' + data.id, data).then(response => (response.data))
  }

  async message_delete(id) {
    return axios.delete('api/messages/' + id).then(response => (response.data))
  }

  async group(id) {
    return axios.get('api/groups/' + id).then(response => (response.data))
  }

  async notifiers() {
    return axios.get('api/notifiers').then(response => (response.data))
  }

  async notifier_save(data) {
    return axios.post('api/notifier/' + data.method, data).then(response => (response.data))
  }

  async notifier_test(data, notifier) {
    return axios.post('api/notifier/' + notifier + '/test', data).then(response => (response.data))
  }

  async renewApiKeys() {
    return axios.get('api/renew').then(response => (response.data))
  }

  async logs() {
    return axios.get('api/logs').then(response => (response.data)) || []
  }

  async logs_last() {
    return axios.get('api/logs/last').then(response => (response.data))
  }

  async theme() {
    return axios.get('api/theme').then(response => (response.data))
  }

  async theme_generate(create = true) {
    if (create) {
      return axios.get('api/theme/create').then(response => (response.data))
    } else {
      return axios.delete('api/theme').then(response => (response.data))
    }
  }

  async theme_save(data) {
    return axios.post('api/theme', data).then(response => (response.data))
  }

  async import(data) {
    return axios.post('api/settings/import', data).then(response => (response.data))
  }

  async check_token(token) {
    const f = {token: token}
    return axios.post('api/users/token', qs.stringify(f)).then(response => (response.data))
  }

  async login(username, password) {
    const f = {username: username, password: password}
    return axios.post('api/login', qs.stringify(f)).then(response => (response.data))
  }

  async logout() {
    return axios.get('api/logout').then(response => (response.data))
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

  async configs() {
    return axios.get('api/settings/configs').then(response => (response.data)) || []
  }

  async configs_save(data) {
    return axios.post('api/settings/configs', data).then(response => (response.data)) || []
  }

  token() {
    return $cookies.get(tokenKey);
  }

  authToken() {
    const tk = $cookies.get(tokenKey)
    if (tk) {
      return {'Authorization': 'Bearer ' + tk};
    } else {
      return {};
    }
  }

  async github_release() {
    return fetch('https://api.github.com/repos/statping-ng/statping-ng/releases/latest').then(response => response.json())
  }

  async allActions(...all) {
    await axios.all([all])
  }

}
const api = new Api()
export default api
