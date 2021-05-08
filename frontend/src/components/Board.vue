<template>
  <div class="board shell">
  <!-- <pre :key="player.ID" v-for="player in game.PlayerIDs"> -->
  <h1 class="title is-6">{{ game.game.ID }}</h1>

  <!-- TURN TRACKER -->
  <div class="columns">
    <div class="shell column is-9">
      <TurnTracker :game="game" />
    </div>
    <div class="shell column is-3">
      <div class="title is-4">{{ boardstates.self.Life }}</div>
      <button class="button is-small" @click="increaseLife()">Increase</button>
      <button class="button is-small" @click="decreaseLife()">Decrease</button>
    </div>
  </div>

    <!-- OPPONENTS -->
    <div class="opponents">
      <div :key="opponent.ID" v-for="opponent in boardstates.boardstates">
        <div v-if="opponent.Username != user.Username">
          <pre>{{ opponent }}</pre> 
        </div>
      </div>
    </div>

    <div class="self shell">
      <h1 class="title">
        {{ user.User.Username }}
        <p class="subtitle">{{ boardstates.self.Commander[0] ? boardstates.self.Commander[0].Name : "" }}</p>
      </h1>

      <div>
        <div class="columns">
          <div class="column">
            <p class="title is-5">Battlefield</p>
            <draggable
              class="card-wrapper bordered battlefield"
              group="board" 
              v-model="boardstates.self.Field"
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div 
              @click="tap(card)"
              v-for="(card, i) in boardstates.self.Field" 
              :key="i" 
              >
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
        </div>
        <div class="columns">
        </div>

        <div class="columns">
          <div class="column hand is-three-quarters">
            <p class="title is-4">Hand</p>
            <draggable
              class="columns card-wrapper"
              v-model="boardstates.self.Hand"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div class="column mtg-card" v-for="(card, i) in boardstates.self.Hand" :key="i">
                <Card v-bind="card"></Card>
              </div>
            </draggable>
          </div>
          <div class="column library is-one-quarter" @click="draw()">
            <p class="title is-5">Library</p>
            <draggable
              class="column card-wrapper"
              v-model="boardstates.self.Library"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in boardstates.self.Library" :key="card.id">
                <Card v-bind="card" hidden="true"/>
              </div>
            </draggable>
          </div>
        </div>
        <hr />
      </div>
    </div>

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
import Opponents from '@/components/Opponents.vue'
import TurnTracker from '@/components/TurnTracker.vue';
import { fetch } from '@/cards.js'
import { mapState } from 'vuex'

export default {
  name: 'board',
  created () {
    this.$store.dispatch('getBoardStates', this.$route.params.id)
    .then(() => this.$store.dispatch('subscribeToBoardState', {
      userID: this.user.User.ID,
      gameID: this.$route.params.id,
    }))
    this.$store.dispatch('getGame', this.$route.params.id)
    .then(() => this.$store.dispatch('subscribeToGame', this.$route.params.id))
  },
  methods: {
    mutateBoardState() {
      return this.$store.dispatch('mutateBoardState', this.boardstates.self)
    },
    mill() {
      const bs = Object.assign({}, this.boardstates.self)
      const card = bs.Library.shift()
      bs.Graveyard.push(card)
      this.$store.dispatch('mutateBoardState', bs)
    },
    draw(num) {
      if (num || num > 1) {
        console.log('you can only draw one card at a time')
      }
      // NB: Not sure if I should handle this here or as an action. 
      // Both ways have their pros and cons, and I've implemented it 
      // as both.
      const bs = Object.assign({}, this.boardstates.self)
      if (bs.Library.length < 1) {
        console.error('you lose the game - cannot draw from an empty library')
      }
      const card = bs.Library.shift()
      bs.Hand.push(card)
      this.$store.dispatch("mutateBoardState", bs)
    },
    tap(card) {
      card.Tapped = !card.Tapped
      this.mutateBoardState()
    },
    increaseLife() {
      const bs = Object.assign({}, this.boardstates.self)
      bs.Life = bs.Life + 1
      this.$store.dispatch("mutateBoardState", bs)
    },
    decreaseLife() {
      const bs = Object.assign({}, this.boardstates.self)
      bs.Life = bs.Life - 1
      this.$store.dispatch("mutateBoardState", bs)
    }
  },
  computed: mapState({
    game: state => state.Game,
    boardstates: state => state.BoardStates,
    user: state => state.User,
  }),
  components: {
    draggable,
    Card,
    PlayerState,
    SelfState,
    Opponents,
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
