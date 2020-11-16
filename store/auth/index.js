import { auth } from '@/plugins/firebase';
import { fireDb } from '@/plugins/firebase';
import { persistance } from '@/plugins/firebase';
import Cookie from 'js-cookie';
import axios from 'axios';
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
  async login({ commit, dispatch }, account) {
    try {   
      let id = [];
      let data = '';
      let user = '';
      let userId = '';
      
      // console.log(id)
      // const loginRef = fireDb.collection('users').doc(id);
      // await loginRef.get().then((doc) => {
      //   if (doc.exists) {
      //     user = doc.data()
      //     userId = doc.id
      //   } else {
      //     console.log("No such document exists")
      //   }
      // }).catch((error) => {
      //   commit("setError", error)
      // })
      // const dataObj = {
      //   id: userId,
      //   data: user
      // }
      
      //commit('setUser', data)
      await auth.setPersistence(persistance.Persistence.LOCAL).then(() => {
        return auth.signInWithEmailAndPassword(account.email, account.password).then(() => {
          
          //dispatch('hideAuthModal')
          commit('setError', null)
        }).catch((error) => {
          commit('setError', error.message)
        })
      }).catch((error) => {
        console.log(error)
      })
      const userUID = auth.currentUser.uid;
      
      const userRef = fireDb.collection('users').doc(userUID)
      console.log(userRef)
      await userRef.get().then((doc) => {
        if (doc.exists) {
          data = doc.data()
        }
      }).catch((error) => {
        fireDb.collection('users').doc(userUID).set({
          email: account.email,
          id: userUID,
          role: "user",
          favorites: []
        }).then(() => {
          console.log("success")
        }).catch(() => {
          console.log("no user added")
        })
      })
      // const token = await auth.currentUser.getIdToken();
      // Cookie.set('user_token', token, {
      //   expires: 1
      // });
      commit('setUser', {
        email: account.email,
        id: userUID,
        role: data.role
      })
      // await auth.signInWithEmailAndPassword(account.email, account.password).then(() => {
      //   commit('setUser', dataObj)
      //   dispatch('hideAuthModal')
      //   commit('setError', null)
      // }).catch((error) => {
      //   commit('setError', error.message)
      // })
    } catch (error) {
      throw error;
    }
    
  },
  async passwordReset({ commit, dispatch }, email) {
    await auth.sendPasswordResetEmail(email).then(() => {
      //dispatch('hideAuthModal')
      commit('setError', null)
    }).catch((error) => {
      commit('setError', error.message)
    })
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
  },
  getUser: (state) => {
    return state.user
  }
}