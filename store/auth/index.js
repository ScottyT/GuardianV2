import { auth } from '@/plugins/firebase';
import { fireDb } from '@/plugins/firebase';
import { persistance } from '@/plugins/firebase';
import Cookie from 'js-cookie';
export const state = () => ({
  user: null,
  error: null,
  authModal: {
    visibility: false,
    callbackAction: {
      name: null,
      payload: null
    }
  }
});

export const mutations = {
  setUser: (state, payload) => {
    state.user = payload
  },
  setError: (state, payload) => {
    state.error = payload
  },
  updateAuthModalVisibility: (state, payload) => {
    state.authModal.visibility = payload.visibility
  }
}

export const actions = {
  async login({ commit }, account) {
    try {
      // await auth.setPersistence(persistance.Persistence.LOCAL).then(() => {
      //   return auth.signInWithEmailAndPassword(account.email, account.password).catch((error) => {
      //       console.log(error.message)
      //       commit('setError', error.message);
      //     });
      //   })
      //   .catch((error) => {
      //     commit('setError', error.message)
      //   })
      
      let id = '';
      let data = '';
      let user = '';
      await fireDb.collection('users').get().then((querySnapshot) => {
        querySnapshot.forEach((doc) => {
          id = doc.id;
          data = doc.data()
        })
      })
      const loginRef = fireDb.collection('users').doc(id);
      await loginRef.get().then((doc) => {
        if (doc.exists) {
          user = doc.data()
        } else {
          console.log("No such document exists")
        }
      }).catch((error) => {
        commit("setError", error)
      })
      // const token = await auth.currentUser.getIdToken();
      // Cookie.set('user_token', token, {
      //   expires: 1
      // });
      //commit('setUser', data)
      await auth.signInWithEmailAndPassword(account.email, account.password).then(() => {
        commit('setUser', data)
      }).catch((error) => {
        commit('setError', error.message)
      })
    } catch (error) {
      throw error;
    }
    
  },
  showAuthModal({ commit }) {
    commit("updateAuthModalVisibility", { visibility: true });
  },
  hideAuthModal({ commit }) {
    commit("updateAuthModalVisibility", { visibility: false });
  }
}

export const getters = {
  getError: (state) => {
    return state.error
  }
}