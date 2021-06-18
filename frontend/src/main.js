import Vue from 'vue'
import VueApollo from 'vue-apollo'
import api from '@/gqlclient'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import './scss/custom.scss'
import VueCookies from 'vue-cookies'
import router from './router'
import App from './App.vue'
import store from './store'
import { AuthPlugin } from './auth'

Vue.use(Buefy)
Vue.use(VueCookies)
Vue.$cookies.config('7d')
Vue.config.productionTip = false
const apolloProvider = new VueApollo({
  defaultClient: api,
});
Vue.use(VueApollo)
Vue.use(AuthPlugin)

const vm = new Vue({
  router,
  store,
  provide: apolloProvider.provide(),
  render: (h) => h(App),
});
vm.$mount('#app');
