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
      labels: []
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
    'counters',
  ],
  methods: {
    addCounter (name) {
      this.trackers[name]++
    },
    removeCounter (name) {
      this.trackers[name]--
    },
    addLabel (name, value) {
    }
  }
}
</script>
<style scoped media="screen">
.mtg-card {
  width: 150px;
  font-size: 14px;
}
div .card-body {
  line-height: .85rem;
  padding: 0.5em;
}
</style>
