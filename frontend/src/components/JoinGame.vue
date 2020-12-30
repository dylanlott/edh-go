<template>
  <div class="join shell">
    <div class="container">
      <div class="columns is-mobile is-centered">
        <div class="column is-half">
          <h1 class="title is-1">Join Game</h1>
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
              @select="(option) => (selected = option)"
            >
              <template slot="empty">No results found</template>
            </b-autocomplete>
          </b-field>
          <h3 class="title is-4">Add the 99.</h3>
          <p>
            We recommend using <a href="www.archidekt.com">Archidekt</a> to generate your decklists so that spelling
            errors and quantities aren't an issue. Select "CSV" on Export.
          </p>
          <br />
          <p>
            <b>Note</b>: There must be exactly 99 cards in this list, they need to be spelled exactly correct, and there
            can't be duplicates except for Basic Lands.
          </p>
          <br />
          <p>Enter cards in CSV Format: <code>1, Card Name</code></p>
          <br />
          <b-field label="Paste your decklist here.">
            <b-input maxlength="200000" v-model="deck.library" type="textarea"></b-input>
          </b-field>

          <b-field label="Add a username?">
            <b-input v-model="username"></b-input>
          </b-field>

          <b-button @click="handleJoinGame()" type="button" class="is-success">Start a new game </b-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import gql from 'graphql-tag';
export default {
  name: 'join',
  data() {
    return {
      data: [],
      loading: false,
      deck: {
        library: '',
        commander: '',
        decklist: '',
      },
      username: '',
      game: {}
    };
  },
  computed: {
    commanderList() {
      this.queryCommanders()
        .then((data) => {
          return data;
        })
        .catch((err) => console.error(err));
    },
  },
  created () {
    this.gameQuery()
  },
  methods: {
    queryCommanders() {
      if (this.deck.commander === '') {
        return;
      }
      this.$apollo
        .query({
          query: gql`
            query($name: String!) {
              search(name: $name) {
                Name
                ID
                Colors
                ColorIdentity
                CMC
                ManaCost
              }
            }
          `,
          variables: {
            name: this.deck.commander,
          },
        })
        .then((resp) => {
          this.data = resp.data.search.map((item) => {
            return { name: item.Name };
          });
          return this.data;
        })
        .catch((err) => {
          console.error('error querying commanders: ', err);
          return err;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    handleJoinGame() {
      const inputGame = {
        ID: this.$route.params.id,
        PlayerIDs: this.game.PlayerIDs.push({
          ID: "",
          Username: "",
        }),
        Decklist: "",
        BoardState: {},
      }
      console.log('inputGame: ', inputGame)
      // this.$apollo.mutate({
      //   // TODO: write join game mutation
      //   mutation: gql`mutation {
      //     joinGame(input: inputJoinGame) {
      //       ID
      //     }
      //   }`,
      //   variables: {
      //     inputJoinGame: inputGame,
      //   },
      //   update: (store, { data }) => {
      //     console.log('handleJoinGame#update#store:', store)
      //     console.log('handleJoinGame#update#data:', data)
      //   }
      // })
    },
    gameQuery() {
      this.$apollo.query({
        query: gql`
          query	($gameID: String) {
            games(gameID: $gameID) {
              ID
              PlayerIDs {
                Username
                ID
              }
            }
          }
        `,
        variables: {
          gameID: this.$route.params.id,
        },
        update (data) {
          console.log('found game: ', data.games[0])
          this.game = data.games[0]
        }
      })
      .then(({ data }) => {
        console.log('this.gameQuery#then#data', data)
        if (data.games[0].length === 0) {
          console.log('failed to find game: ', this.$route.params.id)
        }
        this.game = data.games[0]
        console.log('this.game: ', this.game)
      })
    }
  },
};
</script>
<style>
</style>