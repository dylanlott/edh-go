<template>
  <div @click="tap()"
    class="mtg-card card"
    v-bind:class="{
      'd-none': hidden,
      'tapped': tapped
    }"
  >
    <!-- # TODO: Get images working for cards, but for now text will do.
    // Maybe hovering on a card should expose the
      <img class="card-img-top" src="https://via.placeholder.com/1000x400.jpg">
    -->
    <div class="card-body">
      <p class="card-title"><b>{{ name }}</b></p>
      <p class="card-text">{{ types }}</p>
      <p class="card-text">{{ text }}</p>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Card',
  data () {
    return {
      hidden: false,
      tapped: false,
      trackers: {},
      labels: {},
      counters: {},
      reminders: {},
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
  ],
  methods: {
    addCounter (name) {
      this.trackers[name]++
    },
    removeCounter (name) {
      this.trackers[name]--
    },
    addLabel (name, value) {
    },
    removeLabel (name) {

    },
    updateLabel (name, value) {

    },
    tap () {
      this.tapped = !this.tapped
    },
    moveTo (dst) {

    },
    flip () {

    },
  }
}
</script>
<style scoped media="screen">
.mtg-card {
  margin: .75rem -0.5rem;
  width: 175px;
  font-size: 14px;
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
