import Vue from 'vue'
import io from 'socket.io-client'
import App from './App.vue'
import router from './router'
import store from './store'

var socket = io()
// on connection of webapp
socket.on('connection', (socket) => {
  socket.emit('/hello', 'world')
})

socket.on('/', 'hello', (msg) => {
  console.log('hello world: ', msg)
})

Vue.$socket = socket
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
