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
  },
};

export { AuthPlugin };
