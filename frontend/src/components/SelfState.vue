<template>
  <div>
    {{ boardstate }}
    <!-- <div class="columns">
      <div class="column is-10">
        <p class="title is-5">Battlefield</p>
        <draggable
          class="card-wrapper"
          v-model="boardstate.Field"
          group="people"
          @start="drag = true"
          @end="drag = false"
        >
          <div v-for="card in boardstate.Field" :key="card.id" class="columns">
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
        v-model="boardstate.Exiled"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.Exiled" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Graveyard</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.Graveyard"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.Graveyard" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column">
        <p class="title is-5">Revealed</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.Revealed"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.Revealed" :key="card.id">
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
           <div v-for="card in boardstate.Emblems" :key="card.id">
             <Card v-bind="card"/>
           </div>
        </draggable>
      </div>
      <div class="column library" @click="draw()">
        <p class="title is-5">Library</p>
        <draggable
        class="column card-wrapper"
        v-model="boardstate.Library"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div v-for="card in boardstate.Library" :key="card.id">
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
        v-model="boardstate.Hand"
        group="people"
        @start="drag=true"
        @end="drag=false">
           <div 
            class="column mtg-card"
            v-for="card in boardstate.Hand"
            :key="card.id">
            <Card v-bind="card"></Card>
           </div>
        </draggable>
      </div>
    </div>
    <hr /> -->
  </div>
</template>
<script>
import draggable from 'vuedraggable';
import Card from '@/components/Card';
import { mapState } from 'vuex';

export default {
  name: 'selfstate',
  data() {
    return {
      gameID: this.$route.params.id,
    };
  },
  props: {
    boardstate: Object,
  },
  methods: {
    draw() {
      console.log('TODO');
    },
    shuffle() {
      console.log('TODO');
    },
    handleBoardUpdate() {
      console.log('handle board update');
      this.$apollo
        .mutate({
          // TODO
        })
        .catch((res) => {
          console.log('pushed updated board state: ', res);
          return res;
        })
        .catch((err) => {
          console.log('error pushing board update: ', err);
          return err;
        });
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
