import Vuex from 'vuex'
import Vue from 'vue'
// import gql from 'graphql-tag'
import api from '@/gqlclient'

import {
    gameQuery,
} from '@/gqlQueries'

Vue.use(Vuex)

const BoardStates = {
    state: {
        boardstates: {},
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
            state.boardstates = payload
        },
    },
    actions: {
        mutateBoardStates ({ commit, state }, payload) {
            console.log('store#mutateBoardStates: ', payload)
            commit('update', payload)
            console.log("state? ", state)
            // Should we put this logic here or just update all boardstates
            // and make view logic handle which opponent sees what?
            // If we wanted to keep it separate, we could do 
            // different commit 
            // commit('updateSelf', payload)
            // commit('updateOpponents', payload)
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
        },
        Error: undefined ,
        Loading: false,
    },
    mutations: {
        loading (state, payload) {
            state.loading = payload
        },
        updateGame (state, game) {
            console.log('updateGame: ', game)
            // state.Game.ID = game.ID
        },
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
        gameFailure(state, error) {
            state.error = error
        }
    },
    actions: {
        getGame({ commit }, ID) {
            // console.log("api? ", api)
            console.log("getGame#ID: ", ID)
            api.query({
                query: gameQuery,// TODO: Add the right query  
                variables: {
                  gameID: ID,
                }
            }).then((data) => {
                commit('updateGame', data.data.games[0])
            }).catch((err) => {
                console.log('vuex failed to get game: ', err)
                commit('gameFailure', err)
            }) 
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