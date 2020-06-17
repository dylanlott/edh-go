<template>
  <div class="games shell container is-fluid" v-if="gameID == ''">
    <div class="columns">
      <div class="column">
        <h1 class="title is-1">Welcome, {{ username }}</h1>
      </div>
    </div>
    <hr>
    <div class="hero">
      <div class="container">
        <div class="columns">
          <div class="column">
            <p class="title is-4">Pick your commander</p>
            <!-- TODO: Make this an autocomplete that pulls from a list of all commanders  -->
            <b-input v-model="deck.commander"></b-input>
            <br></br>
            <h3 class="title is-4">Add the 99.</h3>
            <p>Note: There must be exactly 99 cards in this list, they need to be spelled exactly correct,
              and there can't be duplicates. <br></br>
              Standard EDH rules and banlist applies. 
            </p>
            <p>
              We recommend using
              <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling errors
              and quantities aren't an issue.
            </p>
            <b-field label="Add the other 99 cards.">
              <b-input maxlength="200000" v-model="deck.library" type="textarea"></b-input>
            </b-field>
          </div>
        </div>
        <b-button @click="handleCreateGame()" type="button" class="is-success">Start a new game</b-button>
      </div>
    </div>

    <div class="container">
      <h2>Or join an existing game</h2>
      <p>Enter the game code below and your deck and you'll be
        joined into the game</p>
      <div class="">
        <b-input type="text" v-model="joinGameID" placeholder="Game ID"></b-input>
      </div>
      <b-button @click="handleJoinGame()" type="button" class="is-primary">Join Game</b-button>
    </div>

    <div class="columns">
      <div class="column">
        <code>{{ deck }}</code>
      </div>
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
      },
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
      console.log('creating game with deck: ', this.deck)
      this.$apollo.mutate({
        mutation: gql`mutation ($input: InputGame!) {
          createGame(inputGame: $input){
           	id
            created_at
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
            Players: [{
              Username: 'shakezula',
              Commander: [ ...this.deck.commander ],
              Library: [ ...this.deck.library ]
            }]
          }
        }
      })
      .then((data) => {
        console.log('GOT DATA!!!: ', data)
        $id = '1234' // TODO get id from response $gameID 
        router.push({ path: `/games/${id}` })
      })
      .catch((err) => {
        console.error('got error back: ', err)
      })
    },
    handleJoinGame() {
      // this.$apollo.mutate({
      //   mutation: gql``,
      //   variables: {}
      // })
      router.push({ name: 'board', params: { id: this.joinGameID }})
    }
  },
}
</script>