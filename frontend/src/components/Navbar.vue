<template>
    <b-navbar>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
              <img src="../logo.png" alt="EDH Go">
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item href="#">
                Home
            </b-navbar-item>
            <b-navbar-item href="/games">
                Games
            </b-navbar-item>
        </template>

        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <a href="/signup" class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a tag="router-link"
                    href="/login"
                    class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>
<script>
export default {
  name: 'navbar',
  data () {
    return {}
  }
}
</script>
