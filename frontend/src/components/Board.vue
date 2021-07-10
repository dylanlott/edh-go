<template>
  <div class="board container is-fluid" v-if="user && game">
    <!-- TURN TRACKER -->
    <div class="column"><TurnTracker :game="game" /></div>
    <!-- END TURN TRACKER -->

    <!-- ### OPPONENTS BOARDSTATES ### -->
    <div class="columns" :key="player.ID" v-for="player in game.PlayerIDs">
      <div class="columns" v-if="user.ID && player.ID != user.ID">
        <div class="title">{{ player.Username }}</div>
      </div>
    </div>
    <!-- ### END OF OPPONENTS BOARDSTATES ### -->

    <div class="title">{{ user.Username }}</div>
    <!-- ### SELF BOARDSTATE - PUBLIC SECTION ### -->
    <div class="columns" v-if="self">
      <!-- SELF - BATTLEFIELD -->
      <div class="column">
        <p class="title is-5">Battlefield</p>
        <draggable
          class="columns is-mobile"
          @change="handleChange()"
          v-model="self.Field"
          group="people"
          @start="drag = true"
          @end="drag = false"
        >
          <!-- Note: Cards can only be tapped on the battlefield -->
          <div @click="handleTap(card)" class="column mtg-card" v-for="card in self.Field" :key="card.id">
            <Card v-bind="card"></Card>
          </div>
        </draggable>
      </div>
    </div>
    <!-- END SELF BATTLEFIELD -->

    <!-- SELF - PRIVATE SECTION -->
    <div class="columns">
      <div class="column">
        <div class="columns is-mobile is-multiline">
          <div class="column is-full">
            <!-- SELF - HAND-->
            <p class="title is-5">Hand</p>
            <draggable
              class="columns is-flex is-multiline is-mobile is-align-items-flex-start"
              v-model="self.Hand"
              group="people"
              @change="handleChange()"
              @start="drag = true"
              @end="drag = false"
            >
              <div class="column mtg-card" v-for="card in self.Hand" :key="card.id">
                <Card v-bind="card"></Card>
              </div>
            </draggable>
          </div>
        </div>
      </div>
    </div>

    <hr />

    <template>
      <b-navbar>
        <template #start>
          <draggable
            v-model="self.Library"
            group="people"
            @change="handleChange()"
            @start="drag = true"
            @end="drag = false"
          >
            <b-navbar-item @click="handleDraw()" href="#">
              <button class="button is-primary">
                <span class="icon">
                  <i class="fa fa-book"></i>
                </span>
                <span>Draw</span>
              </button>
            </b-navbar-item>
          </draggable>
          <b-navbar-item href="#">
            <button class="button is-primary">
              <span class="icon">
                <i class="fa fa-arrow-up"></i>
              </span>
              <span>Untap</span>
            </button>
          </b-navbar-item>
          <!-- <b-navbar-dropdown label="Info">
              <b-navbar-item href="#"> About </b-navbar-item>
              <b-navbar-item href="#"> Contact </b-navbar-item>
            </b-navbar-dropdown> -->
        </template>

        <template #end>
          <b-navbar-item tag="div">
            <div class="buttons">
              <!-- <a class="button is-primary"> <strong>Sign up</strong> </a>
              <a class="button is-light"> Log in </a> -->
            </div>
          </b-navbar-item>
        </template>
      </b-navbar>
    </template>
  </div>
</template>





    <!-- END OF SELF PRIVATE VIEW -->
    <!-- ### END OF SELF ### -->
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
    this.$store.dispatch('getBoardStates', this.$route.params.id);
    this.$store.dispatch('getGame', this.$route.params.id).then(() => {
      this.game.PlayerIDs.forEach((player) => {
        this.$store.dispatch('subToBoardstate', {
          gameID: this.$route.params.id,
          userID: player.ID,
        });
      });
      this.$store.dispatch('subscribeToGame', this.$route.params.id);
    });
  },
  computed: mapState({
    game: (state) => state.Game.game,
    bs: (state) => state.BoardStates.boardstates,
    self: (state) => state.BoardStates.self,
    user: (state) => state.User.User,
  }),
  methods: {
    handleDraw() {
      this.$store.dispatch('draw', this.self);
    },
    handleChange() {
      this.$store.dispatch('mutateBoardState', this.self);
    },
    handleTap(card) {
      // TODO: Make this a boardstate action
      card.Tapped = true;
      this.handleChange();
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
