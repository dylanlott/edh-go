<template>
  <div class="container">
    <div class="section">
      <div class="columns is-mobile is-centered">
        <div class="column is-half">
          <h1 class="title is-1">Signup for EDH-Go</h1>

          <b-field label="Username" :label-position="labelPosition">
            <b-input v-model="username"></b-input>
          </b-field>

          <b-field
            v-on:keyup.enter="handleSignup()"
            @submit="handleSignup()"
            label="Password"
            :label-position="labelPosition"
          >
            <b-input type="password" v-model="password"></b-input>
          </b-field>

          <b-button @submit="handleSignup()" @click="handleSignup()" v-on:keyup.enter="handleSignup()" type="is-primary"
            >Sign Up</b-button
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'signup',
  data: () => {
    return {
      labelPosition: 'on-border',
      username: '',
      password: '',
    };
  },
  methods: {
    handleSignup() {
      this.$store.dispatch('signup', {
        username: this.username,
        password: this.password,
      })
      .then(() => {
          this.username = ""
          this.password = ""
      })
      .catch((err) => {
          this.username = ""
          this.password = ""
      })
    },
  },
};
</script>

<style lang="sass">
</style>