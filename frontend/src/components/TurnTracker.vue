<template>
  <div class="turn-tracker">
    <div class="columns">
      <section class="column is-11 is-mobile">
        <p class="has-text-primary">{{ game.turn.Player }} - {{ game.turn.Phase }}</p>
        <b-progress :value="progress" size="is-small" show-value></b-progress>
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
      game: {
        ID: "",
        turn: {
          Player: "",
          Phase: "",
          Number: 0,
        },
        PlayerIDs: []
      },
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
    }
  },
  computed: {
    progress () {
      return 100
    }
  },
  apollo: {
    gameSubscription() {
      return {
        query: gql`
          query($gameID: String!) {
            games(gameID: $gameID) {
              id 
              turn {
                Player
                Phase
                Number
              }
              playerIDs {
                username
                id
              }
            }
          } 
        `,
        variables: {
          gameID: this.$route.params.id
        },
        update(data) {
          this.game = data.games[0]
        },
        subscribeToMore: {
          document: gql`subscription($game: InputGame!) {
            gameUpdated(game: $game) {
              id 
              turn {
                Player
                Phase
                Number
              }
            }
          }`,
          variables: {
            game: {
              ID: this.$route.params.id,
              Turn: {
                Player: this.game.turn.Player || "",
                Phase: this.game.turn.Phase || "",
                Number: this.game.turn.Number || 0
              }
            }
          },
          updateQuery: (previousResult, { subscriptionData }) => {
            console.log('turntracker#previousResult: ', previousResult)
            console.log('turntracker#subscriptionData: ', subscriptionData)
          }
        } 
      }
    }
  },
  methods: {
    mutateGame () {
      this.$apollo.mutate({
        mutation: gql`mutation ($input: InputGame!) {
          updateGame(input: $input) {
            id
            turn {
              Player
              Phase
              Number
            }
            playerIDs {
              id
              username
            }
          }
        }
        `,
        variables: {
          input: {
            ID: this.$route.params.id,
            Turn: {
              Player: this.game.turn.Player,
              Phase: this.game.turn.Phase,
              Number: this.game.turn.Number,
            },
            PlayerIDs: this.getPlayerIDs(),
          }
        },
        update: (store, { data }) => {
          return data
        }
      }).then((data) => {
        return data
      }).catch((err) => {
        console.error('error mutating game: ', err)
        return err
      })
    },
    // this is a utility function to properly format received playerIDs into input
    getPlayerIDs () {
      const formatted = this.game.playerIDs.map((u, i) => {
        return {
          Username: u.username,
          ID: u.id,
        }
      })
      return formatted
    },
    tick () {
      // setup is the default phase before the game starts, where chat,
      // rolls for turn, and deck tweaking can occur.
      // TODO: Make it so that only the game creator can leave the setup phase
      // TODO: Allow players to confirm "ready" if (phase === "setup")
      if (this.game.turn.Phase === "setup") {
        console.log('setup phase ending, setting to untap.')
        this.game.turn.Phase = this.phases[0]
        return
      }
      if (this.game.turn.Phase === this.phases[10]) {
        this.game.turn.Phase = this.phases[0]
        return
      }
      var pos = 0
      let current = this.phases.find((phase, i) => {
        pos = i
        if (phase === this.game.turn.Phase) { return true }
      })
      if (current === undefined) {
        // NB: can't detect current phase, so default to setup
        this.game.turn.Phase = "setup"
        return
      }
      // happy path 
      this.currentPos = pos++
      this.game.turn.Phase = this.phases[pos++]
      console.log('set phase: ', this.game.turn.Phase)
      this.mutateGame()
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
