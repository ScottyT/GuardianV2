import {
  fireDb
} from '@/plugins/firebase';
import { auth } from '@/plugins/firebase';
import axios from 'axios';
import VueAxios from 'vue-axios';
export const state = () => ({
  header: {},
  footer: {},
  error: '',
  users: [],
  user: ''
})

export const mutations = {
  setHeader(state, header) {
    state.header = header
  },
  setFooter(state, footer) {
    state.footer = footer
  },
  setError(state, error) {
    state.error = error
  },
  setUsers: (state, payload) => {
    state.users = payload
  },
  userId: (state, payload) => {
    state.user = payload
  }  
}

export const actions = {
  async fetchMenu({
    commit
  }, $prismic) {
    try {
      const header = (await $prismic.api.getByUID('menu', 'header')).data
      const footer = (await $prismic.api.getByUID('menu', 'footer')).data
      commit('setHeader', header)
      commit('setFooter', footer)
    } catch (e) {
      const error = 'Please create a menu document'
      commit('setError', error)
    }
  },
  async fetchUsers({ commit }) {
    try {
      let users = []
      await fireDb.collection('users').get().then((qs) => {
        qs.forEach((doc) => {
          users.push(doc.data())
        })
      })
      commit('setUsers', users)
    } catch (e) {
      commit('setError', e)
    }
  },
  async fetchUserFavs({
    commit
  }, data) {
    var userid = data.id;
    try {
      await fireDb.collection("users").doc(userid).get()
        .then((doc) => {
          commit('project/setFavs', doc.data().favorites)
        })
        .catch((error) => {
          commit('setError', error)
        })
    } catch (e) {
      commit('setError', e)
    }
  },
  // async nuxtServerInit({ dispatch, commit }, { res, req }) {
  //   if (res && res.locals && res.locals.user) {
  //     console.log(res)
  //   }
  //   if (auth.currentUser != null) {
  //     var user = auth.currentUser.uid
  //     commit("userId", user)
  //     console.log(user)

  //   }
  //   console.log("hello")
  //   //console.log("Request: ", req.session)
  // }
  // nuxtServerInit: (process.server && !process.static) ? async function ({ commit, dispatch }, { req }) {
  //   if (!req.headers.cookie) return;
  //   const cookieparser = await import('cookie')
  //   const parsed = cookieparser.parse(req.headers.cookie);
  //   if (!parsed) {
  //     return;
  //   }
  //   const idToken = parsed['user_token']
  //   const JWTDecode = (await import('jwt-decode')).default;
  //   var decodedUser = null;
  //   try {
  //     decodedUser = JWTDecode(idToken)
  //   } catch { }
  //   if (decodedUser == null) {
  //     return;
  //   }
  //   console.log(decodedUser)
  //   commit('auth/setUser', {
  //     email: decodedUser.email
  //   })
  //   dispatch('project/fetchProjects', null, { root: true })
  // } : () => { console.log('other way') },
}
export const getters = {
  user(state, getters, rootState, rootGetters) {
    return rootGetters['auth/getUser']
  }
}
