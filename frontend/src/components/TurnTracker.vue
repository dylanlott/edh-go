<template>
  <div class="turn-tracker">
    <div class="columns">
      <section class="column is-11 is-mobile">
        <p class="has-text-primary">{{ Game.Turn.Player }} - {{ Game.Turn.Phase }}</p>
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
      Game: {
        ID: "",
        Turn: {
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
        variables: {
          gameID: this.$route.params.id
        },
        // update(data) {
        //   this.Game = data.games[0]
        //   console.log('#TurnTracker#update set this.Game: ', this.Game)
        // },
      //   subscribeToMore: {
      //     document: gql`subscription($game: InputGame!) {
      //       gameUpdated(game: $game) {
      //         ID 
      //         Turn {
      //           Player
      //           Phase
      //           Number
      //         }
      //       }
      //     }`,
      //     variables: {
      //       game: {
      //         ID: this.$route.params.id,
      //         Turn: {
      //           Player: this.Game.Turn.Player || "",
      //           Phase: this.Game.Turn.Phase || "",
      //           Number: this.Game.Turn.Number || 0
      //         },
      //         PlayerIDs: this.Game.PlayerIDs ? this.Game.PlayerIDs : []
      //       }
      //     },
      //     updateQuery: (previousResult, { subscriptionData }) => {
      //       console.log('turntracker#previousResult: ', previousResult)
      //       console.log('turntracker#subscriptionData: ', subscriptionData)
      //     }
      //   } 
      }
    }
  },
  methods: {
    mutateGame () {
      console.log('mutating game with turntracker: ', this.Game)
      this.$apollo.mutate({
        mutation: gql`mutation ($input: InputGame!) {
          updateGame(input: $input) {
            ID 
            Turn {
              Player
              Phase
              Number
            }
            PlayerIDs {
              ID 
              Username
            }
          }
        }
        `,
        variables: {
          input: {
            ID: this.$route.params.id,
            Turn: {
              Player: this.Game.Turn.Player,
              Phase: this.Game.Turn.Phase,
              Number: this.Game.Turn.Number,
            },
            PlayerIDs: this.getPlayerIDs(),
          }
        },
        results (data) {
          console.log('results? ', data)
        }
      }).then((data) => {
        console.log('TURN TRACKER DATA: ', data)
        return data
      }).catch((err) => {
        console.error('TurnTracker: Error mutating game: ', err)
        return err
      })
    },
    // this is a utility function to properly format received playerIDs into input
    getPlayerIDs () {
      const formatted = this.Game.PlayerIDs.map((u, i) => {
        return {
          Username: u.Username,
          ID: u.ID,
        }
      })
      return formatted
    },
    tick () {
      // setup is the default phase before the game starts, where chat,
      // rolls for turn, and deck tweaking can occur.
      // TODO: Make it so that only the game creator can leave the setup phase
      // TODO: Allow players to confirm "ready" if (phase === "setup")
      if (this.Game.Turn.Phase === "setup") {
        console.log('setup phase ending, setting to untap.')
        this.Game.Turn.Phase = this.phases[0]
        return
      }
      if (this.Game.Turn.Phase === this.phases[10]) {
        this.Game.Turn.Phase = this.phases[0]
        return
      }
      var pos = 0
      let current = this.phases.find((phase, i) => {
        pos = i
        if (phase === this.Game.Turn.Phase) { return true }
      })
      if (current === undefined) {
        // NB: can't detect current phase, so default to setup
        this.Game.Turn.Phase = "setup"
        return
      }
      // happy path 
      this.currentPos = pos++
      this.Game.Turn.Phase = this.phases[pos++]
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
