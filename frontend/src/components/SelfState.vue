<template>
  <div>
    <!-- {{ boardstate }} -->
    <div class="columns">
      <div class="column is-10">
        <p class="title is-5">Battlefield</p>
        <draggable
          class="card-wrapper"
          v-model="self.Field"
          group="people"
          @start="drag = true"
          @end="drag = false"
        >
          <div v-for="card in self.Field" :key="card.id" class="columns">
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
        v-model="self.Exiled"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in self.Exiled" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Graveyard</p>
        <draggable
        class="column card-wrapper"
        v-model="self.Graveyard"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in self.Graveyard" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Revealed</p>
        <draggable
        class="column card-wrapper"
        v-model="self.Revealed"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in self.Revealed" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Emblems/Counters</p>
        <draggable
        class="column card-wrapper"
        v-model="self.emblems"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in self.Emblems" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column library" @click="draw(self)">
        <p class="title is-5">Library</p>
        <draggable
        class="column card-wrapper"
        v-model="self.Library"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in self.Library" :key="card.id">
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
        v-model="self.Hand"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div 
            class="column mtg-card"
            v-for="card in self.Hand"
            :key="card.id">
            <Card v-bind="card"></Card>
           </div>
        </draggable>
      </div>
    </div>
    <hr />
  </div>
</template>
<script>
import draggable from 'vuedraggable';
import { mapState } from 'vuex';
import Card from '@/components/Card';
export default {
  name: 'selfstate',
  data() {
    return {
      gameID: this.$route.params.id,
    };
  },
  props: {
    playerID: String,
  },
  computed: {
    self() {
      return this.$store.getters.self
    },
  ...mapState({
    boardstates: state => state.BoardStates.boardstates,
  }),
  },
  methods: {
    draw() {
      console.log(this.playerID, this.boardstates)
      const bs = Object.assign({}, this.boardstates[this.playerID])
      console.log('bs: ', bs)
      const card = bs.Library[0]
      bs.Library.shift()
      bs.Hand.push(card)
      console.log('boardstate: ', bs)
    },
    mutateBoardState(bs) {
      return this.$store.dispatch('mutateBoardState', bs);
    },
  },
  components: {
    Card,
    draggable,
  },
};
</script>
<style media="screen">
.library {
  visibility: true;
}
</style>
