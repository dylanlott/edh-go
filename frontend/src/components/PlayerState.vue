<template id="">
  <div class="shell">
    <div class="graveyard">
      <p>Graveyard</p>
    </div>
  </div>
</template>
<script>
export default {
  name: 'playerstate',
  data () { return {}},
  created () {
    console.log('game id: ', this.$route.params.id)
  },
  props: [
    'boardstate'
  ]
}
</script>
