<template>
  <div
    class="mtg-card card"
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
    <div class="card-body">
      <b-collapse class="card" animation="slide" aria-id="contentIdForA11y3">
        <template #trigger="props">
          <div class="card-header" role="button" aria-controls="contentIdForA11y3">
            <p class="card-header-title">
              {{ Name }}
            </p>
            <a class="card-header-icon">
              <b-icon v-if="CMC">
                {{ CMC }}
              </b-icon>
              <b-icon>
                {{ ManaCost }}
              </b-icon>
              <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"> </b-icon>
            </a>
          </div>
        </template>
        <div class="card-content">
          <div class="content">
            <p class="card-text">{{ Text }}</p>
            <div class="columns" v-if="Power || Toughness">
              <div class="column">{{ Power }} / {{ Toughness }}</div>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <!-- <a class="card-footer-item">Save</a>
                <a class="card-footer-item">Edit</a>
                <a class="card-footer-item">Delete</a> -->
        </footer>
      </b-collapse>
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
<style scoped media="screen">
.mtg-card {
  margin: 0.75rem 0rem;
  width: 175px;
  font-size: 12px;
}
div .card-body {
  line-height: 0.85rem;
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
