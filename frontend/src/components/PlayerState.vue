<template id="">
  <div class="player state">
    <div class="columns">
      <div class="battlefield column">
        <p class="title is-5">Battlefield</p>
      </div>
      <div class="library column">
        <p class="title is-5">Libarary</p>
      </div>
      <div class="exiled column">
        <p class="title is-5">Exiled</p>
      </div>
      <div class="graveyard column">
        <p class="title is-5">Graveyard</p>
      </div>
      <div class="revealed column">
        <p class="title is-5">Revealed</p>
      </div>
      <div class="emblems column">
        <p class="title is-5">Emblems/Counters</p>
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
      turn: {
        player: {},
        phase: {}
      },
      props: {
        library: Array,
        graveyard: Array,
        exiled: Array,
        battlefield: Array,
        controlled: Array
      }
    }
  },
  created () {
    console.log('game id: ', this.$route.params.id)
  },
}
</script>
<style media="screen">
</style>
