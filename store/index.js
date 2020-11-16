import {
  fireDb
} from '@/plugins/firebase';
export const state = () => ({
  header: {},
  footer: {},
  error: '',
  users: []
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
    console.log(userid)
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
  }
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
