<template>
  <div class="board shell">
    <!-- <pre>{{ user }}</pre> -->
    <!-- v-if="boardstate.User.ID != user.User.ID" -->
    <div class="columns">
      <pre>
      </pre>
      <!-- <b-collapse
        class="card"
        animation="slide"
        v-for="(boardstate, index) of opponents"
        :key="index"
        :open="opponentStateOpen == index"
        @open="opponentStateOpen = index"
      >
        <template #trigger="props">
          <div class="card-header" role="button">
            <p class="card-header-title">
              {{ boardstate.Username }}
            </p>
            <a class="card-header-icon">
              <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"> </b-icon>
            </a>
          </div>
        </template>
        <div class="card-content">
          <div class="content">
            <Opponent :boardstate="boardstate" />
          </div>
        </div>
      </b-collapse> -->
    </div>

    <!-- TURN TRACKER -->
    <div class="columns">
      <div class="shell column is-9">
        <TurnTracker :game="game" />
      </div>
      <div class="shell column is-3">
        <!-- <div class="title is-4">{{ boardstates[user.ID].Life }}</div> -->
        <button class="button is-small" @click="increaseLife()">Increase</button>
        <button class="button is-small" @click="decreaseLife()">Decrease</button>
      </div>
    </div>

    <SelfState :boardstate="self"></SelfState>

    <!-- CONTROL PANEL -->
    <div class="shell controlpanel columns">
      <div class="columns">
        <!-- <div class="column">
          <button class="button is-small is-primary">Collapse All</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Untap</button>
        </div> -->
        <div class="column">
          <button @click="draw()" class="button is-small is-primary">Draw</button>
        </div>
        <!-- <div class="column">
          <button class="button is-small is-primary">Shuffle</button>
        </div> -->
        <div class="column">
          <button @click="mill()" class="button is-small is-primary">Mill</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import _ from 'lodash';
import draggable from 'vuedraggable';
import Card from '@/components/Card';
import PlayerState from '@/components/PlayerState.vue';
import SelfState from '@/components/SelfState.vue';
import Opponent from '@/components/Opponent.vue';
import TurnTracker from '@/components/TurnTracker.vue';
import { mapState } from 'vuex';

export default {
  name: 'board',
  data() {
    return {
      opponentStateOpen: false,
    };
  },
  created() {
    this.$store.dispatch('getBoardStates', this.$route.params.id).then(() =>
      this.$store.dispatch('subscribeToBoardState', {
        userID: this.user.ID,
        gameID: this.$route.params.id,
      })
    )
   
    // get the game info 
    this.$store
      .dispatch('getGame', this.$route.params.id)
      .then(() => this.$store.dispatch('subscribeToGame', this.$route.params.id));
     
    // sub to all boardstate updates, self included
    this.$store.dispatch('subAll', this.$route.params.id);
  },
  computed: mapState({
    game: (state) => state.Game,
    boardstates: (state) => state.BoardStates,
    user: (state) => state.User.User,
    self: (state) => state.BoardStates[state.User.User.ID]
  }),
  methods: {
    mutateBoardState() {
      return this.$store.dispatch('mutateBoardState', this.boardstates[this.user.ID]);
    }, 
    print (...args) {
      console.log(...args)
    }
  },
  components: {
    draggable,
    Card,
    PlayerState,
    SelfState,
    Opponent,
    TurnTracker,
  },
};
</script>
<style media="screen" scoped>
.shell {
  padding: 0.5rem;
  border: 1px solid #efefef;
  margin: 0.25rem 0rem;
}

.bordered {
  border: 1px #000;
}
</style>
