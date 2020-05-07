<template id="">
  <div class="shell">
    <div class="row">
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Battlefield</p>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Libarary</p>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Exiled</p>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Graveyard</p>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Revealed</p>
      </div>
      <div class="col-sm">
        <p class="bg-dark text-center text-white">Emblems/Counters</p>
      </div>
    </div>
    <div class="row shell">
      <p class="col-sm bg-dark text-center text-white">Hand</p>
      <div class="row">
      </div>
    </div>
  </div>
</template>
<script>
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
}

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
  },
  components: {
    Card
  }
}
</script>
