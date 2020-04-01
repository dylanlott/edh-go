<template>
  <div class="game">
    <main>
      <section>
          <header>
            <h1>Game state</h1>
          </header>
      </section>
      <section>
        <aside class="fullwidth">
          Your Board
          <aside class="player">
            <Player :id="player.id" :boardstate="player.boardstate"/>
          </aside>
        </aside>
      </section>

      <section>
        <aside class="fullwidth boardstate">
          Board State
            <aside
            class="opponent"
            v-for="opponent in opponents"
            :key="opponent.id">
              <Opponent :id="opponent.id" :boardstate="opponent.boardstate"/>
              <br>
            </aside>
        </aside>
      </section>
    </main>
  </div>
</template>

<script>
// @ is an alias to /src
import Player from '@/components/Player.vue'
import Opponent from '@/components/Opponent.vue'

export default {
  name: 'Game',
  data () {
    return {
      player: {
        id: 'ax8ga8w',
        boardstate: {
          graveyard: [],
          library: [],
          exiled: [],
          battlefield: [],
          meta: {}
        }
      },
      opponents: [
        {
          boardstate: {},
          id: '1'
        },
        {
          boardstate: {},
          id: '2'
        },
        {
          boardstate: {},
          id: '3'
        },
        {
          boardstate: {},
          id: '4'
        }
      ]
    }
  },
  components: {
    Player,
    Opponent
  }
}
</script>
<style media="screen">
.boardstate {
  width: auto;
}
.player__list {

}
.player {
  width: auto;
}
.fullwidth {
  width: 100%;
}

/* flex grid for player boards */
.flex-grid {
  display: flex;
}
.col {
  flex: 1;
}

@media (max-width: 400px) {
  .flex-grid {
    display: block;
  }
}
</style>
