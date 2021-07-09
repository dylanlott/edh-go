<template>
  <div
    v-bind:class="{
      'd-none': hidden,
      tapped: Tapped,
      flipped: flipped,
    }"
  >
    <!-- # TODO: Get images working for cards, but for now text will do.
    // Maybe hovering on a card should expose the
      <img class="card-img-top" src="https://via.placeholder.com/1000x400.jpg">
    -->
    <div class="card mtg-card">
        <div class="card-content">
          <div class="content">
            {{ Name }}
          </div>
        </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Card',
  data() {
    return {
      hidden: false, // if a card can be seen at all - visibility off
      flipped: false, // if a card is upside down or not
      trackers: {}, // player-assigned trackers
      labels: {}, // player assigned labels
      counters: {}, // game-assigned counters such as poison or infect
      reminders: {}, // untap effects, etb effect reminders, etc...
    };
  },
  props: [
    'ID',
    'Name',
    'CMC',
    'ManaCost',
    'colorIdentity',
    'Power',
    'Toughness',
    'Text',
    'Types',
    'Supertypes',
    'Subtypes',
    'ScryfallID',
    'Tapped',
    'Flipped',
    'open',
  ],
  methods: {
    addCounter(name) {
      this.trackers[name]++;
    },
    removeCounter(name) {
      this.trackers[name]--;
    },
    addLabel(name, value) {
      this.labels.name = value;
    },
    removeLabel(name) {
      delete this.labels.name;
    },
    updateLabel(name, value) {
      this.labels.name;
    },
    moveTo(dst) {},
    flip() {
      this.flipped = !this.flipped;
    },
  },
};
</script>
