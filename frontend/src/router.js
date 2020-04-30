import Vue from 'vue';
import Router from 'vue-router';

import Home from '@/components/Home.vue';
import Login from '@/components/Login.vue';
import Games from '@/components/Games.vue'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/games',
      name: 'games',
      component: Games
    }
  ],
});
