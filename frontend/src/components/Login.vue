<template>
  <section class="shell">
    <div class="columns is-centered">
      <div class="column is-4">
        <form v-on:keyup.enter="onLoginClick()">
          <h1 class="title">Login</h1>
          <b-field label="Login">
            <b-input 
              v-on:keyup.enter="onLoginClick()"
              v-model="username"></b-input>
          </b-field>
          <b-field label="Password">
            <b-input 
              v-on:keyup.enter="onLoginClick()"
              v-model="password"></b-input>
          </b-field>
          <b-button 
            v-on:keyup.enter="onLoginClick()"
            @click="onLoginClick()"
            type="submit"
            class="is-primary"
          >Log In</b-button>
         </form>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  computed: {
    isInputValid() {
      return this.username.length > 0;
    },
  },
  methods: {
    onLoginClick() {
      if (this.isInputValid) {
        this.$setCurrentUser(this.username);
        this.$router.push({ path: '/games' });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.form {
  max-width: 480px;
  margin: 0 auto;
}
</style>
