<template>
  <div class="board shell">
    <h1 class="title shell">{{ gameID }}</h1>
    <TurnTracker gameID="gameID" />

    <!-- OPPONENTS -->
    <div class="opponents">
      <div :key="b.id" v-for="b in boardstates" class="shell">
        <h1 class="title">{{ b.username }}</h1>
        <PlayerState v-bind="b"></PlayerState>
      </div>
    </div>
    <hr />

    <!-- SELF -->
    <div class="self shell">
      <h1 class="title">
        {{ self.boardstate.User.username }}
        <p class="subtitle">{{ self.boardstate.Commander[0].Name || undefined }}</p>
      </h1>

      <div>
        <div class="columns">
          <div class="column is-10">
            <p class="title is-5">Battlefield</p>
            <draggable
              class="card-wrapper bordered"
              v-model="self.boardstate.Field"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.Field" :key="card.id" class="columns">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
        </div>
        <div class="columns">
          <div class="column">
            <p class="title is-5">Exiled</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Exiled"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.Exiled" :key="card.id">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
          <div class="column">
            <p class="title is-5">Graveyard</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Graveyard"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.Graveyard" :key="card.id">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
          <div class="column">
            <p class="title is-5">Revealed</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Revealed"
              group="board"
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.Revealed" :key="card.id">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
          <div class="column">
            <p class="title is-5">Emblems/Counters</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.emblems"
              group="board"
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.emblems" :key="card.id">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
          <div class="column library" @click="draw()">
            <p class="title is-5">Library</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Library"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div v-for="card in self.boardstate.Library" :key="card.id">
                <Card v-bind="card" hidden="true" />
              </div>
            </draggable>
          </div>
        </div>
        <div class="columns">
          <div class="column">
            <p class="title is-4">Hand</p>
            <draggable
              class="columns card-wrapper is-vcentered"
              v-model="self.boardstate.Hand"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="handleUpdateState()"
            >
              <div class="column mtg-card" v-for="card in self.boardstate.Hand" :key="card.id">
                <Card v-bind="card"></Card>
              </div>
            </draggable>
          </div>
        </div>
        <hr />
        <!-- <code>{{ self }}</code> -->
      </div>
      <!-- END OF SELF STATE -->
    </div>

    <!-- CONTROL PANEL -->
    <div class="shell controlpanel columns">
      <div class="columns">
        <div class="column">
          <button class="button is-small is-primary">Collapse All</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Untap</button>
        </div>
        <div class="column">
          <button @click="draw()" class="button is-small is-primary">Draw</button>
        </div>
        <div class="column">
          <button class="button is-small is-primary">Shuffle</button>
        </div>
        <div class="column">
          <button @click="mill()" class="button is-small is-primary">Mill</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import _ from 'lodash';
import gql from 'graphql-tag';
import draggable from 'vuedraggable';
import Card from '@/components/Card';
import PlayerState from '@/components/PlayerState.vue';
import SelfState from '@/components/SelfState.vue';
import TurnTracker from '@/components/TurnTracker.vue';
import { 
  selfStateQuery, 
  updateBoardStateQuery 
} from '@/gqlQueries';

