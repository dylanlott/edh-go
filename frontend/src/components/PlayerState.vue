<template id="">
  <div class="player state">
    <div class="row">
      <div class="graveyard col-sm">
        <p class="text-center bg-secondary text-white">Battlefield</p>
      </div>
      <div class="library col-sm">
        <p class="text-center bg-secondary text-white">Libarary</p>
      </div>
      <div class="exiled col-sm">
        <p class="text-center bg-secondary text-white">Exiled</p>
      </div>
      <div class="hand col-sm">
        <p class="text-center bg-secondary text-white">Graveyard</p>
      </div>
      <div class="revealed col-sm">
        <p class="text-center bg-secondary text-white">Revealed</p>
      </div>
      <div class="emblems col-sm">
        <p class="text-center bg-secondary text-white">Emblems/Counters</p>
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
      boardstate: {
        battlefield: [],
        hand: [],
        library: [],
        graveyard: [],
        exiled: [],
        stolen: [],
        attacking: [],
        blocking: [],
        stack: [],
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
