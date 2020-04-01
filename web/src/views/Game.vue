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
        </aside>
      </section>

      <section>
        <aside class="fullwidth boardstate">
          Board State
            <aside
            class="player"
            v-for="player in players"
            :key="player.id">
              <Player :id="player.id" :boardstate="player.boardstate"/>
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

export default {
  name: 'Game',
  data () {
    return {
      players: [
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
    Player
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
