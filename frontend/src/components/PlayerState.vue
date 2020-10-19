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
      GameID: this.$route.params.id,
    }
  },
  methods: {
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
  },
  apollo: {
    opponents() {
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
        update(data) {
          this.Game = data.Games[0]
        },
        subscribeToMore: {
          document: gql`subscription($game: InputGame!) {
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
          }`,
          variables: {
            Game: {
              ID: this.$route.params.id,
              Turn: {
                Player: this.Game.Turn.Player || "",
                Phase: this.Game.Turn.Phase || "",
                Number: this.Game.Turn.Number || 0
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
