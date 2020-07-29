<template>
  <div @click="tap()"
    class="mtg-card card"
    v-bind:class="{
      'd-none': hidden,
      'tapped': tapped,
      'flipped': flipped
    }"
  >
    <!-- # TODO: Get images working for cards, but for now text will do.
    // Maybe hovering on a card should expose the
      <img class="card-img-top" src="https://via.placeholder.com/1000x400.jpg">
    -->
    <div class="card-body">
      <p class="card-title"><b>{{ name }}</b></p>
      <p class="card-text">{{ convertedManaCost }} - {{ manaCost }} </p>
      <p class="card-text">{{ types }}</p>
      <p class="card-text">{{ text }}</p>
      <div class="columns">
        <div class="column">{{ power }} / {{ toughness }}</div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Card',
  data () {
    return {
      hidden: false, // if a card can be seen at all - visibility off
      flipped: false, // if a card is upside down or not
    //   tapped: false, // if a card is tapped 
      trackers: {}, // player-assigned trackers 
      labels: {}, // player assigned labels
      counters: {}, // game-assigned counters such as poison or infect
      reminders: {}, // untap effects, etb effect reminders, etc...
    }
  },
  props: [
    'id',
    'name',
    'convertedManaCost',
    'manaCost',
    'colorIdentity',
    'power',
    'toughness',
    'text',
    'types',
    'supertypes',
    'subtypes',
    'types,',
    'image',
    'tapped',
  ],
  methods: {
    addCounter (name) {
      this.trackers[name]++
    },
    removeCounter (name) {
      this.trackers[name]--
    },
    addLabel (name, value) {
      this.labels.name = value
    },
    removeLabel (name) {
      delete this.labels.name
    },
    updateLabel (name, value) {
      this.labels.name
    },
    tap () {
      this.tapped = !this.tapped
    },
    moveTo (dst) {
    },
    flip () {
      this.flipped = !this.flipped
    },
  }
}
</script>
<style scoped media="screen">
.mtg-card {
  margin: .75rem 0rem;
  width: 175px;
  font-size: 12px;
}
div .card-body {
  line-height: .85rem;
  padding: 0.5em;
}

.tapped {
  -webkit-transform: rotate(90deg);
  -moz-transform: rotate(90deg);
  -o-transform: rotate(90deg);
  -ms-transform: rotate(90deg);
  transform: rotate(90deg);
}
</style>
