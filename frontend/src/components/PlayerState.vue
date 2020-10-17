<template>
  <div class="player">
    <div v-for="(player, i) in opponents" :key="i">
      {{ player }} 
    </div> 
  </div>
</template>
<script>
export default {
  name: 'playerstate',
  data () {
    return {
      gameID: this.$route.params.id,
    }
  },
  created () {
    console.log('PlayerState#gameID: ', this.$route.params.id)
  },
  methods: {
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
  },
  apollo: {
    opponents() {
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
              playerIDs {
                username
                id
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
              },
              PlayerIDs: this.getPlayerIDs()
            }
          },
          updateQuery: (previousResult, { subscriptionData }) => {
            console.log('playerState#previousResult: ', previousResult)
            console.log('playerState#subscriptionData: ', subscriptionData)
          }
        } 
      }
    }
  }
}
</script>
<style media="screen">
</style>
