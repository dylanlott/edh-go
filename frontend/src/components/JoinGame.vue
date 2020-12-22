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

          <b-button @click="handleCreateGame()" type="button" class="is-success"> Start a new game </b-button>
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
    handleCreateGame() {
      this.$apollo
        .mutate({
          mutation: gql`
            mutation($inputGame: InputCreateGame!) {
              createGame(input: $inputGame) {
                ID
                CreatedAt
                Turn {
                  Number
                  Player
                  Phase
                }
                PlayerIDs {
                  Username
                }
              }
            }
          `,
          variables: {
            inputGame: {
              ID: '',
              Players: [
                {
                  GameID: this.$route.params.id,
                  User: {
                    Username: this.username,
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
                  Controlled: [],
                },
              ],
            },
          },
        })
        .then((res) => {
          console.log('pushing route to: ', res.data.createGame.ID);
          router.push({ path: `/games/${res.data.createGame.ID}` });
        })
        .catch((err) => {
          console.error('got error back: ', err);
          return err;
        });
    },
  },
};
</script>
<style>
</style>