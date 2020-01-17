import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

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
        onlineServices: (state) => (online) => {
            return state.services.filter(s => s.online === online)
        },
        groupById: (state) => (id) => {
            return state.groups.find(g => g.id === id)
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
