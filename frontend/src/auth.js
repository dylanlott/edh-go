const AuthPlugin = {
  install(Vue, options) {
    Vue.prototype.$setCurrentUser = function (user) {
      Vue.prototype.user = user;
    };
    Vue.prototype.$currentUser = function () {
      return Vue.prototype.user;
    };
    Vue.prototype.$logoutUser = function (user) {
      Vue.prototype.user = {};
    };
  },
};

export { AuthPlugin };
