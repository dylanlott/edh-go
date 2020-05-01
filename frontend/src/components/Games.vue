<template>
  <div class="" v-if="gameID == ''">
    <h1>Welcome, {{ username }}</h1>
    <hr>
    <div class="jumbotron">
      <div class="container">
        <h2>Pick your Commander.</h2>
        <div class="input-group">
          <input type="text" class="form-control"
          v-model="deck.commander">
        </div>
        <h3>Add the 99.</h3>
          <p>Note: There must be exactly 99 cards in this list, they need to be spelled exactly correct,
            and there can't be duplicates. Standard EDH rules and banlist applies. We recommend using
            <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling errors
            and quantities aren't an issue.
          </p>
          <div class="form-group">
            <label for="exampleFormControlTextarea1">Add the rest of your 99 cards.</label>
            <textarea class="form-control" id="exampleFormControlTextarea1" rows="3"></textarea>
          </div>
          <button @click="handleCreateGame()"
          type="button"
          class="btn btn-success">Start a new game</button>
      </div>
    </div>

    <div class="container">
      <h2>Or join an existing game</h2>
      <p>Enter the game code below and your deck and you'll be
      joined into the game</p>
      <div class="input-group">
        <input type="text"
        class="form-control"
        v-model="id"
        placeholder="Game ID">
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
      deck: {
        library: [],
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
    // TODO: Make sure this works.
    decks() {
      const user = this.$currentUser();
      return {
        query: gql`
        {
          decks {
            id
            user
            name
          }
        }
        `
      }
    }
  },
  computed: {
    players: [],
    username () {
      return this.$currentUser()
    },
  },
  methods: {
    handleCreateGame() {
      this.$apollo.mutate({
        mutation: gql``,
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
    }
  },
  apollo: {

  }
}
</script>
