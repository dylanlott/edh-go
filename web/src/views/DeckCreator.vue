<template id="deckcreator">
  <section>
    <header>
      <h1>Deck Creator</h1>
      <p>Import a list of cards and we'll create a deck for you.</p>
      <p>Card names must be spelled exactly how they appear on the card.</p>
      <p>Each card name must be on its own line.</p>
      <p>You must have exactly <u>99 cards plus a commander for a total of 100 cards.</u></p>
      <p>We recommend using <a href="http://www.archidekt.com" target="_blank">Archidekt</a> to get this list of card names.</p>
    </header>
    <section>
      <v-form>
        <h3>Commander</h3>
        <p>Select the commander for this EDH deck</p>
        <input label="Commander">
        <textarea cols="50" rows="20"></textarea>
      </v-form>
    </section>
  </section>
</template>
<script>
export default {
  name: 'DeckCreator'
}
</script>
