import Vuex from 'vuex'
import Vue from 'vue'
import api from '@/gqlclient'
import gql from 'graphql-tag';
import router from '@/router';
import Cookies from 'js-cookie';


import {
    gameQuery,
    gameUpdateQuery,
    boardstateSubscription,
} from '@/gqlQueries'
import { boardstates, boardstatesSubscription } from '../gqlQueries';

Vue.use(Vuex)

const BoardStates = {
    state: {
        // boardstates is a map of user ID's to BoardStates.
        self: {
            User: {
                Username: "",
                ID: ""
            },
            GameID: "",
            Life: 0,
            Library: [],
            Commander: [],
            Field: [],
            Hand: [],
            Graveyard: [],
            Exiled: [],
            Revealed: [],
            Controlled: [],
            Counters: [],
        },
        boardstates: {},
        error: undefined
    },
    mutations: {
        error(state, payload) {
            state.error = payload
        },
        updateBoardStates(state, payload) {
            // update each boardstate by player ID
            payload.boardstates.forEach((bs) => {
                if (bs.User.ID == payload.selfID) {
                    state.self = bs
                }
                state.boardstates[bs.User.ID] = bs
            })
        },
    },
    actions: {
        mutateBoardStates({ commit }, payload) {
            // commit('update', payload)
            // Should we put this logic here or just update all boardstates
            // and make view logic handle which opponent sees what?
            // If we wanted to keep it separate, we could do 
            // different commit 
            // commit('updateSelf', payload)
            // commit('updateOpponents', payload)
        },
        // gets all boardstates from server, but doesn't subscribe
        getBoardStates({ commit }, gameID, selfID) {
            api.query({
                query: boardstates,
                variables: {
                    gameID: gameID
                }
            })
            .then((resp) => {
                commit('updateBoardStates', resp.data, selfID)
                return Promise.resolve(resp.data)
            })
            .catch((err) => {
                console.error("failed to get boardstates: ", err)
                commit('error', err)
                return Promise.reject(err)
            })
        },
        // used for subscribing to single board updates
        subscribeToBoardState({ state, commit }, payload) {
            console.log("subscribeToBoardStates#payload: ", payload)
            const sub = api.subscribe({
                query: boardstateSubscription,// TODO: Add the right query  
                variables: {
                    gameID: payload.gameID,
                    userID: payload.userID,
                    inputBoardState: state.self,
                },
            })
            sub.subscribe({
                next(data) {
                    console.log('BOARDSTATE SUBSCRIPTION DATA RECEIVED: ', data)
                },
                error(err) {
                    commit('error', err)
                    console.error('vuex action: boardstate subscription error: ', err)
                }
            })
        }
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
    },
    mutations: {
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
                const sub = api.subscribe({
                    query: gameUpdateQuery,// nb: this is where we use the subscription { } query
                    variables: {
                        game: {
                            ID: ID,
                            PlayerIDs: state.game.PlayerIDs.map((v) => {
                                return { Username: v.Username, ID: v.ID }
                            }),
                            Turn: state.Turn,
                        }
                    }
                })
                sub.subscribe({
                    next(data) {
                        console.log('GAME SUBSCRIPTION DATA RECEIVED: ', data)
                        commit('updateGame', data.data.gameUpdated)
                    },
                    error(err) {
                        commit('error', err)
                        console.log('vuex error: subscribeToGame: game subscription error: ', err)
                    }
                })
            })
        },
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
                commit('updateGame', res.data.joinGame)
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
            Username: Cookies.get("username") || window.localStorage.getItem("username"),
            ID: Cookies.get("userID") || window.localStorage.getItem("userID"),
            Token: Cookies.get("token") || window.localStorage.getItem("token")
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