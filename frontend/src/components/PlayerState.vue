<template id="">
  <div class="shell">
<div class="accordion" id="accordionExample">
  <div class="card">
    <div class="card-header" id="headingOne">
      <h4 class="mb-0">
        <button class="btn btn-link text-info" type="button" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
          Battlefield
        </button>
      </h4>
    </div>

    <div id="collapseOne" class="collapse show" aria-labelledby="headingOne" data-parent="#accordionExample">
      <div class="card-body">
        <div class="" v-for="card in boardstate.battlefield">
        </div>
        <div v-if="boardstate.battlefield.length < 1">No cards.</div>
      </div>
    </div>
  </div>
  <div class="card">
    <div class="card-header" id="headingTwo">
      <h4 class="mb-0">
        <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
          Graveyard
        </button>
      </h4>
    </div>
    <div id="collapseTwo" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionExample">
      <div class="card-body">
        <div v-if="boardstate.graveyard.length < 1">No graveyard cards.</div>
      </div>
    </div>
  </div>
  <div class="card">
    <div class="card-header" id="headingThree">
      <h4 class="mb-0">
        <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
          Exiled
        </button>
      </h4>
    </div>
    <div id="collapseThree" class="collapse" aria-labelledby="headingThree" data-parent="#accordionExample">
      <div class="card-body">
        <div v-if="boardstate.exiled.length < 1">No graveyard cards.</div>
      </div>
    </div>
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
