<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container">
      <a class="navbar-brand" href="/">EDH Go</a>
      <button class="navbar-toggler"
      type="button"
      data-toggle="collapse"
      data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent" 
      aria-expanded="false"
      aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
          <a class="nav-link" href="/">Home</a>
          <a class="nav-link" href="/games">Games</a>
          <a class="nav-link" href="/login">Sign Up</a>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script>
export default {
  name: 'navbar',
  data () {
    return {}
  }
}
</script>
