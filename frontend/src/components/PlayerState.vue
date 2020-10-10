<template id="">
  <div class="player">
    <p class="title is-4">(opponent)</p>
    <pre>{{ playerState }}</pre>
    <!-- <div class="columns is-mobile is-full">
      <div class="battlefield column">
        <div class="columns is-mobile">
          <div class="column" v-for="(card, i) in battlefield" :key="i"></div>
        </div>
      </div>
    </div> -->
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
  created () {
    console.log('PlayerState#gameID: ', this.$route.params.id)
  }
}
</script>
<style media="screen">
</style>
