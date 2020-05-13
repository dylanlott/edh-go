<template>
  <div>
    <div class="row">
      <div class="col-lg">
        <p class="bg-dark text-center text-white">Battlefield</p>
        <draggable
        class="col-sm card-wrapper"
        v-model="boardstate.battlefield"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.battlefield" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
    </div>
    <div class="row">
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Exiled</p>
        <draggable
        class="col-sm card-wrapper"
        v-model="boardstate.exiled"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.exiled" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Graveyard</p>
        <draggable
        class="col-sm card-wrapper"
        v-model="boardstate.graveyard"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.graveyard" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Revealed</p>
        <draggable
        class="col-sm card-wrapper"
        v-model="boardstate.revealed"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.revealed" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Emblems/Counters</p>
        <draggable
        class="col-sm card-wrapper"
        v-model="boardstate.emblems"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.emblems" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="col-sm library" @click="draw()">
        <p class="bg-dark text-center text-white">Library</p>
        <draggable
        class="col-sm card-wrapper"
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
      <h3>Hand</h3>
      <div class="col-sm">
        <draggable
        class="row card-wrapper"
        v-model="boardstate.hand"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div class="mtg-card"
           v-for="card in boardstate.hand"
           :key="card.id">
            <Card v-bind="card"></Card>
           </div>
        </draggable>
      </div>
    </div>
    <code>{{ boardstate }}</code>
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
