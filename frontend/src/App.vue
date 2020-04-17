<template>
  <div class="">
    <Navbar></Navbar>
    <div class="container">
      <h1 class="mt-5 mb-4">EDH Go</h1>
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
