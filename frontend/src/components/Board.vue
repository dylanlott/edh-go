<template>
<div class="board">
  <h1 class="display-5">{{ gameID }}</h1>
  <div class="turn-track">
    <b>Active Turn: {{ turn.player }}</b>
    <div class="progress">
      <div class="progress-bar bg-success" role="progressbar" style="width: 25%" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100"></div>
    </div>
  </div>
  <div class="opponents">
    <div v-for="o in opponents" class="shell">
      <h1>{{ o.username }}</h1>
      <PlayerState v-bind="o.boardstate"></PlayerState>
    </div>
  </div>
  <hr>
  <div class="self shell">
    <h1>{{ self.username }}</h1>
    <SelfState></SelfState>
  </div>
  <div class="container controlpanel">
    <div class="row">
      <button class="btn btn-primary btn-sm">Toggle boards</button>
    </div>
  </div>
</div>
</template>
<script>
import PlayerState from '@/components/PlayerState.vue'
import SelfState from '@/components/SelfState.vue'
import Card from '@/components/Card.vue'
import draggable from 'vuedraggable'

const testCard = {
  id: '1',
  name: 'Karlov of the Ghost Council',
  convertedManaCost: '3',
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
      gameID: this.$route.params.id,
      turn: {
        phase: '',
        player: 'player1'
      },
      phases: [
        'untap',
        'upkeep',
        'draw',
        'main phase 1',
        'combat',
        'declare attackers',
        'declare blockers',
        'resolve combat damage',
        'main phase 2',
        'end phase',
        'discard'
      ],
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
    PlayerState,
    SelfState,
    Card
  }
}
</script>
<style media="screen" scoped>
.shell {
  padding: .5rem;
  border: 1px solid #efefef;
  margin: .25rem 0.0rem;
}
</style>
