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
<style scope="global" lang="scss" media="screen">
@import "~bulma/sass/utilities/_all";

* {
}
// $primary: #3C34B2;
// $twitter: #18C572;
$primary: #000;
$twitter: #fff;
$primary-invert: findColorInvert($primary);
$twitter-invert: findColorInvert($twitter);

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
    "white": ($white, $black),
    "black": ($black, $white),
    "light": ($light, $light-invert),
    "dark": ($dark, $dark-invert),
    "primary": ($primary, $primary-invert),
    "info": ($info, $info-invert),
    "success": ($success, $success-invert),
    "warning": ($warning, $warning-invert),
    "danger": ($danger, $danger-invert),
    "twitter": ($twitter, $twitter-invert)
);

// Links
$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

</style>
