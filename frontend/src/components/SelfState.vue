<template>
  <div>
    <div class="columns">
      <div class="column">
        <div class="columns">
          <div class="column is-10">
            <p class="title is-5">Battlefield</p>
            <draggable
              class="card-wrapper"
              v-model="boardstate.battlefield"
              group="people"
              @start="drag=true"
              @end="drag=false">
              <div 
                v-for="card in boardstate.battlefield" 
                :key="card.id"
                class="columns">
                <Card v-bind="card"/>
              </div>
            </draggable>
          </div>
        </div>
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
    <hr>
    <div class="container shell">
      <p class="title is-4">Hand</p>
      <div class="columns">
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
    <!-- <code>{{ boardstate }}</code> -->
  </div>
</template>
<script>
import draggable from 'vuedraggable'
import Card from '@/components/Card'

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

export default {
  name: 'selfstate',
  data () {
    return {
      boardstate: {
        graveyard: [],
        library: [testCard2, testCard],
        exiled: [],
        hand: [],
        battlefield: [],
        emblems: [],
        revealed: []
      },
    }
  },
  watch: {
    boardstate: {
      handler(newState, oldState) {
        // TODO: Emit event here to graphQL that records boardstate mutations
        console.log('newState: ', newState)
        console.log('oldState: ', oldState)
      },
      deep: true
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
    draggable,
  }
}
</script>
<style media="screen">
  .library {
    hidden: true;
  }
</style>
