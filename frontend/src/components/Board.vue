<template>
  <div class="board shell" v-if="user && game">
    <!-- TURN TRACKER -->
    <div class="shell column">
      <TurnTracker :game="game" />
    </div>
    <div class="columns shell" :key="player.ID" v-for="player in game.PlayerIDs">
      <!-- ### OPPONENTS BOARDSTATES ### -->
      <div class="column" v-if="user.ID && player.ID != user.ID">
        <div class="title">{{ player.Username }} {{ player.ID }}</div>
      </div>
    </div>
    <div class="columns shell">
      <!-- ### END OF OPPONENTS BOARDSTATES ### -->

      <!-- ### SELF BOARDSTATE ### -->
      <div class="column" v-if="self">
        <div class="title is-6">{{ user.Username }} {{ user.ID }}</div>

        <!-- SELF - BATTLEFIELD -->
        <div class="column">
          <p class="title is-5">Battlefield</p>
          <draggable class="card-wrapper" v-model="self.Field" group="people" @start="drag = true" @end="drag = false">
            <div v-for="card in self.Field" :key="card.id" class="columns">
              <Card v-bind="card" />
            </div>
          </draggable>
        </div>
        <!-- END SELF BATTLEFIELD -->

        <!-- SELF - LIBRARY  -->
        <div class="column" @click="handleDraw()">
          <p class="title is-5">Library</p>
          <draggable
            class="column card-wrapper"
            v-model="self.Library"
            group="people"
            @start="drag = true"
            @end="drag = false"
          >
            <div v-for="card in self.Library" :key="card.id">
              <Card v-bind="card" hidden="true" />
            </div>
          </draggable>
        </div>
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
    this.$store.dispatch('getGame', this.$route.params.id).then(() => {
      this.game.PlayerIDs.forEach((player) => {
        this.$store.dispatch('subToBoardstate', {
          gameID: this.$route.params.id,
          userID: player.ID,
        });
      });
      this.$store.dispatch('subscribeToGame', this.$route.params.id);
      this.$store.dispatch('getBoardStates', this.$route.params.id);
    });

    // get initial boardstates
    console.log('getting game ', this.$route.params.id)
  },
  computed: mapState({
    game: (state) => state.Game.game,
    bs: (state) => state.BoardStates.boardstates,
    self: (state) => state.BoardStates.self,
    user: (state) => state.User.User,
  }),
  methods: {
    handleIncLife() {},
    handleDecLife() {},
    handleDraw() {
      this.$store.dispatch('draw', this.self)
    },
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
