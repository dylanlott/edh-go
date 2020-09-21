<template>
  <div class="board shell">
    <h1 class="title shell">{{ gameID }}</h1>
    <TurnTracker gameID="gameID" />

  <!-- LIFE TRACKER -->
    <div>
      <div class="title is-4">{{ self.boardstate.Life }}</div>
      <button class="button" @click="increaseLife()">Increase</button>
      <button class="button" @click="decreaseLife()">Decrease</button>
    </div>

    <!-- OPPONENTS -->
    <div class="opponents">
      <div :key="b.id" v-for="b in boardstates" class="shell">
        <h1 class="title">{{ b.username }}</h1>
        <PlayerState v-bind="b"></PlayerState>
      </div>
    </div>
    <hr />

    <!-- STACK-->
    <div class="columns">
      <div 
      class="column is-full is-multiline">
        <draggable
          v-model="self.boardstate.Stack" 
        ></draggable>
      </div>
    </div>
    <!-- SELF -->
    <div class="self shell">
      <h1 class="title">
        {{ self.boardstate.User.username }}
        <p class="subtitle">{{ 
          self.boardstate.Commander ? self.boardstate.Commander[0].Name : ""
        }}</p>
      </h1>

      <div>
        <div class="columns">
          <div class="column is-full is-multiline">
            <p class="title is-5">Battlefield</p>
            <draggable
              class="card-wrapper bordered is-multiline battlefield"
              group="board" 
              v-model="self.boardstate.Field"
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div 
              class="is-multiline"
              @click="tap(card)"
              v-for="(card, i) in self.boardstate.Field" 
              :key="i" 
              >
                <Card v-bind="card" />
              </div>
            </draggable>
          </div>
        </div>
        <div class="columns">
          <!-- <div class="column">
            <p class="title is-5">Exiled</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Exiled"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in self.boardstate.Exiled" :key="card.id">
                <Card v-bind="card" @click="tap()"/>
              </div>
            </draggable>
          </div> -->
          <!-- <div class="column">
            <p class="title is-5">Graveyard</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Graveyard"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in self.boardstate.Graveyard" :key="card.id">
                <Card v-bind="card" />
              </div>
            </draggable>
          </div> -->
          <!-- <div class="column">
            <p class="title is-5">Revealed</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Revealed"
              group="board"
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in self.boardstate.Revealed" :key="card.id">
                <Card v-bind="card" @click="tap(card)"/>
              </div>
            </draggable>
          </div> -->
          <!-- <div class="column">
            <p class="title is-5">Emblems/Counters</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.emblems"
              group="board"
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in self.boardstate.emblems" :key="card.id">
                <Card v-bind="card" @click="tap(card)"/>
              </div>
            </draggable>
          </div> -->
          <div class="column library" @click="draw()">
            <p class="title is-5">Library</p>
            <draggable
              class="column card-wrapper"
              v-model="self.boardstate.Library"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div v-for="card in self.boardstate.Library" :key="card.id">
                <Card v-bind="card" hidden="true"/>
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
              @change="mutateBoardState()"
            >
              <div class="column mtg-card" v-for="card in self.boardstate.Hand" :key="card.id">
                <Card v-bind="card"></Card>
              </div>
            </draggable>
          </div>
        </div>
        <hr />
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
  updateBoardStateQuery,
  boardstates,
  boardstatesSubscription
} from '@/gqlQueries';
import router from '@/router'

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
      this.mutateBoardState()
    },
    mill() {
      const card = this.self.boardstate.Library.shift();
      this.self.boardstate.Graveyard.push(card);
      this.mutateBoardState()
    },
    increaseLife () {
      this.self.boardstate.Life++
      this.mutateBoardState()
    },
    decreaseLife() {
      this.self.boardstate.Life--
      this.mutateBoardState()
    },
    tap(card) {
      card.Tapped = !card.Tapped
      this.mutateBoardState()
    },
    mutateBoardState() {
      this.self.boardstate.User = {
        Username: this.$currentUser()
      }
      this.self.boardstate.GameID = this.$route.params.id
      console.log('right before update: ', this.self.boardstate)
      this.$apollo.mutate({
        mutation: updateBoardStateQuery,
        variables: {
          boardstate: this.self.boardstate,
        },
      })
      .then((res) => {
        const bs = res.data.updateBoardState
        this.self.boardstate.Life = bs.Life
        this.self.boardstate.Library = bs.Library
        this.self.boardstate.Commander = bs.Commander
        this.self.boardstate.Graveyard = bs.Graveyard
        this.self.boardstate.Exiled = bs.Exiled
        this.self.boardstate.Hand = bs.Hand
        this.self.boardstate.Revealed = bs.Revealed
        this.self.boardstate.Field = bs.Field
        this.self.boardstate.Controlled = bs.Controlled
        return res 
      })
      .catch((err) => {
        console.log('error mutating boardstate: ', err)
        return err
      })
    },
    handleActivity(val) {
      // console.log('logging activity: ', val)
    }
  },
  watch: {
    self: {
      handler (newVal, oldVal) {
        // we don't want to mutate state with this, 
        // or else we'll get infinite loops.
        // This is only where we should emit ActivityLog events.
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
          console.log('setting life: ', data.boardstates[0])
          this.self.boardstate = data.boardstates[0];
        },
      };
    },
    boardstates() {
      // TODO: get gameID and userID here so they're not tied to `self`
      // NB: This is where opponent boardstates come in to the Board.
      return {
        query: boardstates,
        variables: { gameID: this.$route.params.id },
        subscribeToMore: {
          document: boardstatesSubscription,
          variables: {
            boardstate: {
              User: {
                Username: this.$currentUser(),
              },
              Life: this.self.boardstate.Life ? this.self.boardstate.Life : 40,
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
          console.log('setting life: ', data.boardstates[0])
          console.log('subscription results: ', data);
        },
        error(err) {
          if (err == "Error: GraphQL error: game does not exist")  {
            // push to error page 
            router.push({ name: 'GameDoesNotExist'})
          }
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

.battlefield {
  height: 200px;
}
</style>
