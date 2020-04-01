import Vue from 'vue'
import VueSocketio from 'vue-socket.io'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

// Socket.io configuration
Vue.use(new VueSocketio({
  debug: true,
  connection: 'http://localhost:8000/socket.io',
  vuex: {
    store,
    actionPrefix: 'SOCKET_',
    mutationPrefix: 'SOCKET_'
  }
}))

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
