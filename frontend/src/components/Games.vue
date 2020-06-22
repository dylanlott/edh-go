<template>
  <div class="games shell container" v-if="gameID == ''">
    <div class="hero">
      <div class="container">
        <div class="columns">
          <div class="column">
            <h1 class="title is-1">Welcome, {{ username }}</h1>
          </div>
        </div>
        <div class="columns">
          <div class="column">
            <p class="title is-4">Pick your commander</p>
            <p class="content"><b>Selected:</b> {{ selected }}</p>
            <b-field label="Find a JS framework">
              <b-autocomplete 
              rounded 
              v-model="deck.commander" 
              :data="commanderSearch" 
              placeholder="e.g. Jarad, Golgari Lich Lord" 
              icon="magnify"
              clearable 
              @select="option => selected = option">
                <template slot="empty">No results found</template>
              </b-autocomplete>
            </b-field>
            <b-input v-model="deck.commander"></b-input>
            <br></br>
            <h3 class="title is-4">Add the 99.</h3>
            <p>We recommend using
              <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling errors
              and quantities aren't an issue. Select "CSV" on Export.</p>
            <p><b>Note</b>: There must be exactly 99 cards in this list, they need to be spelled exactly correct,
              and there can't be duplicates except for Basic Lands.
            </p>
            <p>Standard EDH rules and banlist applies.</p> 
            <p>Enter cards in CSV Format: <code>1, Card Name</code></p>
            <b-field label="Add the other 99 cards.">
              <b-input maxlength="200000" v-model="deck.library" type="textarea"></b-input>
            </b-field>
          </div>
        </div>
        <b-button @click="handleCreateGame()" type="button" class="is-success">Start a new game</b-button>
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
      selected: {},
      deck: {
        library: '',
        commander: '',
        decklist: ''
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
    // NB: We use computed values to run validation and sanitization
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
    commander () {
      const trimmed = this.deck.commander.trim()
      const list = [{
        Name: trimmed
      }]
      return list 
    },
    commanderSearch () {
      this.$apollo.query({
        query: gql`query($name: String!) {
         search(name: $name) {
            Name
            ID
            Colors
            ColorIdentity
            CMC
            ManaCost
         }
        }`,
        variables: {
          name: this.deck.commander,
        }
      })
      .then((response) => {
        console.log('@ commander search response.data: ', response.data)
        return response.data
      })
    },
    library () {
      const split = this.deck.library.split('\n')
      console.log('split: ', split)
      const lib = split.map((card) => { 
        return { Name: card }
      })
      return lib
    },
    decklist () {
      return this.deck.library
    }
  },
  methods: {
    handleCreateGame() {
      console.log('creating game with deck: ', this.deck)
      console.log('using decklist: ', this.decklist)
      this.$apollo.mutate({
        mutation: gql`mutation ($inputGame: InputGame!) {
          createGame(input: $inputGame){
           	id
            created_at
            turn {
              Number
              Player
              Phase
            }
        		players {
              GameID
              Commander {
                Name
                ID
              }
            }
          }
        }`,
        variables: {
          inputGame: {
            ID: "",
            Turn: {
              Player: this.$currentUser(),
              Phase: "setup",
              Number: 0
            },
            Players: [{
              GameID: "",
              User: {
                Username: this.$currentUser()
              },
              Commander: this.commander,
              Library: [],
              Decklist: this.decklist,
              Graveyard: [],
              Exiled: [],
              Field: [],
              Hand: [],
              Revealed: [],
              Controlled: []
            }]
          }
        }
      })
      .then((res) => {
        console.log('@ Create Game Response: ', console.table(res))
        const id = res.data.id
        router.push({ path: `/games/${res.data.createGame.id}` })
      })
      .catch((err) => {
        console.error('got error back: ', err)
      })
    },
    handleJoinGame() {
      router.push({ name: 'board', params: { id: this.joinGameID }})
    }
  },
}
</script>