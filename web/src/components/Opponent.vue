<template>
  <section>
    <aside class="opponent">
      <h3>Player {{ id }}</h3>
    </aside>
  </section>
</template>
<script>
export default {
  name: 'Opponent',
  props: ['id', 'boardstate']
}
</script>
