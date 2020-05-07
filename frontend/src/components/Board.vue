<template>
<div class="board">
  <h1 class="display-5">Boardstate</h1>
  <div class="turn-track">
    <b>Active Turn: {{ turn.player }}</b>
    <div class="progress">
      <div class="progress-bar bg-success" role="progressbar" style="width: 25%" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100"></div>
    </div>
  </div>
  <div class="opponents">
    <div v-for="o in opponents" class="shell">
      <h1>{{ o.username }}</h1>
      <PlayerState ></PlayerState>
    </div>
  </div>
  <hr>
  <div class="self shell">
    <h1>Self</h1>
    <SelfState></SelfState>
  </div>
</div>
</template>
<script>
import PlayerState from '@/components/PlayerState.vue'
import SelfState from '@/components/SelfState.vue'

export default {
  name: 'board',
  data () {
    return {
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
          library: [],
          graveyard: [],
          exiled: [],
          battlefield: [],
          controlled: [],
        }
      }]
    }
  },
  components: {
    PlayerState,
    SelfState
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
