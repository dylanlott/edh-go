const AuthPlugin = {
  install(Vue, options) {
    Vue.prototype.$setCurrentUser = function (user) {
      return Vue.$cookies.set('username', user)
    };
    Vue.prototype.$currentUser = function () {
      return Vue.$cookies.get('username')
    };
    Vue.prototype.$currentUserID = function () {
      return Vue.$cookies.get('userID')
    };
    Vue.prototype.$logoutUser = function () {
      Vue.$cookies.set('username', '')
      Vue.$cookies.set('userID', '')
      Vue.$cookies.set('token', '')
    };
    Vue.prototype.$setToken = function (token) {
      return Vue.$cookies.set('token', token)
    }
  },
};

export { AuthPlugin };
