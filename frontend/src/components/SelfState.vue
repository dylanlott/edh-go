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
      },
    }
  },
  apollo: {

  },
  methods: {
    draw () {
      console.log('this.boardstates: ', this.boardstates)
      if (this.boardstates[0].Library.length == 0) {
        return
      }

      const card = this.boardstates[0].Library.shift()
      this.boardstates[0].Hand.push(card)
    },
    shuffle () {
      console.log('TODO')
    },
    handleBoardUpdate() {
      console.log('TODO')
    }
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
