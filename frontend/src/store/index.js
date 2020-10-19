import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

const BoardStates = {
    state: {
        Boardstate: {
            User: {
                Username: ""
            },
            Life: 0,
            Commander: [],
            Library: [],
            Graveyard: [],
            Exiled: [],
            Field: [],
            Hand: [],
            Revealed: [],
            Controlled: []
        }
    },
    mutations: {

    },
    actions: {

    }
}

const Game = {
    state: {
        Game: {
            ID: "",
            Turn: {
                Player: "",
                Phase: "",
                Number: 0
            },
            PlayerIDs: []
        }
    },
    mutations: {},
    actions: {}
}

const User = {
    state: {
        User: {
            Username: "",
            ID: "",
            Token: ""
        }
    }
}

const store = new Vuex.Store({
    modules: {
        BoardStates,
        Game,
        User
    }
})

export default store