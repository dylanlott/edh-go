import Vue from 'vue';
import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import VueApollo from 'vue-apollo';
import Vuex from 'vuex'
import { split } from 'apollo-link';
import { WebSocketLink } from 'apollo-link-ws';
import { getMainDefinition } from 'apollo-utilities';

// TODO: Remove bootstrap cause it sucks
// import 'bootstrap';

// Buefy
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import './scss/custom.scss';
Vue.use(Buefy)

import router from './router';
import App from './App.vue';
import store from './store';
import { AuthPlugin } from './auth';

Vue.config.productionTip = false;

const httpLink = new HttpLink({
  uri: 'http://localhost:8080/graphql',
});
const wsLink = new WebSocketLink({
  uri: 'ws://localhost:8080/graphql',
  options: {
    reconnect: true,
  },
});
const link = split(
  ({ query }) => {
    const { kind, operation } = getMainDefinition(query);
    return kind === 'OperationDefinition' && operation === 'subscription';
  },
  wsLink,
  httpLink,
);
const apolloClient = new ApolloClient({
  link: link,
  cache: new InMemoryCache({
    addTypename: false
  }),
});
const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
});

Vue.use(VueApollo);
Vue.use(AuthPlugin);

const vm = new Vue({
  router,
  store,
  provide: apolloProvider.provide(),
  render: (h) => h(App),
});
vm.$mount('#app');
