import Vuex from 'vuex'
import Vue from 'vue'
import gql from 'graphql-tag'
import Cookies from 'js-cookie'
import { uuid } from '@/uuid'
import { ToastProgrammatic as Toast } from 'buefy'
import api from '@/gqlclient'
import router from '@/router'
import {
    login,
    signup,
    gameQuery,
    gameUpdateQuery,
    updateGame,
    boardstates,
    boardstateSubscription,
} from '@/gqlQueries'
import { updateBoardStateQuery } from '../gqlQueries';

Vue.use(Vuex)

const ls = window.localStorage

const BoardStates = {
    // BoardStates should only ever be updated by push from the server, we 
    // don't normally want to update these ourselves. 
    state: {
        // boardstates holds an object with all player boardstates
        // keyed by the player's UUID (ID)
        boardstates: {},
        // we load our personal state into self but we still don't want to 
        // violate our one-direction data flow principle
        self: {},
        // error is an error message from the API. 
        // TODO: Make these decay 
        error: undefined
    },
    mutations: {
        error(state, payload) {
            state.error = payload
            Toast.open({
                message: `Boardstate error: ${state.error}`,
                duration: 3000,
                position: "is-bottom",
                type: "is-danger",
            })
        },
        // takes an array of boardstates and updates each of them
        // payload is iterated over and checked for it's user's ID
        // and then assigned to the boardstates object keyed by that ID
        updateBoardStates(state, payload) {
            // update each boardstate by player ID
            payload.forEach((bs) => {
                console.log('updating boardstate:', bs)
                if (bs.User.ID == "" || bs.User.ID == undefined) {
                    state.error = "boardstate did not have ID"
                    return
                }
                state.boardstates[bs.User.ID] = bs
            })
        },
        updateSelf(state, payload) {
            state.self = payload
        }
    },
    actions: {
        draw({ commit, dispatch }, boardstate) {
            const bs = Object.assign({}, boardstate)
            console.log('draw#boardstate:', bs)
            if (bs.Library.length < 1) {
                // handle player losing issue
                commit('error', 'you cannot draw from an empty library. you lose the game.')
                console.error('cannot draw from an empty library.')
            }
            const card = bs.Library.shift()
            console.log('card: ', card)
            bs.Hand.push(card)
            dispatch('mutateBoardState', bs)
        },
        // mutate boardstate will mutate a boardstate and then commit the result.
        mutateBoardState({ commit }, payload) {
            api.mutate({
                mutation: updateBoardStateQuery,
                variables: {
                    boardstate: payload,
                },
            })
                .then((resp) => {
                    commit('updateBoardStates', [resp.data.updateBoardState])
                })
                .catch((err) => {
                    console.error('mutateBoardState: error updating boardstate: ', err)
                    commit('error', 'something went wrong.')
                })
        },
        // gets all boardstates from server, but doesn't subscribe
        getBoardStates({ state, commit, rootState }, gameID) {
            return new Promise((resolve, reject) => {
                api.query({
                    query: boardstates,
                    variables: {
                        gameID: gameID
                    }
                })
                .then((resp) => {
                    console.log('GET BOARDSTATES SUCCEEDED: ', resp)
                    resp.data.boardstates.forEach((boardstate) => {
                        if (boardstate.User.ID === rootState.User.User.ID) {
                            console.log("self detected: ", boardstate)
                            state.boardstates[rootState.User.User.ID] = boardstate
                            commit('updateSelf', boardstate)
                        } else {
                            commit('updateBoardStates', resp.data.boardstates)
                        }
                    })
                    // Note: we don't need to put this resp into an array because its already a list
                    return resolve(resp.data)
                })
                .catch((err) => {
                    console.error('GET BOARDSTATES FAILED: ', err)
                    commit('error', err)
                    return reject(err)
                })
            })
        },
        // used for subscribing to single board updates
        subToBoardstate({ rootState, commit }, payload) {
            const sub = api.subscribe({
                query: boardstateSubscription,// TODO: Add the right query  
                variables: {
                    gameID: payload.gameID,
                    userID: payload.userID,
                },
            })
            sub.subscribe({
                next(data) {
                    console.log("received boardstate update: ", data.data.boardstateUpdated)
                    // detect self vs opponents here and assign accordingly 
                    if (data.data.boardstateUpdated.User.ID == rootState.User.User.ID) {
                        console.log("SELF RECEIVED: ", data.data.boardstateUpdated)
                    }
                    commit('updateBoardStates', [data.data.boardstateUpdated])
                },
                error(err) {
                    commit('error', err)
                    console.error('subscribeToBoardstate: boardstate subscription error: ', err)
                }
            })
        },
    },
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
            Toast.open({
                message: `Game error: ${state.error}`,
                duration: 3000,
                position: "is-bottom",
                type: "is-danger",
            })
        },
        updateGame(state, game) {
            // merge updated game over current game
            const g = Object.assign(state.game, game)
            console.table('updateGame mutation: ', g)
            state.game = g
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
            return new Promise((resolve, reject) => {
                api.query({
                    query: gameQuery,// TODO: Add the right query  
                    variables: {
                        gameID: ID,
                    }
                }).then((resp) => {
                    commit('updateGame', resp.data.games[0])
                    return resolve(resp.data.games[0])
                }).catch((err) => {
                    console.error('vuex failed to get game: ', err)
                    commit('gameFailure', err)
                    return reject(err)
                })
            })
        },
        subscribeToGame({ state, commit }, ID) {
            api.query({
                query: gameQuery,
                variables: {
                    gameID: ID,
                }
            })
            .then(data => {
                if (data.data.games.length === 0) {
                    commit('error', 'no game received from subscription')
                    return
                }
                const sub = api.subscribe({
                    query: gameUpdateQuery,// nb: this is where we use the subscription { } query
                    variables: {
                        gameID: ID,
                    }
                })
                sub.subscribe({
                    next(data) {
                        commit('updateGame', data.data.gameUpdated)
                    },
                    error(err) {
                        console.error('vuex error: subscribeToGame: game subscription error: ', err)
                        commit('error', err)
                    }
                })
            })
        },
        joinGame({ commit }, payload) {
            return new Promise((resolve, reject) => {
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
                    console.log("updating game after joing: ", res.data.joinGame)
                    commit('updateGame', res.data.joinGame)
                    router.push({ path: `/games/${res.data.joinGame.ID}` })
                    return resolve(res)
                })
                .catch((err) => {
                    commit('error', 'error joining game')
                    console.error('error joining game: ', err)
                    return reject(err)
                })
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
        updateGame({ commit }, payload) {
            return api.mutate({
                mutation: updateGame,
                variables: {
                    input: payload,
                }
            })
            .then((data) => {
                console.log('updateGame is setting: ', data.data.updateGame)
                commit('updateGame', data.data.updateGame)
                return data
            })
            .catch((err) => {
                console.error('updateGame failed: ', err)
                return err
            })
        }
    },
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
    mutations: {
        setUser(state, payload) {
            state.User.Username = payload.Username
            Cookies.set("username", payload.Username)
            state.User.ID = payload.ID
            Cookies.set("userID", payload.ID)
            state.User.Token = payload.Token
            Cookies.set("token", payload.Token)
            Cookies.set("user_info", JSON.stringify(payload))
        },
        logout(state) {
            // set back to default state
            state.User = {Username: "", Token: "", ID: ""}
            // and then remove all cookies
            Cookies.remove("username")
            Cookies.remove("userID")
            Cookies.remove("token")
            Cookies.remove("user_info")
        },
        loading(state, bool) {
            state.loading = bool
        },
        error(state, message) {
            state.error = message
            Toast.open({
                message: `User error: ${state.error}`,
                duration: 3000,
                position: "is-bottom",
                type: "is-danger",
            })
        }
    },
    actions: {
        login({ commit }, payload) {
            commit('loading', true)
            api.mutate({
                mutation: login,
                variables: {
                    username: payload.username,
                    password: payload.password,
                }
            })
            .then((data) => {
                commit('setUser', data.data.login)
                router.push({ path: '/games' });
                return data
            })
            .catch((err) => {
                console.error('login error: ', err)
                commit('error', 'failed to login')
                return err
            })
        },
        logout({ commit }) {
            commit('logout')
            router.push({ path: '/' })
        },
        signup({ commit, dispatch }, payload) {
            commit('loading', true)
            api.mutate({
                mutation: signup,
                variables: {
                    username: payload.username,
                    password: payload.password,
                }
            })
            .then((resp) => {
                dispatch('login', payload) 
                .catch((err) => commit('error', 'failed to login'))
            })
            .catch((err) => {
                commit('error', 'failed to signup')
            })
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