import { auth } from '@/plugins/firebase';
import { fireDb } from '@/plugins/firebase';
import { persistance } from '@/plugins/firebase';
import Cookie from 'js-cookie';
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
    let id = '';
    let data = '';
    const userRef = fireDb.collection('users').get().then((querySnapshot) => {
      querySnapshot.forEach((doc) => {
        id = doc.id;
        data = doc.data()
      })
    })
    const loginRef = userRef.doc(id);
    const token = await auth.currentUser.getIdToken();
    Cookie.set('user_token', token, {
      expires: 1
    });
    commit('setUser', userRef)
  }
}