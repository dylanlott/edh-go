const AuthPlugin = {
  install(Vue, options) {
    Vue.prototype.$setCurrentUser = function (user) {
      return window.localStorage.setItem('user', user)
    };
    Vue.prototype.$currentUser = function () {
      const user = window.localStorage.getItem('user')
      if (!user) {
        return undefined
      }
      return window.localStorage.getItem('user')
    };
    Vue.prototype.$logoutUser = function () {
      window.localStorage.removeItem('user')
      delete Vue.prototype.user
    };
    Vue.prototype.$setUsername = function (user) {
      Vue.$cookies.set('username', user)
    };
    Vue.prototype.$getUsername = function () {
      return Vue.$cookies.get('username')
    };
  },
};

export { AuthPlugin };
