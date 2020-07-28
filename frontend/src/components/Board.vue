<template>
  <div class="board shell">
    <h1 class="title shell">{{ gameID }}</h1>
    <TurnTracker gameID="gameID"/>
    <div class="opponents">
      <div :key="b.id" v-for="b in boardstates" class="shell">
        <h1 class="title">{{ b.username }}</h1>
        <PlayerState v-bind="b"></PlayerState>
      </div >
    </div>
    <hr>
    <div class="self shell">
      <h1 class="title">{{ self.username }}</h1>
      <SelfState
        :boardstate="self.boardstate"
      ></SelfState>
    </div>
    <div class="shell controlpanel columns">
      <div class="columns">
        <div class="column">
          <button class="button is-small is-primary">Collapse All</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Untap</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Draw</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Shuffle</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import gql from 'graphql-tag';
import PlayerState from '@/components/PlayerState.vue'
import SelfState from '@/components/SelfState.vue'
import TurnTracker from '@/components/TurnTracker.vue'

export default {
  name: 'board',
  data () {
    return {
      // TODO: This needs to be modeled after BoardState
      gameID: this.$route.params.id,
      locked: false,  // `locked` is set to true once the players and turn order are decided.
      mulligan: true, // `mulligan` is set to true until no one is mulling anymore.
      self: {
        username: this.$currentUser(),
        boardstate: {
          commander: [],
          library: [],
          graveyard: [],
          exiled: [],
          battlefield: [],
          hand: [],
          controlled: [],
          revealed: [],
        }
      },
    }
  },
  computed: {
    username: (state) => this.$currentUser(),
  },
  apollo: {
    selfstate() {
      return {
        query: gql`
          query($gameID: String!, $userID: String) {
            boardstates(gameID: $gameID, userID: $userID) {
              Commander{ Name }
              Library { Name }
              Graveyard { Name }
              Exiled { Name }
              Field { Name }
              Hand { Name }
              Revealed { Name }
              Controlled { Name } 
            }
          }
        `,
        variables: ({ 
          gameID: this.$route.params.id,
          userID: this.$currentUser()
        }),
        update(data) {
          this.selfstate = data.boardstates[0]
          console.table(this.selfstate)
        }
      }
    },
    boardstates() {
      // get gameID and userID here so they're not tied to `self` 
      return {
        query: gql`
          query($gameID: String!) {
            boardstates(gameID: $gameID) {
              Commander{ Name }
              Library { Name }
              Graveyard { Name }
              Exiled { Name }
              Field { Name }
              Hand { Name }
              Revealed { Name }
              Controlled { Name } 
            }
          }
        `,
        variables: ({ gameID: this.$route.params.id }),
        subscribeToMore: {
          document: gql`
            subscription($boardstate: InputBoardState!) {
              boardUpdate(boardstate: $boardstate) {
                GameID
                Commander{ Name }
                Library { Name }
                Graveyard { Name }
                Exiled { Name }
                Field { Name }
                Hand { Name }
                Revealed { Name }
                Controlled { Name } 
              }
            }
          `,
          variables: () => {
            const vars = {
              boardstate: {
                User: {
                  Username: this.$currentUser()
                },
                GameID: this.$route.params.id,
                Commander: [ ...this.self.boardstate.commander ],
                Library: [ ...this.self.boardstate.library ],
                Graveyard: [ ...this.self.boardstate.graveyard ],
                Exiled: [...this.self.boardstate.exiled ],
                Field: [...this.self.boardstate.battlefield ],
                Hand: [...this.self.boardstate.hand ],
                Revealed: [...this.self.boardstate.revealed ],
                Controlled: [...this.self.boardstate.controlled]
              },
            }
            return vars
          },
          updateQuery: (prev, { subscriptionData }) => {
            console.log('opponents()#previous: ', prev)
            console.log('opponents()#subscriptionData: ', subscriptionData)
          }
        },
        error(err) {
          console.log('error getting boardstates: ', err)
          const notif = this.$buefy.notification.open({
            duration: 5000,
            message: `Error occurred when fetching opponents boardstates. Check your game ID and try again.`,
            position: 'is-top-right',
            type: 'is-danger',
            hasIcon: true
          })
        }
      }
    }
  },
  components: {
    TurnTracker,
    PlayerState,
    SelfState,
  },

}
</script>
<style media="screen" scoped>
.shell {
  padding: .5rem;
  border: 1px solid #efefef;
  margin: .25rem 0.0rem;
}
</style>
