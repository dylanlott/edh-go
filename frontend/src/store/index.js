import Vuex from 'vuex'
import Vue from 'vue'
import gql from 'graphql-tag';


Vue.use(Vuex)

const BoardStates = {
    state: {
        boardstate: {},
        loading: false,
        error: undefined
    },
    mutations: {
        request (state, payload) {
            state.loading = true
            state.error = undefined
        },
        error (state, payload) {
            state.loading = false
            state.error = payload
        },
        update (state, payload) {
            state.loading = false
            state.boardstate = payload
        }
    },
    actions: {
        mutateBoardState ({ commit }, payload) {
            console.log('vuex mutate boardstate hit: ', payload)
            console.log('apollo? ', this.$apollo)
            commit('update', payload)
        },
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
    mutations: {
        updateTurnRequest (state) {

        },
        updateTurnFailed (state, err) {

        },
        updateTurnSuccess (state, turn) {
            // state.Game.Turn 
        },
        opponentsRequest (state) {

        },
        opponentsSuccess (state, opps) {

        },
    },
    actions: {
        getGame({ commit }, ID) {
            console.log("getGame: ", ID)
        }
    }
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