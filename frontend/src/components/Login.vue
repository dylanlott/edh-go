<template>
  <div class="row">
    <form class="col-12 form"
          v-on:submit.prevent="onLoginClick">

          <div>
            <h1>Signup for EDH Go</h1>
          </div>
      <div class="input-group">
        <input type="text"
               class="form-control"
               :class="{'is-invalid': !isInputValid}"
               placeholder="Username..."
               v-model.trim="userName"
               required>
      </div>
      <div class="input-group">
        <input type="text"
          class="form-control"
          placeholder="Password"
          v-model.trim="password"
        >
        <div class="input-group-append">
          <button class="btn btn-secondary"
                  type="submit">Log in</button>
        </div>
      </div>
      <hr>
      <div>
        <h2>Sign up if you haven't already</h2>
        <div class="input-group">
          <input type="text"
          class="form-control"
          placeholder="Email">
          <input type="text"
          class="form-control"
          placeholder="Password">
        </div>
      </div>

    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: '',
      password: '',
      email: ''
    };
  },
  computed: {
    isInputValid() {
      return this.userName.length > 0;
    },
  },
  methods: {
    onLoginClick() {
      if (this.isInputValid) {
        this.$setCurrentUser(this.userName);
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
