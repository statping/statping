import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

export const HAS_ALL_DATA = 'HAS_ALL_DATA'
export const HAS_PUBLIC_DATA = 'HAS_PUBLIC_DATA'

export const GET_CORE = 'GET_CORE'
export const GET_SERVICES = 'GET_SERVICES'
export const GET_TOKEN = 'GET_TOKEN'
export const GET_GROUPS = 'GET_GROUPS'
export const GET_MESSAGES = 'GET_MESSAGES'
export const GET_NOTIFIERS = 'GET_NOTIFIERS'
export const GET_USERS = 'GET_USERS'

export default new Vuex.Store({
    state: {
        hasAllData: false,
        hasPublicData: false,
        core: {},
        token: null,
        services: [],
        groups: [],
        messages: [],
        users: [],
        notifiers: []
    },
    getters: {
        hasAllData: state => state.hasAllData,
        hasPublicData: state => state.hasPublicData,
        core: state => state.core,
        token: state => state.token,
        services: state => state.services,
        groups: state => state.groups,
        messages: state => state.messages,
        users: state => state.users,
        notifiers: state => state.notifiers,

        serviceById: (state) => (id) => {
            return state.services.find(s => s.id === id)
        },
        serviceByName: (state) => (name) => {
            return state.services.find(s => s.name === name)
        },
        servicesInGroup: (state) => (id) => {
            return state.services.filter(s => s.group_id === id)
        },
        servicesInOrder: (state) => () => {
          return state.services.sort((a, b) => a.order_id - b.order_id)
        },
        onlineServices: (state) => (online) => {
            return state.services.filter(s => s.online === online)
        },
        groupById: (state) => (id) => {
            return state.groups.find(g => g.id === id)
        },
        cleanGroups: (state) => () => {
          return state.groups.filter(g => g.name !== '')
        },
        userById: (state) => (id) => {
            return state.users.find(u => u.id === id)
        },
        messageById: (state) => (id) => {
            return state.messages.find(m => m.id === id)
        },
    },
    mutations: {
        setHasAllData(state, bool) {
            state.hasAllData = bool
        },
        setHasPublicData(state, bool) {
            state.hasPublicData = bool
        },
        setCore(state, core) {
            state.core = core
        },
        setToken(state, token) {
            state.token = token
        },
        setServices(state, services) {
            state.services = services
        },
        setGroups(state, groups) {
            state.groups = groups
        },
        setMessages(state, messages) {
            state.messages = messages
        },
        setUsers(state, users) {
            state.users = users
        },
        setNotifiers(state, notifiers) {
            state.notifiers = notifiers
        }
    },
    actions: {

    }
});
