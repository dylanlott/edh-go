<template>
  <div class="turn-tracker">
    <div class="columns">
      <section class="column is-11 is-mobile">
        <p class="has-text-primary">
          <!-- <b> {{ players[turn.player]['username'] }} - {{ phases[turn.phase] }} </b> -->
          <b>{{ players }}</b>
          <b>{{ turn }}</b>
        </p>
        <b-progress :value="progress" size="is-medium" show-value></b-progress>
      </section>
      <section class="column is-1">
        <b-button
          type="button"
          @click="tick()"
          class="is-success">
          Next 
        </b-button>
      </section>
    </div>
  </div>
</template>
<script>
import gql from 'graphql-tag';

export default {
  name: 'turntracker',
  data () {
    return {
      turn: {
        // starting indexes for game.
        phase: 0,
        player: 0
      },
      game: {},
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
  apollo: {
    game: {
      query: gql`
          subscription($game: InputGame!) {
            gameUpdated(game: $game) {
              ID
              Turn {
                Player
                Phase
                Number
              }
              PlayerIDs {
                Username
                ID
              }
            }
          } 
        `,
      update(data) {
        console.log('game#update: ', data)
      },
      subscribeToMore: {
        document: gql`
          subscription($game: InputGame!) {
            gameUpdated(game: $game) {
              ID
              Turn {
                Player
                Phase
                Number
              }
              PlayerIDs {
                Username
                ID
              }
            }
          } 
        `,
        updateQuery: (previousResult, { subscriptionData}) => {
          console.log("game#previousResult: ", previousResult)
          console.log("game#subscriptionData: ", subscriptionData)
        }
      }
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
    margin: 0.5rem;
  }
</style>
