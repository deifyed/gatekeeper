import Axios from 'axios'

import config from '~@/app.config.js'

const axios = Axios.create({
  baseURL: config.GATEKEEPER_URL,
  withCredentials: true, // Important. This tells the browser to send cookies
})

const state = () => ({
  user: null,
})

const actions = {
  async refresh({ commit }) {
    try {
      const response = await axios.request({
        url: '/userinfo',
        method: 'get',
      })

      if (response.status === 200) commit('user', response.data)
      else if (response.status === 401) console.warn('not authenticated')
    } catch (e) {
      console.error(e)
    }
  },
  async login() {
    let url = new URL(config.GATEKEEPER_URL)
    url.pathname = "/login"
    url.searchParams.set("redirect", config.BASE_URL)

    window.location.href = url.toString()
  },
  async logout({ commit }) {
    await axios.request({
      url: '/logout',
      method: 'POST'
    })
    
    commit('user', null)
  }
}

const mutations = {
  user(state, data) {
    state.user = data
  }
}

const getters = {
  isAuthenticated: state => {
    return state.user !== null
  }
}

export default {
  namespaced: true,
  actions,
  state,
  mutations,
  getters,
}
