<template>
  <div class="">
    <Navbar></Navbar>
    <div class="container">
      <router-view/>
    </div>
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
@import "node_modules/bootstrap/scss/bootstrap";

$theme-colors: (
  "primary": #3C34B2,
  "danger": #BC2E2E,
  "success": #18C572
);
</style>
