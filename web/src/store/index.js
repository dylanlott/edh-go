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
    game: {}
  },
  mutations: {
    SOCKET_CONNECT (state) {
      state.socket.connected = true
    },
    SOCKET_DISCONNECT (state) {
      state.socket.connected = false
    },
    GAME_EVENT_JOIN (state, data) {
      state.game = data
    },
    GAME_EVENT_LEAVE (state, data) {
      state.game = {}
    }
  },
  actions: {
  },
  modules: {
  }
})
