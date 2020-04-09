<template>
  <div id="app">
    <main>
      <Navbar></Navbar>
      <router-view/>
    </main>
  </div>
</template>
<script>
import Navbar from '@/components/Navbar'

export default {
  name: 'main',
  components: {
    Navbar
  },
  created () {
  }
}
</script>