export default {
  name: 'board',
  data() {
    return {
      gameID: this.$route.params.id,
      locked: false, // `locked` is set to true once the players and turn order are decided.
      mulligan: true, // `mulligan` is set to true until no one is mulling anymore.
      self: {
        user: {
          username: this.$currentUser(),
        },
        boardstate: {
          User: {
            username: this.$currentUser(),
          },
          Commander: [],
          Library: [],
          Field: [],
        },
      },
    };
  },
  computed: {
    username: (state) => this.$currentUser(),
  },
  methods: {
    draw() {
      const card = this.self.boardstate.Library.shift();
      this.self.boardstate.Hand.push(card);
      this.handleUpdateState()
    },
    mill() {
      const card = this.self.boardstate.Library.shift();
      this.self.boardstate.Graveyard.push(card);
      this.handleUpdateState()
    },
    handleUpdateState() {
      const self = this
      console.log('updating state: ', self)
      // _.throttle(this.mutateBoardState, 500)
      this.mutateBoardState()
    },
    mutateBoardState() {
      const self = this
      self.self.boardstate.User = {
        Username: this.$currentUser()
      }
      self.self.boardstate.GameID = this.$route.params.id

      this.$apollo.mutate({
        mutation: updateBoardStateQuery,
        variables: {
          boardstate: self.self.boardstate,
        },
      })
      .then((data) => {
        console.log('successfully mutated state: ', data)
      })
    },
    handleActivity(val) {
      console.log('logging activity: ', val)
    }
  },
  watch: {
    self: {
      handler (newVal, oldVal) {
        // we don't want to mutate with this, or else we'll get infinite loops.
        // instead we should only send out activity log events here.
        this.handleActivity(newVal)
      },
      deep: true
    }
  },
  apollo: {
    selfstate() {
      return {
        query: selfStateQuery,
        variables: {
          gameID: this.$route.params.id,
          userID: this.$currentUser(),
        },
        update(data) {
          this.self.boardstate = data.boardstates[0];
          console.table('selfstate#update: ', this.self.boardstate);
        },
      };
    },
    boardstates() {
      // get gameID and userID here so they're not tied to `self`
      return {
        query: gql`
          query($gameID: String!) {
            boardstates(gameID: $gameID) {
              Commander {
                Name
              }
              Library {
                Name
              }
              Graveyard {
                Name
              }
              Exiled {
                Name
              }
              Field {
                Name
              }
              Hand {
                Name
              }
              Revealed {
                Name
              }
              Controlled {
                Name
              }
            }
          }
        `,
        variables: { gameID: this.$route.params.id },
        subscribeToMore: {
          document: gql`
            subscription($boardstate: InputBoardState!) {
              boardUpdate(boardstate: $boardstate) {
                GameID
                Commander {
                  Name
                }
                Library {
                  Name
                }
                Graveyard {
                  Name
                }
                Exiled {
                  Name
                }
                Field {
                  Name
                }
                Hand {
                  Name
                }
                Revealed {
                  Name
                }
                Controlled {
                  Name
                }
              }
            }
          `,
          variables: {
            boardstate: {
              User: {
                Username: this.$currentUser(),
              },
              GameID: this.$route.params.id,
              Commander: this.self.boardstate.Commander ? [...this.self.boardstate.Commander] : [],
              Library: this.self.boardstate.Library ? [...this.self.boardstate.Library] : [],
              Graveyard: this.self.boardstate.Graveyard ? [...this.self.boardstate.Graveyard] : [],
              Exiled: this.self.boardstate.Exiled ? [...this.self.boardstate.Exiled] : [],
              Field: this.self.boardstate.Field ? [...this.self.boardstate.Field] : [],
              Hand: this.self.boardstate.Hand ? [...this.self.boardstate.Hand] : [],
              Revealed: this.self.boardstate.Revealed ? [...this.self.boardstate.Revealed] : [],
              Controlled: this.self.boardstate.Controlled ? [...this.self.boardstate.Controlled] : [],
            },
          },
        },
        results({ data }) {
          console.log('subscription results: ', data);
        },
        error(err) {
          console.log('error getting boardstates: ', err);
          const notif = this.$buefy.notification.open({
            duration: 5000,
            message: `Error occurred when fetching opponents boardstates. Check your game ID and try again.`,
            position: 'is-top-right',
            type: 'is-danger',
            hasIcon: true,
          });
        },
      };
    },
  },
  components: {
    draggable,
    Card,
    TurnTracker,
    PlayerState,
    SelfState,
  },
};
</script>
<style media="screen" scoped>
.shell {
  padding: 0.5rem;
  border: 1px solid #efefef;
  margin: 0.25rem 0rem;
}

.bordered {
  border: 1px #000;
}
</style>
