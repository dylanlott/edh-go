<template>
  <div class="board shell">
    <pre>
      <!-- {{ game.Turn }} -->
      <!-- {{ user }}
      {{ boardstates }} -->
    </pre>
    <!-- <pre :key="player.ID" v-for="player in game.PlayerIDs">
      {{ player }}
    </pre> -->
    <!-- <h1 class="title shell">{{ gameID }}</h1> -->

  <!-- LIFE TRACKER -->
  <!-- TODO -->

  <!-- TURN TRACKER -->
  <div class="columns">
    <div class="shell column is-9">
      <TurnTracker :game="game" />
    </div>
    <!-- <div class="shell column is-3">
      <div class="title is-4">{{ boardstates.self.Life }}</div>
      <button class="button is-small" @click="increaseLife()">Increase</button>
      <button class="button is-small" @click="decreaseLife()">Decrease</button>
    </div> -->
  </div>

    <!-- OPPONENTS -->
    <!-- <div class="opponents">
      <div :key="opponent.ID" v-for="opponent in boardstates">
        <div v-if="opponent.Username != self.User.Username"></div>
      </div>
      <div :key="player.ID" v-for="player in game.PlayerIDs">
        <div v-if="player.Username !== self.User.Username">
          <h1 class="title">{{ player.Username }}</h1>
        </div>

        <div :key="p.ID" v-for="p in game.PlayerIDs">
          username: {{ p.Username }} {{ self.User.Username}}
          <div v-if="p.Username !== self.User.Username">
            {{ p }}
          </div>
          <div v-else>
            You are {{ p }}
          </div>
        </div>
      </div>
    </div> -->

    <!-- <hr /> -->

    <div class="self shell">
      <h1 class="title">
        {{ boardstates.self.User.username }}
        <!-- <p class="subtitle">{{ 
          boardstates.self.Commander ? boardstates.self.Commander[0].Name : ""
        }}</p> -->
      </h1>

      <!-- {{ boardstates.self }} -->

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
          <div class="column library is-one-quarter" @click="$store.dispatch('draw', boardstates.self)">
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
    .then((resp) => this.$store.dispatch('subscribeToBoardState', {
      userID: this.user.User.ID,
      gameID: this.$route.params.id,
    }))
    this.$store.dispatch('getGame', this.$route.params.id)
    .then((resp) => this.$store.dispatch('subscribeToGame', this.$route.params.id))
  },
  methods: {
    handleActivity(val) {
      console.log('logging activity: ', val)
      return
    },
    mutateBoardState() {
      console.log('mutate board state hit')
    },
  },
  watch: {
    // self: {
    //   handler (newVal, oldVal) {
    //     // we don't want to mutate state with this, 
    //     // or else we'll get infinite loops.
    //     // This is only where we should emit ActivityLog events.
    //     // this.handleActivity(newVal)

    //     // TODO: Call mutate board state action here.
    //   },
    //   deep: true
    // }
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
