import Vuex from 'vuex'
import Vue from 'vue'
import api from '@/gqlclient'
import gql from 'graphql-tag';
import router from '@/router'

import {
    gameQuery,
    gameUpdateQuery
} from '@/gqlQueries'

Vue.use(Vuex)

const BoardStates = {
    state: {
        // NB: Should self act as a cache to the server? 
        self: {},
        boardstates: {},
        loading: false,
        error: undefined
    },
    mutations: {
        request(state) {
            state.loading = true
            state.error = undefined
        },
        error(state, payload) {
            state.loading = false
            state.error = payload
        },
        update(state, payload) {
            state.loading = false
            state.boardstates = payload
        },
    },
    actions: {
        mutateBoardStates({ commit, state }, payload) {
            console.log('store#mutateBoardStates: ', payload)
            commit('update', payload)
            console.log("state: ", state.boardstates)
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
        // game should be identical in structure to the object we get 
        // back from the server. Thus the capitalization.
        game: {
            ID: "",
            Turn: {
                Player: "",
                Phase: "",
                Number: 0
            },
            PlayerIDs: []
        },
        error: undefined,
        loading: false,
    },
    mutations: {
        loading(state, payload) {
            state.loading = payload
        },
        error(state, err) {
            state.error = err 
        },
        updateGame(state, game) {
            state.game.ID = game.ID
            state.game.PlayerIDs = game.PlayerIDs.map((v) => {
                return { Username: v.Username, ID: v.ID }
            }),
            state.game.Turn = game.Turn
        },
        updateTurn(state, turn) {
            state.game.Turn = turn
        },
        gameFailure(state, error) {
            state.error = error
        },
        setStack(state, error) {

        },
    },
    actions: {
        getGame({ commit }, ID) {
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
        },
        subscribeToGame({ commit, state }, ID) {
            api.query({
                query: gameQuery,// TODO: Add the right query  
                variables: {
                    gameID: ID,
                }
            })
            .then(data => {
                if (data.data.games.length === 0) {
                    console.log('no game received from subscription')
                    return
                }
                commit('updateGame', data.data.games[0])
                api.subscribe({
                    query: gameUpdateQuery,// nb: this is where we use the subscription { } query
                    variables: {
                        game: {
                            ID: ID,
                            PlayerIDs: state.game.PlayerIDs.map((v) => {
                                return { Username: v.Username, ID: v.ID }
                            }),
                        }
                    }
                })
                .subscribe({
                    next(data) {
                        console.log('subscribeToGame received data: ', data)
                        commit('updateGame', data.data.games[0])
                        // self.game.PlayerIDs = data.data.gameUpdated.PlayerIDs
                        // self.game.Turn = data.data.gameUpdated.Turn
                    },
                    error(err) {
                        commit('error', err)
                        console.log('vuex error: subscribeToGame: game subscription error: ', err)
                    }
                })
            })
        },
        // TODO: Should joinGame simultaneously subscribe to Game?
        joinGame({ commit }, payload) {
            console.log('joinGame#payload: ', payload)
            api.mutate({
                mutation: gql`mutation ($InputJoinGame: InputJoinGame) {
                  joinGame(input: $InputJoinGame) {
                    ID
                    PlayerIDs {
                      Username
                    }
                  }
                }`,
                variables: {
                    InputJoinGame: payload.inputGame,
                }
            })
            .then((res) => {
                console.log('joinGame#res: ', res)
                commit('updateGame', res.data.games[0])
                router.push({ path: `/games/${res.data.joinGame.ID}` })
                return Promise.resolve(res)
            })
            .catch((err) => {
                commit('error', 'error joining game')
                console.log('error joining game: ', err)
                return err
            })
        },
        createGame({ state }, payload) {
            console.log('createGame#state: ', state)
            this.$apollo.mutate({
                mutation: gql`mutation ($inputGame: InputCreateGame!) {
                  createGame(input: $inputGame){
                    ID	
                    CreatedAt 
                    Turn {
                      Number
                      Player
                      Phase
                    }
                    PlayerIDs {
                      Username
                    }
                  }
                }`,
                variables: {
                    inputGame: payload.inputGame
                }
            })
            .then((res) => {
                console.log('pushing route to: ', res.data.createGame.ID)
                router.push({ path: `/games/${res.data.createGame.ID}` })
            })
            .catch((err) => {
                commit('error', err)
                console.error('got error back: ', err)
                return err
            })
        },
    }
}

const User = {
    state: {
        User: {
            Username: "",
            ID: "",
            Token: ""
        }
    },
    mutations:{
    },
    actions:{
        login() {
        },
        logout() {
        },
        getUser() {
        }
    }
}

const Cards = {
    state: {
        list: [],
        error: undefined,
    },
    mutations: {},
    actions: {},
}

const store = new Vuex.Store({
    modules: {
        BoardStates,
        Game,
        User,
        Cards,
    }
})

export default store