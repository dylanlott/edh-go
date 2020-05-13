<template>
  <div class="turn-tracker">
    <div class="row">
      <div class="col-10">
        <b>{{ players[turn.player]['username'] }} Phase: {{ phases[turn.phase] }}</b>
      </div>
      <div class="col-2">
        <button
          type="button"
          @click="tick()"
          class="btn btn-success">
          Next Phase
        </button>
      </div>
    </div>
    <div class="progress">
      <div class="progress-bar bg-success"
      role="progressbar"
      :style="progress"
      :aria-valuenow="progress"
      aria-valuemin="0"
      aria-valuemax="100"></div>
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
      return `width: ${v}%`
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
</style>
