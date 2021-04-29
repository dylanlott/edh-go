import Vuex from 'vuex'
import Vue from 'vue'
import gql from 'graphql-tag';
import Cookies from 'js-cookie';
import { uuid } from '@/uuid';
import api from '@/gqlclient'
import router from '@/router';

const ls = window.localStorage

import {
    gameQuery,
    gameUpdateQuery,
    boardstates,
    boardstateSubscription,
} from '@/gqlQueries'

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
            console.log('updateBoardStates payload: ', payload)
            // update each boardstate by player ID
            payload.boardstates.forEach((bs) => {
                state.boardstates[bs.User.ID] = bs
                // assign our own boardstate to `self` for easier control
                if (bs.User.Username == payload.self) {
                    state.self = bs
                }
            })
        },
        updateBoardstate(state, payload) {
            if (payload.ID == "" || payload.ID == undefined) {
                state.error = "boardstate did not have ID"
                return
            }
            
            // assign boardstates to keys by their ID
            state.boardstates[payload.ID] = payload
        }
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
        getBoardStates({ commit, state, rootState }, gameID) {
            api.query({
                query: boardstates,
                variables: {
                    gameID: gameID
                }
            })
            .then((resp) => {
                console.log('getBoardStates committing: ', resp.data)
                commit('updateBoardStates', {
                    boardstates: resp.data.boardstates, 
                    self: rootState.User.User.Username,
                })
                return resp.data
            })
            .catch((err) => {
                console.error("failed to get boardstates: ", err)
                commit('error', err)
                return err
            })
        },
        // used for subscribing to single board updates
        subscribeToBoardState({ state, commit, rootState }, payload) {
            // self.GameID == "" since it's not assigned before that query fires
            // we need to assign that first 
            state.self.GameID = payload.gameID
            const sub = api.subscribe({
                query: boardstateSubscription,// TODO: Add the right query  
                variables: {
                    inputBoardState: state.self,
                },
            })
            sub.subscribe({
                next(data) {
                    console.log('BOARDSTATE SUBSCRIPTION DATA RECEIVED: ', data)
                    // commit('updateBoardStates', {
                    //     boardstates: data.data.boardstatePosted,
                    //     self: rootState.User.User.Username,
                    // })
                    commit('updateBoardstate', data.data.boardstatePosted)
                },
                error(err) {
                    commit('error', err)
                    console.error('subscribeToBoardstate: boardstate subscription error: ', err)
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
        loading: false,
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
                        console.log('subscribeToGame is updating game with ', data.data.gameUpdated)
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
            api.mutate({
                mutation: gql`mutation ($InputJoinGame: InputJoinGame) {
                  joinGame(input: $InputJoinGame) {
                    ID
                    PlayerIDs {
                      Username
                      ID
                    }
                    Turn {
                        Phase
                        Player
                        Number
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
        createGame({ commit }, payload) {
            commit('loading', true)
            return new Promise((resolve, reject) => {
                api.mutate({
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
                            ID
                        }
                    }
                    }`,
                    variables: {
                        inputGame: payload
                    }
                })
                .then((res) => {
                    commit('updateGame', res.data.createGame)
                    router.push({ path: `/games/${res.data.createGame.ID}` })
                    return resolve(res)
                })
                .catch((err) => {
                    commit('error', 'error creating game')
                    console.error('error createGame: ', err)
                    return reject(err)
                })
            })
        },
    }
}

const User = {
    state: {
        User: {
            Username: Cookies.get("username") || ls.getItem("username"),
            ID: Cookies.get("userID") || ls.getItem("userID") || uuid(),
            Token: Cookies.get("token") || ls.getItem("token")
        },
        loading: false,
        error: undefined,
    },
    mutations:{
        setUser(state, payload) {
            state.User.Username = payload.Username
            Cookies.set("username", payload.Username)
            state.User.ID = payload.ID
            Cookies.set("userID", payload.ID)
            state.User.Token = payload.Token
            Cookies.set("token", payload.Token)
            Cookies.set("user_info", JSON.stringify(payload))
        },
        loading(state, bool) {
            state.loading = bool
        },
        error(state, message) {
            state.error = message
        }
    },
    actions:{
        login({ commit }, payload) {
            commit('loading', true)
            api.mutate({
                mutation: gql`mutation($username: String!, $password: String!) {
                    login(username: $username, password: $password) {
                        Username
                        ID
                        Token 
                    }
                }`,
                variables: {
                    username: payload.username,
                    password: payload.password,
                }
            }) 
            .then((data) => {
                commit('setUser', data.data.login)
                router.push({ path: '/games' });
                return Promise.resolve(data) 
            })
            .catch((err) => {
                console.error('login error: ', err)
                commit('error', 'failed to login')
                Promise.reject(err)
            })
        },
        logout() {
            // TODO: Clear cookies and localStorage and delete token on server
            console.log('logout action hit')
        },
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