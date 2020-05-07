<template>
<div class="board">
  <h1>Boardstate</h1>

  <div class="opponents">
    <div v-for="o in opponents" class="shell">
      <h1>{{ o.username }}</h1>
      <PlayerState
      ></PlayerState>
    </div>
  </div>
  <hr>
  <div class="self">
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
          hand: [],
          controlled: [],
        }
      }, {
        id: 2,
        username: "player2",
        boardstate: {
          library: [],
          graveyard: [],
          exiled: [],
          battlefield: [],
          hand: [],
          controlled: [],
        }
      }, {
        id: 3,
        username: "player3",
        boardstate: {
          library: [],
          graveyard: [],
          exiled: [],
          battlefield: [],
          hand: [],
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
<style media="screen">
.shell {
  padding: .5rem;
  border: 1px solid #efefef;
  margin: .25rem 0.0rem;
}
</style>
