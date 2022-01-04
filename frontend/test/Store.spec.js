// import Vuex from 'vuex'
// import Api from '../src/components/API'
// import thisStore from '../src/store'
//
// import { createLocalVue } from '@vue/test-utils'

// const localVue = createLocalVue()
// localVue.use(Vuex)
//
// describe('MyName test', async () => {
//     const services = [
//         { id: 1, title: 'Apple', order_id: 3 },
//         { id: 2, title: 'Orange', order_id: 2},
//         { id: 3, title: 'Carrot', order_id: 1}
//     ]
//
//     const store = new Vuex.Store(thisStore)
//
//     await store.dispatch('loadRequired')
//
//     console.log(store.getters.services)
//
//     expect(store.getters.services).toEqual(services.slice(0, 20))
// })