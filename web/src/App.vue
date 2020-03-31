<template>
  <div id="app">
    <main>
      <div id="nav">
        <router-link to="/">EDH Go</router-link> |
        <router-link to="/about">About</router-link>
      </div>
      <router-view/>
    </main>
  </div>
</template>
<script>
export default {
  name: 'main',
  components: {},
  created () {
    console.log('created')
  }
}
</script>
