<template>
  <div>
    <Navbar></Navbar>
      <router-view/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
export default {
  components: {
    Navbar
  }
};
</script>
<style lang="scss" media="screen">
@import "./scss/custom.scss";
</style>
