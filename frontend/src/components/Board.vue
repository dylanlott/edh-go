<template>
  <div class="board shell">
      <div class="columns shell" :key="player.ID" v-for="player in game.PlayerIDs">
        <!-- FIND OPPONENTS  -->
        <div class="column" v-if="player.ID != user.ID">
          // {{ player.ID }}
          // BOARDSTATE: {{ boardstates[player.ID] ? boardstates[player.ID].Commander[0].Name : "" }}
        </div>
        <!-- FIND SELF  -->
        <div class="column" v-if="player.ID == user.ID">
          // SELF: {{ player.ID }}
          // BOARDSTATE: {{ boardstates[player.ID] ? boardstates[player.ID].Commander[0].Name : "" }}
        </div>
      </div>

      <!-- TURN TRACKER -->
      <div class="shell column is-9">
        <TurnTracker :game="game" />
      </div>

      </div>
    <!-- <div class="columns">
      <div class="shell column is-3">
        <div class="title is-4">{{ boardstates[user.ID].Life }}</div>
        <button class="button is-small" @click="increaseLife()">Increase</button>
        <button class="button is-small" @click="decreaseLife()">Decrease</button>
      </div>
    </div> -->

    <!-- <SelfState :boardstate="self"></SelfState> -->

    <!-- CONTROL PANEL -->
    <!-- <div class="shell controlpanel columns">
      <div class="columns">
        <div class="column">
          <button class="button is-small is-primary">Collapse All</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Untap</button>
        </div>
        <div class="column">
          <button @click="draw()" class="button is-small is-primary">Draw</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Shuffle</button>
        </div>
        <div class="column">
          <button @click="mill()" class="button is-small is-primary">Mill</button>
        </div>
      </div>
    </div> -->
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
    // Start from the game:
    // Get the root game and then register boardstate listeners off of that.
    this.$store
      .dispatch('getGame', this.$route.params.id)
      .then(() => this.$store.dispatch('subscribeToGame', this.$route.params.id));

    // this.$store.dispatch('subscribeToGame', this.$route.params.id)
    this.$store
      .dispatch('getBoardStates', this.$route.params.id)
      .then(() => this.$store.dispatch('subAll', this.$route.params.id))
  },
  computed: mapState({
    game: (state) => state.Game.game,
    boardstates: (state) => state.BoardStates.boardstates,
    user: (state) => state.User.User,
  }),
  methods: {
    mutateBoardState() {
      // QUESTION: Does this work like I think it does? I think it does, but need to make sure.
      return this.$store.dispatch('mutateBoardState', this.boardstates[this.user.ID]);
    },
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
