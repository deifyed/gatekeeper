import Axios from 'axios'

const axios = Axios.create({
  baseURL: 'http://localhost:4554/api/test',
  withCredentials: true, // Important. This tells the browser to send cookies
})

const actions = {
  async requestOpen() {
    const response = await axios.request({
      url: '/open',
      method: 'get',
      validateStatus: () => true
    })
    
    return response.status
  },
  async requestClosed() {
    const response = await axios.request({
      url: '/closed',
      method: 'get',
      validateStatus: () => true
    })
    
    return response.status
  }
}

export default {
  namespaced: true,
  actions,
}
