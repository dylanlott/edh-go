<template>
  <div class="container">
    <div class="columns">
      <div class="column">
        <p class="title is-1">Scores</p>
        <!-- First Row -->
        <div class="tile is-ancestor">
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Mana</p>
              <p class="subtitle">How much, how often, what color?</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Creatures</p>
              <p class="subtitle">What kind, what color, what power, what toughness?</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Life</p>
              <p class="subtitle">How much, how often, how little?</p>
            </article>
          </div>
        </div>
        <!-- Middle Row -->
        <div class="tile is-ancestor">
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Life</p>
              <p class="subtitle">How much, how often, how little?</p>
            </article>
          </div>
        </div>
        <!-- Third Row -->
        <div class="tile is-ancestor">
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Cards</p>
              <p class="subtitle">How much, how often, what color?</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Casts</p>
              <p class="subtitle">What kind, what color, what power, what toughness?</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Counters</p>
              <p class="subtitle">How much, how often, how little?</p>
            </article>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'score',
}
</script>

<style>

</style>