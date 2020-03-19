import { shallowMount, mount } from '@vue/test-utils'

import FormLogin from "../../src/forms/Login.vue"

const wrapper = shallowMount(FormLogin)

describe('Login Form', () => {

    it('has a created hook', () => {
        expect(typeof FormLogin.methods.checkForm).toBe('function')
    })

    it('should login', async () => {

        expect(wrapper.vm.$data.loading).toBe(false)

        expect(wrapper.vm.$data.username).toBe('')
        expect(wrapper.vm.$data.password).toBe('')

        wrapper.setData({ username: 'admin' })
        wrapper.setData({ password: 'admin' })

        wrapper.find('button').trigger('click')

        await wrapper.vm.$nextTick()

        expect(wrapper.vm.$data.loading).toBe(true)
        done()

    })

});