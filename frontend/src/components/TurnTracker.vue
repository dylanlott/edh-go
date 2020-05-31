<template>
  <div class="turn-tracker shell">
    <div class="columns">
      <section class="column is-12 is-mobile">
        <p class="has-text-primary">
          {{ players[turn.player]['username'] }}
          Phase: {{ phases[turn.phase] }}
        </p>
        <b-progress :value="progress" size="is-medium" show-value>
        </b-progress>
      </section>
      <section class="column is-2">
        <b-button
          type="button"
          @click="tick()"
          class="is-success">
          Next Phase
        </b-button>
      </section>
    </div>
  </div>
</template>
<script>
export default {
  name: 'turntracker',
  data () {
    return {
      turn: {
        // starting indexes for game.
        phase: 0,
        player: 0
      },
      players: [
        // TODO: Make these props.
        {
          id: '1',
          username: 'shakezula'
        },
        {
          id: '2',
          username: 'player2'
        }
      ],
      phases: [
        'untap',
        'upkeep',
        'draw',
        'main phase 1',
        'combat',
        'declare attackers',
        'declare blockers',
        'first strike / double strike',
        'resolve combat damage',
        'main phase 2',
        'end phase',
        'discard'
      ],
    }
  },
  computed: {
    progress () {
      const total = this.phases.length
      const current = this.turn.phase
      const v = ((current / total) * 100)
      return v 
    }
  },
  methods: {
    tick () {
      if ((this.turn.phase + 1) >= this.phases.length) {
        // turn ends, tick player over
        this.turn.phase = 0

        // if player is last in array, tick back to beginning rotation
        if ((this.turn.player + 1) >= this.players.length) {
          this.turn.player = 0
          return
        }

        this.turn.player++
        return
      }

      this.turn.phase++
      return
    },
  }
}
</script>
<style media="screen">
  .progress {
    margin: 0.5rem 0rem;
  }
  .turn-tracker {
    margin: 0.25rem;
  }
</style>
