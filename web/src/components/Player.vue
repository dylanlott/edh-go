<template>
  <aside>
    <header>
      <h4>Player {{ id }}</h4>
    </header>
    <section>
      <h4>Battlefield</h4>
      {{ boardstate.battlefiedl }}
    </section>
    <section>
      <h4>Graveyard</h4>
      {{ boardstate.graveyard }}
    </section>
    <section>
      <h4>Exiled</h4>
      {{ boardstate.exiled }}
    </section>
    <code>
      {{ boardstate }}
    </code>
  </aside>
</template>
<script>
export default {
  name: 'Player',
  props: ['boardstate', 'id']
}
</script>
