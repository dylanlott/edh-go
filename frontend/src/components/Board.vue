<template>
  <div class="board shell">
    <!-- <h1 class="title shell">{{ gameID }}</h1> -->

  <!-- LIFE TRACKER -->
  <div class="columns">
    <div class="shell column is-9">
      <!-- <TurnTracker gameID="GameID" /> -->
    </div>
    <div class="shell column is-3">
      <div class="title is-4">{{ self.boardstate.Life }}</div>
      <button class="button is-small" @click="increaseLife()">Increase</button>
      <button class="button is-small" @click="decreaseLife()">Decrease</button>
    </div>
  </div>

    <!-- OPPONENTS -->
    <div class="opponents">
      <!-- <div :key="b.id" v-for="o in opponents" class="shell"> -->
        <!-- <pre> {{ b }} </pre> -->
        <!-- <h1 class="title">{{ b.username }}</h1> -->
        <!-- <PlayerState v-bind="b"></PlayerState> -->
      <!-- </div> -->
    </div>
    <hr />

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
          <div class="column">
            <p class="title is-5">Battlefield</p>
            <draggable
              class="card-wrapper bordered battlefield"
              group="board" 
              v-model="self.boardstate.Field"
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div 
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
        </div>
        <div class="columns">
          <div class="column hand is-three-quarters">
            <p class="title is-4">Hand</p>
            <draggable
              class="columns card-wrapper"
              v-model="self.boardstate.Hand"
              group="board" 
              @start="drag = true"
              @end="drag = false"
              @change="mutateBoardState()"
            >
              <div class="column mtg-card" v-for="(card, i) in self.boardstate.Hand" :key="i">
                <Card v-bind="card"></Card>
              </div>
            </draggable>
          </div>
          <div class="column library is-one-quarter" @click="draw()">
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
// import TurnTracker from '@/components/TurnTracker.vue';
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
      locked: false, // `locked` is set to true once the players and turn order are decided.
      mulligan: true, // `mulligan` is set to true until no one is mulling anymore.
      self: {
        GameID: this.$route.params.id,
        User: {
          Username: this.$currentUser(),
        },
        boardstate: {
          GameID: this.$route.params.id,
          User: {
            Username: this.$currentUser(),
          },
        },
      },
    };
  },
  created () {
    console.log('route ID: ', this.$route.params.id)
  },
  methods: {
    routeGameID() {
      return this.$route.params.id
    },
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

    // @param `src` is the source field of cards the target card is in. 
    // @param `target` is the card that's being fetched
    // @param `dst` is the destination field of the fetched card
    // NB: We always want to pass cards around by ID, since we're 
    // planning on these being unique.
    // @returns: `src`, `dst`
    fetch (src, target, dst) {
      let obj = src.find((v, idx)=> {
        if (v.ID === target.ID) {
          console.log(`target found, moving ${target} from ${src} -> ${dst}`)
          src2 = src.splice(1, idx)
          dst2 = dst.push(v)
          console.log(`target moved: src: ${src2} dst: ${dst2}`)
          // TODO: this.mutateBoardState()
          return src2, dst2
        }
      })
      if (obj === undefined) {
        console.error(`unable to find target ${target}`)
        return src, dst
      }
      // if we get here, we have a weird result. 
      // log and return src.
      console.log('weird, we shouldnt be here', src, dst)
      return src, dst
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
      console.log('mutateBoardState#mutating with boardstate: ', this.self.boardstate)
      this.$apollo.mutate({
        mutation: updateBoardStateQuery,
        variables: {
          boardstate: this.self.boardstate,
        },
      })
      .then((res) => {
        console.log('mutateBoardState#setting boardstate: ', res.data.updateBoardState)
        this.self.boardstate = res.data.updateBoardState
        return res 
      })
      .catch((err) => {
        console.log('error mutating boardstate: ', err)
        return err
      })
    },
    handleActivity(val) {
      // console.log('logging activity: ', val)
    },
    sleepFor (sleepDuration) {
      var now = new Date().getTime();
      while(new Date().getTime() < now + sleepDuration){ /* do nothing */ } 
    }
  },
  watch: {
    self: {
      handler (newVal, oldVal) {
        // we don't want to mutate state with this, 
        // or else we'll get infinite loops.
        // This is only where we should emit ActivityLog events.
        // this.handleActivity(newVal)
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
          var updated = Object.assign(this.self.boardstate, data.boardstates[0])
          console.log('selfstate#updated object: ', updated)
          this.self.boardstate = updated
        },
        results (data) {
          console.log("selfstate results: ", data)
          return data
        }
      };
    },
    // boardstates() {
    //   // TODO: This is not correctly async with the route ID. Maybe I need to 
    //   // refactor these into my own methods and call them individually instead of
    //   // relying on Apollo's auto-call-on-load magic. 

    //   // TODO: get gameID and userID here so they're not tied to `self`
    //   // NB: This is where opponent boardstates come in to the Board.
    //   return {
    //     query: boardstates,
    //     variables: { gameID: this.routeGameID() },
    //     subscribeToMore: {
    //       document: boardstatesSubscription,
    //       // updateQuery: (previousResult, { subscriptionData }) => {
    //       //   console.log('previousResult: ', previousResult)
    //       //   console.log('subscriptionData: ', subscriptionData)
    //       // },
    //       variables: {
    //         boardstate: {
    //           User: {
    //             Username: this.$currentUser(),
    //           },
    //           Life: this.self.boardstate.Life ? this.self.boardstate.Life : 40,
    //           GameID: this.self.boardstate.GameID,
    //           Commander: this.self.boardstate.Commander ? [...this.self.boardstate.Commander] : [],
    //           Library: this.self.boardstate.Library ? [...this.self.boardstate.Library] : [],
    //           Graveyard: this.self.boardstate.Graveyard ? [...this.self.boardstate.Graveyard] : [],
    //           Exiled: this.self.boardstate.Exiled ? [...this.self.boardstate.Exiled] : [],
    //           Field: this.self.boardstate.Field ? [...this.self.boardstate.Field] : [],
    //           Hand: this.self.boardstate.Hand ? [...this.self.boardstate.Hand] : [],
    //           Revealed: this.self.boardstate.Revealed ? [...this.self.boardstate.Revealed] : [],
    //           Controlled: this.self.boardstate.Controlled ? [...this.self.boardstate.Controlled] : [],
    //         },
    //       },
    //     },
    //     results(data) {
    //       console.log('boardstates#data: ', data)
    //       return data
    //     },
    //     error(err) {
    //       console.log('HIT BOARDSTATES ERROR:', err)
    //       if (err == "Error: GraphQL error: game does not exist")  {
    //         // push to error page 
    //         router.push({ name: 'GameDoesNotExist'})
    //       }
    //       console.log('error getting boardstates: ', err);
    //       const notif = this.$buefy.notification.open({
    //         duration: 5000,
    //         message: `Error occurred when fetching opponents boardstates. Check your game ID and try again.`,
    //         position: 'is-top-right',
    //         type: 'is-danger',
    //         hasIcon: true,
    //       });
    //     },
    //   };
    // },
  },
  components: {
    draggable,
    Card,
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
