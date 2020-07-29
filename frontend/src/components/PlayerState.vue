<template id="">
  <div class="player state">
    <div class="columns">
      <div class="battlefield column">
        <span class="tag is-dark">Battlefield</span>
      </div>
      <div class="library column">
        <span class="tag is-dark">Library</span>
      </div>
      <div class="exiled column">
        <span class="tag is-dark">Exiled</span>
      </div>
      <div class="graveyard column">
        <span class="tag is-dark">Graveyard</span>
      </div>
      <div class="revealed column">
        <span class="tag is-dark">Revealed</span>
      </div>
      <div class="emblems column">
        <span class="tag is-dark">Emblems/Counters</span>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'playerstate',
  data () {
    return {
      gameID: this.$route.params.id,
    }
  },
  props: {
    library: Array,
    hand: Array,
    graveyard: Array,
    exiled: Array,
    battlefield: Array,
    controlled: Array,
    revealed: Array,
    emblems: Array
  },
  created () {
    console.log('PlayerState#gameID: ', this.$route.params.id)
  }
}
</script>
<style media="screen">
</style>
