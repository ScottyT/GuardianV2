import { auth } from '@/plugins/firebase';
import { persistance } from '@/plugins/firebase';
export const state = () => ({
  user: null,
  error: null
});

export const mutations = {
  setUser: (state, payload) => {
    state.user = payload
  },
  setError: (state, payload) => {
    state.error = payload
  }
}

export const actions = {
  async login({ commit }, account) {
    await auth.setPersistence(persistance.Persistence.LOCAL).then(() => {
      return auth.signInWithEmailAndPassword(account.email, account.password);
    })
    .catch((error) => {
      commit('setError', error.message)
    })

    
  }
}