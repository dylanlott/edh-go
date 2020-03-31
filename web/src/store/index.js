import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    notification: {
      show: false,
      message: ''
    },
    socket: {
      connected: false
    },
    user: {},
    board: {},
    notifications: {},
    game: {}
  },
  mutations: {
    SOCKET_CONNECT (state) {
      state.socket.connected = true
    },
    SOCKET_DISCONNECT (state) {
      state.socket.connected = false
    }
  },
  actions: {
  },
  modules: {
  }
})
