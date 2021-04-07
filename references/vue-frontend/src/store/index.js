import { createStore } from 'vuex'

import auth from './auth'
import upstream from './upstream'

export default createStore({
  modules: {
    auth,
    upstream,
  },
})
