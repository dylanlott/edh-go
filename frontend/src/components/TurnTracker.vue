<template>
  <div class="turn-tracker">
    <div class="columns">
      <section class="column is-11 is-mobile">
        <p class="has-text-primary">
          {{ game.game.Turn.Player }} |
          {{ game.game.Turn.Phase }} |
          {{ game.game.Turn.Number }}</p>
        <b-progress :value="progress" size="is-small" show-value></b-progress>
      </section>
      <section class="column is-1">
        <b-button
          type="button"
          @click="handleTick(game.game)"
          class="is-success">
          Next
        </b-button>
      </section>
    </div>
  </div>
</template>
<script>
export default {
  name: 'TurnTracker',
  data () {
    return {
      currentPos: 0,
      currentPhase: "",
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
        'end step',
      ],
      phasesMap: {
        0: 'untap',
        1: 'upkeep',
        2: 'draw',
        3: 'main phase 1',
        4: 'combat',
        5: 'declare attackers',
        6: 'declare blockers',
        7: 'first strike / double strike',
        8: 'resolve combat damage',
        9: 'main phase 2',
        10: 'end step',
      }
    }
  },
  // we take in Game so we don't ever have to fetch it.
  // we reduce calls to server this way.
  props: [ 'game' ],
  computed: {
    progress () {
      return 100
    },
  },
  methods: {
    handleTick (game) {
      const g = this.tick(game)
      this.$store.dispatch('updateGame', g)
    },
    // tick takes a `game` object and checks for existence of a Turn. 
    // If no Turn exists it will put it into the default `setup` phase.
    // Otherwise it tries to handle the turn as normal.
    tick (game) {
      // create a new object so that we're operating on our own data
      const g = Object.assign({}, game)
      // determine if game turn has been set. 
      // if not, set it to the default setup phase.
      if (!g.Turn) {
        console.log("no existing game.Turn detected, setting to defaults")
        g.Turn = {
          Number: "0",
          Phase: "setup",
          Player: g.PlayerIDs ? g.PlayerIDs[0].Username : ""
        }
        return g
      }
      
      if (!g.Turn.Number) {
        g.Turn.Number = 0
      }
      
      // setup is the default phase before the game starts, where chat,
      // rolls for turn, and deck tweaking can occur.
      // TODO: Make it so that only the game creator can leave the setup phase
      // TODO: Allow players to confirm "ready" if (phase === "setup")
      if (g.Turn.Phase === "setup") {
        g.Turn.Phase = this.phases[0]
        return g
      }

      // detect if we're at the end of a turn cycle
      if (g.Turn.Phase === this.phases[10]) {
        // set the phase back to the beginning of the turn cycle
        g.Turn.Phase = this.phases[0]
        // tick the turn number up
        g.Turn.Number = g.Turn.Number++
        return g
      }

      // if we're not at the end of a cycle, find where we are
      var pos = 0
      let current = this.phases.find((phase, i) => {
        pos = i
        if (phase === g.Turn.Phase) { return true }
      })
      if (current === undefined) {
        // NB: can't detect current phase, so default to setup
        // and return the game object
        g.Turn.Phase = "setup"
        return g
      }
      
      // happy path 
      this.currentPos = pos++
      g.Turn.Phase = this.phases[pos++]
      return g
    },
  }
}
</script>
<style media="screen">
  .progress {
    margin: 0.5rem 0rem;
  }
  .turn-tracker {
    margin: 0.5rem;
  }
</style>
