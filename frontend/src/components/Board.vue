<template>
  <div class="board shell" v-if="user && game">
    <!-- TURN TRACKER -->
    <div class="shell column is-12">
      <TurnTracker :game="game" />
    </div>
    <div class="columns shell" :key="player.ID" v-for="player in game.PlayerIDs">
      <!-- ### OPPONENTS BOARDSTATES ### -->
      <div class="column" v-if="user.ID && player.ID != user.ID">
        <div class="title is-6">{{ player.Username }} {{ player.ID }}</div>
      </div>
      <!-- ### END OF OPPONENTS BOARDSTATES ### -->

      <!-- ### SELF BOARDSTATE ### -->
      <!-- LIFE  -->
      <div class="column" v-if="user.ID ? player.ID == user.ID : false">
        <div class="title is-6">{{ user.Username }} {{ user.ID }}</div>
      </div>

      <!-- ### END OF SELF ### -->
    </div>
  </div>
</template>
<script>
import _ from 'lodash';
import draggable from 'vuedraggable';
import Card from '@/components/Card';
import TurnTracker from '@/components/TurnTracker.vue';
import { mapState } from 'vuex';

export default {
  name: 'board',
  data() {
    return {};
  },
  created() {
    // Get the root game and then register boardstate listeners off of that.
    this.$store.dispatch('getGame', this.$route.params.id)
      .then(() => {
        // sub to game updates
        this.$store.dispatch('subscribeToGame', this.$route.params.id);
        // sub to player boardstate updates
        this.game.PlayerIDs.forEach((player) => {
          this.$store.dispatch('subToBoardstate', {
            gameID: this.$route.params.id,
            userID: player.ID,
          });
        });
      });

    // get initial boardstates
    this.$store.dispatch('getBoardStates', this.$route.params.id);
  },
  computed: mapState({
    game: (state) => state.Game.game,
    bs: (state) => state.BoardStates.boardstates,
    user: (state) => state.User.User,
  }),
  methods: {
    handleIncLife() {},
    handleDecLife() {},
    handleDraw() {},
  },
  components: {
    draggable,
    Card,
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
