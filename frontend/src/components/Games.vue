<template>
  <div class="" v-if="gameID == ''">
      <div class="container">
        <div class="columns is-mobile is-centered">
          <div class="column is-half">
            <h1 class="title is-1">Welcome, {{ username }}</h1>
            <p class="title is-4">Pick your commander</p>
            <!-- <p class="content" v-if="!!selected"><b>Commander:</b> {{ selected.name }}</p> -->
            <b-field label="Select a Commander">
              <b-autocomplete 
              :data="data"
              v-model="deck.commander"
              clearable 
              field="name"
              placeholder="e.g. Jarad, Golgari Lich Lord" 
              @typing="queryCommanders"
              @select="option => selected = option">
                <template slot="empty">No results found</template>
              </b-autocomplete>
            </b-field>
            <h3 class="title is-4">Add the 99.</h3>
            <p>We recommend using
              <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling errors
              and quantities aren't an issue. Select "CSV" on Export.</p>
              <br>
            <p><b>Note</b>: There must be exactly 99 cards in this list, they need to be spelled exactly correct,
              and there can't be duplicates except for Basic Lands.
            </p>
            <br>
            <p>Enter cards in CSV Format: <code>1, Card Name</code></p>
            <br>
            <b-field label="Paste your decklist here.">
              <b-input maxlength="200000" v-model="deck.library" type="textarea"></b-input>
            </b-field>

            <b-button 
              @click="handleCreateGame()" 
              type="button" 
              class="is-success">Start a new game</b-button>
          </div>
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
      selected: '',
      results: [],
      loading: false,
      deck: {
        library: '',
        commander: '',
        decklist: ''
      },
      data: []
    }
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      if (!vm.$currentUser()) {
        vm.$router.push('login');
      }
    });
  },
  computed: {
    commanderList() {
      this.queryCommanders()
        .then((data) => {
          return data
        })
        .catch((err) => console.error(err))
    },
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
    library () {
      const split = this.deck.library.split('\n')
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
      this.$apollo.mutate({
        mutation: gql`mutation ($inputGame: InputCreateGame!) {
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
              Life: 40,
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
        const id = res.data.id
        router.push({ path: `/games/${res.data.createGame.id}` })
      })
      .catch((err) => {
        console.error('got error back: ', err)
        return err
      })
    },
    handleJoinGame() {
      router.push({ name: 'board', params: { id: this.joinGameID }})
    },
    queryCommanders() {
      if (this.deck.commander === "") {
        return 
      }

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
      .then((resp) => {
        this.data = resp.data.search.map((item) => {
          return ({ name: item.Name })
        })
        return this.data
      })
      .catch((err) => {
        console.error('error querying commanders: ', err)
        return err
      })
      .finally(() => {
        this.loading = false
      })
    }
  },
}
</script>
<style>
  .shell {
    margin: 0.5rem;
  }
</style>