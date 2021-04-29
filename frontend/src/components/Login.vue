<template>
  <section class="shell">
    <div class="columns is-centered">
      <div class="column is-4">
        <form v-on:keyup.enter="onLoginClick()">
          <h1 class="title">Login</h1>
          <b-field label="Username">
            <b-input v-on:keyup.enter="onLoginClick()" v-model="username"></b-input>
          </b-field>
          <b-field label="Password">
            <b-input type="password" v-on:keyup.enter="onLoginClick()" v-model="password"></b-input>
          </b-field>
          <b-button v-on:keyup.enter="onLoginClick()" @click="onLoginClick()" type="submit" class="is-primary"
            >Log In</b-button
          >
        </form>
      </div>
    </div>
  </section>
</template>
<script>
import { mapState } from 'vuex';
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  computed: {
    ...mapState({
      loading: (state) => state.User.loading,
    }),
    isInputValid() {
      if (this.username.length === 0) {
        return false;
      }
      if (this.password.length === 0) {
        return false;
      }
      return true;
    },
  },
  methods: {
    onLoginClick() {
      if (this.isInputValid) {
        this.$store.dispatch('login', {
          username: this.username,
          password: this.password,
        });
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
