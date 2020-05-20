import Vue from 'vue';
import Router from 'vue-router';

import Landing from '@/components/Landing.vue';
import Login from '@/components/Login.vue';
import Games from '@/components/Games.vue';
import Board from '@/components/Board.vue';
import Card from '@/components/Card.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
        component: Landing,
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
    },
    {
      path: '/games/:id',
      name: 'board',
      component: Board,
    },
    {
      path: '/card/:id',
      name: 'card',
      component: Card,
    }
  ],
});
