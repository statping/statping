import Vuex from 'vuex'
import Vue from 'vue'
import Api from "./API"

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
            notifiers: [],
            admin: false
        },
    getters: {
        hasAllData: state => state.hasAllData,
        hasPublicData: state => state.hasPublicData,
        core: state => state.core,
        token: state => state.token,
        services: state => state.services,
        groups: state => state.groups,
        messages: state => state.messages,
        incidents: state => state.incidents,
        users: state => state.users,
        notifiers: state => state.notifiers,

        isAdmin: state => state.admin,

        servicesInOrder: state => state.services.sort((a, b) => a.order_id - b.order_id),
        groupsInOrder: state => state.groups.sort((a, b) => a.order_id - b.order_id),
        groupsClean: state => state.groups.filter(g => g.name !== '').sort((a, b) => a.order_id - b.order_id),
        groupsCleanInOrder: state => state.groups.filter(g => g.name !== '').sort((a, b) => a.order_id - b.order_id).sort((a, b) => a.order_id - b.order_id),

        serviceByAll: (state) => (element) => {
            if (element % 1 === 0) {
                return state.services.find(s => s.id == element)
            } else {
                return state.services.find(s => s.permalink === element)
            }
        },
        serviceById: (state) => (id) => {
            return state.services.find(s => s.id == id)
        },
        serviceByPermalink: (state) => (permalink) => {
            return state.services.find(s => s.permalink === permalink)
        },
        servicesInGroup: (state) => (id) => {
            return state.services.filter(s => s.group_id === id).sort((a, b) => a.order_id - b.order_id)
        },
        serviceMessages: (state) => (id) => {
            return state.messages.filter(s => s.service === id)
        },
        onlineServices: (state) => (online) => {
            return state.services.filter(s => s.online === online)
        },
        groupById: (state) => (id) => {
            return state.groups.find(g => g.id === id)
        },
        cleanGroups: (state) => () => {
            return state.groups.filter(g => g.name !== '').sort((a, b) => a.order_id - b.order_id)
        },
        userById: (state) => (id) => {
            return state.users.find(u => u.id === id)
        },
        messageById: (state) => (id) => {
            return state.messages.find(m => m.id === id)
        },
    },
    mutations: {
        setHasAllData (state, bool) {
            state.hasAllData = bool
        },
        setHasPublicData (state, bool) {
            state.hasPublicData = bool
        },
        setCore (state, core) {
            state.core = core
        },
        setToken (state, token) {
            state.token = token
        },
        setServices (state, services) {
            state.services = services
        },
        setGroups (state, groups) {
            state.groups = groups
        },
        setMessages (state, messages) {
            state.messages = messages
        },
        setUsers (state, users) {
            state.users = users
        },
        setNotifiers (state, notifiers) {
            state.notifiers = notifiers
        },
        setAdmin (state, admin) {
            state.admin = admin
        },
    },
    actions: {
        async getAllServices(context) {
            const services = await Api.services()
            context.commit("setServices", services);
        },
        async loadRequired(context) {
            const core = await Api.core()
            context.commit("setCore", core);
            const groups = await Api.groups()
            context.commit("setGroups", groups);
            const services = await Api.services()
            context.commit("setServices", services);
            const messages = await Api.messages()
            context.commit("setMessages", messages)
            context.commit("setHasPublicData", true)
            // if (core.logged_in) {
            //     const notifiers = await Api.notifiers()
            //     context.commit("setNotifiers", notifiers);
            //     const users = await Api.users()
            //     context.commit("setUsers", users);
            //     const integrations = await Api.integrations()
            //     context.commit("setIntegrations", integrations);
            // }
            window.console.log('finished loading required data')
        },
        async loadAdmin(context) {
            const core = await Api.core()
            context.commit("setCore", core);
            const groups = await Api.groups()
            context.commit("setGroups", groups);
            const services = await Api.services()
            context.commit("setServices", services);
            const messages = await Api.messages()
            context.commit("setMessages", messages)
            context.commit("setHasPublicData", true)
            const notifiers = await Api.notifiers()
            context.commit("setNotifiers", notifiers);
            const users = await Api.users()
            context.commit("setUsers", users);
        }
    }
});
