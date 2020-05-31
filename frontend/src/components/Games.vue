<template>
  <div class="games" v-if="gameID == ''">
    <h1>Welcome, {{ username }}</h1>
    <hr>
    <div class="hero">
      <div class="container">
        <b-field label="Pick your Commander.">
          <b-input v-model="deck.commander"></b-input>
        </b-field>
        <h3>Add the 99.</h3>
          <p>Note: There must be exactly 99 cards in this list, they need to be spelled exactly correct,
            and there can't be duplicates. Standard EDH rules and banlist applies. We recommend using
            <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling errors
            and quantities aren't an issue.
          </p>
          <div>
            <b-field label="Add the other 99 cards.">
              <b-input maxlength="200" v-model="deck.library" type="textarea"></b-input>
            </b-field>
          </div>
          <b-button @click="handleCreateGame()"
          type="button"
          class="is-success">Start a new game</b-button>
      </div>
    </div>

    <div class="container">
      <h2>Or join an existing game</h2>
      <p>Enter the game code below and your deck and you'll be
      joined into the game</p>
      <div class="">
        <b-input type="text"
        v-model="joinGameID"
        placeholder="Game ID"></b-input>
      </div>
      <b-button @click="handleJoinGame()"
      type="button"
      class="is-primary">Join Game</b-button>
    </div>
  </div>
</template>
<script>
import router from '@/router'
import gql from 'graphql-tag';

export default {
  name: 'game',
  data () {
    return {
      id: '',
      gameID: '',
      joinGameID: '',
      deck: {
        library: '',
        commander: ''
      }
    }
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      if (!vm.$currentUser()) {
        vm.$router.push('login');
      }
    });
  },
  apollo: {

  },
  computed: {
    players () {
      return [{
        id: 'player1',
        username: 'player1',
        deck: {
          commander: [],
          library: [],
        }
      }]
    },
    username () {
      return this.$currentUser()
    },
  },
  methods: {
    handleCreateGame() {
      this.$apollo.mutate({
        mutation: gql`mutation {
          createGame(input: {
            Players:[
              {
                Deck: {
                  name: "Karlov",
                  commander:"Karlov of the Ghost Council",
                  cards: [
                    "Teysa, Envoy of Ghosts"
                  ]
                },
                Username: "shakezula"
              }
            ]
          }){
           	id
            created_at3
        		players {
              GameID
              Commander {
                Name
              }
            }
          }
        }`,
        variables: {
          inputGame: {
            players: [{
              Deck: [],
              Username: 'shakezula'
            }]
          }
        }
      })
      router.push({ path: '/games/1234' })
    },
    handleJoinGame() {
      // this.$apollo.mutate({
      //   mutation: gql``,
      //   variables: {}
      // })
      router.push({ name: 'board', params: { id: this.joinGameID }})
    }
  },
  apollo: {

  }
}
</script>