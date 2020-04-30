<template>
  <div class="" v-if="gameID == ''">
    <h1>Welcome, {{ username }}</h1>
    <h1>Join a game</h1>
    <div class="input-group">
      <input type="text"
      class="form-control"
      v-model="id"
      placeholder="Game ID">
    </div>
    <h3>Enter your decklist</h3>
    <div class="input-group">
      <input type="textarea"
      class="form-control"
      v-model="decklist">
    </div>
  </div>
</template>
<script>
export default {
  name: 'game',
  data () {
    return {
      id: '',
      gameID: '',
    }
  },
  beforeRouteEnter(to, from, next) {
    next((vm) => {
      if (!vm.$currentUser()) {
        vm.$router.push('login');
      }
    });
  },
  computed: {
    username () {
      return this.$currentUser()
    },
    decklist () {
      return window.localStorage.getItem('decklist') || []
    }
  },
  methods: {
  }
}
</script>
