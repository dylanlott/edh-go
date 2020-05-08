<template>
  <div class="mtg-card card">
    <img class="card-img-top" src="https://via.placeholder.com/1000x400.jpg">
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
  width: 175px;
  font-size: 14px;
}
div .card-body {
  line-height: .85rem;
  padding: 0.5em;
}
</style>
