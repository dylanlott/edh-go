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
      <!-- Gets the game, and only shows the Opponent boardstates from the Game PlayerIDs  -->
      <div :key="player.ID" v-for="player in game.PlayerIDs">
        <div v-if="player.Username !== self.User.Username">
          <h1 class="title">{{ player.Username }}</h1>
        </div>
        <div v-if="game.PlayerIDs.length === 1">
          <h1>No other players have joined this game.</h1>
        </div>
      </div>
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
      game: {
        PlayerIDs: [],
        Turn: {
          Player: this.$currentUser(),
          Phase: "setup",
          Number: 0
        },
      },
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
    // console.log('route ID: ', this.$route.params.id)
  },
  methods: {
    gameID() {
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
      this.$apollo.mutate({
        mutation: updateBoardStateQuery,
        variables: {
          boardstate: this.self.boardstate,
        },
      })
      .then((res) => {
        // console.log('mutateBoardState#setting boardstate: ', res.data.updateBoardState)
        this.self.boardstate = res.data.updateBoardState
        return res 
      })
      .catch((err) => {
        console.log('error mutating boardstate: ', err)
        return err
      })
    },
    mutateGameState() {
      // console.log
    },
    handleActivity(val) {
      return
      // console.log('logging activity: ', val)
    },
    sleepFor (sleepDuration) {
      var now = new Date().getTime();
      while(new Date().getTime() < now + sleepDuration){ /* do nothing */ } 
    },
    getPlayerIDs () {
      return this.game.PlayerIDs
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
    // Loads the user's state from the Route and UserID.
    // Selfstate is used for interacting with the Player's board.
    selfstate() {
      return {
        query: selfStateQuery,
        variables: {
          gameID: this.$route.params.id,
          userID: this.$currentUser(),
        },
        update(data) {
          this.self.boardstate = data.boardstates[0] 
        },
        results (data) {
          return data.boardstates[0] 
        }
      };
    },
    // Queries for the Game by route ID. This is responsible for loading up opponents, 
    // turn tracking, and eventually chat and metagame functionality.
    game() {
      return {
        query: gql`
          query	($gameID: String) {
            games(gameID: $gameID) {
              ID
              PlayerIDs {
                Username
                ID
              }
              Turn {
                Player
                Phase
                Number
              }
            }
          }
        `,
        variables: {
          gameID: this.gameID(),
        },
        update(data) {
          // this.$store.dispatch('')
          if (data.games.length === 0) {
            console.error(`game with ID ${this.gameID()}`)
            return []
          }
          return data.games[0]
        },
        // TODO: This code clobbers state. We need to subscribe to updates AFTER 
        // we have set the PlayerIDs correctly for a given GameID.
        // subscribeToMore: {
        //   // this should be the game updated subscription
        //   document: gql`subscription($game: InputGame!) {
        //     gameUpdated(game: $game) {
        //       ID
        //       PlayerIDs {
        //         Username
        //         ID
        //       }
        //       Turn {
        //         Player
        //         Phase
        //         Number
        //       }
        //     }  
        //   }`,
        //   variables: {
        //     game: {
        //       ID: this.$route.params.id,
        //       Turn: {
        //         Player: this.self.User.Username,
        //         Number: 0,
        //         Phase: "" 
        //       },
        //       PlayerIDs: this.getPlayerIDs(),
        //     }
        //   },
        //   updateQuery: (prevResult, {subData}) => {
        //     console.log('prevResult: ', prevResult)
        //     console.log('subData: ', subData)
        //     return subData[0]
        //   }
        // },
        results (data) {
          console.log('game query results: ', results)
        }
      } 
    }
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
