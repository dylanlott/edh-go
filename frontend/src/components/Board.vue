<template>
  <div class="board shell">
    <h1 class="title shell">{{ gameID }}</h1>
    <TurnTracker gameID="gameID"/>
    <div class="opponents">
      <div :key="o.id" v-for="o in opponents" class="shell">
        <h1 class="title">{{ o.username }}</h1>
        <PlayerState v-bind="o.boardstate"></PlayerState>
      </div >
    </div>
    <hr>
    <div class="self shell">
      <h1 class="title">{{ self.username }}</h1>
      <SelfState></SelfState>
    </div>
    <div class="shell controlpanel columns">
      <div class="columns">
        <div class="column">
          <button class="button is-small is-primary">Collapse All</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Untap</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Draw</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Shuffle</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import PlayerState from '@/components/PlayerState.vue'
import SelfState from '@/components/SelfState.vue'
import Card from '@/components/Card.vue'
import TurnTracker from '@/components/TurnTracker.vue'
import draggable from 'vuedraggable'

const testCard = {
  id: '1',
  name: 'Karlov of the Ghost Council',
  convertedManaCost: '3',
  manaCost: "3 B W",
  colorIdentity: 'BU',
  power: '7',
  toughness: '8',
  text: 'When this card enters the battlefield, make Brenden mill 10 cards.',
  types: 'Legendary Creature Wizard',
  image: '',
  counters: {
    "Skithiryx": {
      type: "poison",
      count: 2
    },
    "Glory of Warfare": {
      type: "+1/+1",
      count: "1"
    }
  },
  labels: ["test label", "countered by Teysa"]
}

export default {
  name: 'board',
  data () {
    return {
      // TODO: This needs to be modeled after BoardState
      gameID: this.$route.params.id,
      self: {
        id: 4,
        username: "shakezula",
        boardstate: {
          library: [],
          graveyard: [],
          exiled: [],
          battlefield: [],
          hand: [],
          controlled: [],
        }
      },
      opponents: [{
        id: 1,
        username: "player1",
        boardstate: {
          library: [testCard],
          graveyard: [testCard],
          exiled: [testCard],
          battlefield: [testCard],
          controlled: [testCard],
        }
      }]
    }
  },
  components: {
    TurnTracker,
    PlayerState,
    SelfState,
    Card
  },

}
</script>
<style media="screen" scoped>
.shell {
  padding: .5rem;
  border: 1px solid #efefef;
  margin: .25rem 0.0rem;
}
</style>
