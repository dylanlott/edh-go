<template id="">
  <div class="shell">
    <div class="row">
      <div class="graveyard col-sm">
        <p>Graveyard</p>
      </div>
      <div class="library col-sm">
        <p>Libarary</p>
      </div>
      <div class="Exiled col-sm">
        <p>Exiled</p>
      </div>
    </div>
    <div class="row">
    <div class="Hand col-sm">
      <p>Hand</p>
    </div>
    <div class="battlefield col-sm">
      <p>Battlefield</p>
    </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'selfstate',
  data () {
    return {
      graveyard: [],
      library: [],
      exiled: [],
      hand: [],
      battlefield: []
    }
  }
}
</script>
