<template>
  <div>
    <div class="columns">
      <div class="column is-10">
        <p class="title is-5">Battlefield</p>
        <draggable
          class="card-wrapper"
          v-model="boardstate.field"
          group="people"
          @start="drag=true"
          @end="drag=false">
          <div 
            v-for="card in boardstate.field" 
            :key="card.id"
            class="columns">
            <Card v-bind="card"/>
          </div>
        </draggable>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <p class="title is-5">Exiled</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.exiled"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.exiled" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Graveyard</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.graveyard"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.graveyard" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Revealed</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.revealed"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.revealed" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Emblems/Counters</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.emblems"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.emblems" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column library" @click="draw()">
        <p class="title is-5">Library</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.library"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.library" :key="card.id">
             <Card v-bind="card" hidden="true"/>
           </div>
        </draggable>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <p class="title is-4">Hand</p>
        <draggable
        class="columns card-wrapper"
        v-model="boardstate.hand"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div 
            class="column mtg-card"
            v-for="card in boardstate.hand"
            :key="card.id">
            <Card v-bind="card"></Card>
           </div>
        </draggable>
      </div>
    </div>
    <hr>
    <code>{{ boardstates }}</code>
  </div>
</template>
<script>
import draggable from 'vuedraggable'
import Card from '@/components/Card'
import gql from 'graphql-tag'

const testCard = {
  id: '1',
  name: 'Karlov of the Ghost Council',
  convertedManaCost: '3',
  colorIdentity: 'BU',
  power: '7',
  toughness: '8',
  text: 'When this card enters the battlefield, make Brenden mill 10 cards.',
  types: 'Legendary Creature Wizard',
  image: '',
  counters: {}
}

const testCard2 = {
  id: '2',
  name: 'Ghost Council of Orzhova',
  convertedManaCost: '3',
  colorIdentity: 'BU',
  power: '7',
  toughness: '8',
  text: 'When this card enters the battlefield, make Brenden mill 10 cards.',
  types: 'Legendary Creature Wizard',
  image: '',
  counters: {}
}

const updateBoardStateQuery = gql`
  mutation ($boardstate: InputBoardState!) {
    updateBoardState(input: $boardstate) {
      User {
        username
      }
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
      Revealed {
        Name
      }
    }
  }
`

const getBoardstate = gql`
  query($gameID: String!) {
    boardstates(gameID: $gameID) {
      User {
        username
        id
      }
      Library {
        Name
        ID
      }
      Graveyard {
        Name
        ID
      }
      Exiled {
        Name
        ID
      }
      Field {
        Name
        ID
      }
      Hand {
        Name
        ID
      }
      Revealed {
        Name
        ID
      }
      Controlled {
        Name
        ID
      }
    }
  }
`

const boardstateSubscription = gql`
  subscription ($boardstate: InputBoardState!) {
    boardUpdate(boardstate: $boardstate) {
      GameID
      User {
        username
      }
    }
  }
`

export default {
  name: 'selfstate',
  data () {
    return {
      gameID: this.$route.params.id,
      boardstates: [],
      boardstate: {
        graveyard: [],
        library: [],
        exiled: [],
        hand: [],
        battlefield: [],
        emblems: [],
        revealed: [],
        // commander: [{
        //   Name: this.boardstate.Commander.Name,
        //   ID: this.boardstate.Commander.ID
        // }]
      },
    }
  },
  apollo: {
    boardstates() {
      return {
        query: getBoardstate,
        variables() {
          return {
            gameID: this.$route.params.id 
          }
        }
        // subscribeToMore: {
        //   document: boardstateSubscription,
        //   variables: {
        //     boardstate: {
        //       User: {
        //         Username: this.$currentUser()
        //       },
        //       Commander: [{
        //         Name: this.boardstate.Commander.Name
        //       }],
        //       GameID: this.$route.params.id,
        //     }
        //   },
        //   updateQuery: (prev, { subscriptionData }) => {
        //     console.log('selfstate # prev: ', prev)
        //     console.log('selfstate # subscription data: ', subscriptionData)
        //     return Object.assign({}, prev, subscriptionData)
        //   },
        // }
      }
    } 
  },
  methods: {
    draw () {
      if (this.boardstate.library.length == 0) {
        return
      }

      const card = this.boardstate.library.shift()
      this.boardstate.hand.push(card)
    },
    shuffle () {
      console.log('shuffling deck')
      // TODO: This should trigger a shuffle on the server and update the
      // library on the client side
    },
  },
  components: {
    Card,
    draggable
  }
}
</script>
<style media="screen">
  .library {
    visibility: true;
  }
</style>
