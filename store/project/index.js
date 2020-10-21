import {
  auth
} from '@/plugins/firebase';
import {
  fireDb
} from '@/plugins/firebase';
import axios from 'axios';
import VueAxios from 'vue-axios';

export const state = () => ({
  projects: null,
  error: ''
});
export const mutations = {
  setProjects: (state, payload) => {
    state.projects = payload
  },
  setError(state, error) {
    state.error = error
  }
};
export const actions = {
  async fetchProjects({ commit, getters }) {
    const list = []
    try {
      var userData = getters['user'] ? getters['user'] : null;
      if (userData.role == 'admin') {
        await fireDb.collection('projects').get().then((qs) => {
          qs.forEach((doc) => {
            list.push(doc.data())           
          })
          commit('setProjects', list)
        })
        
      }
      if (userData.role == 'user') {
        await fireDb.collection('projects').where("client", "==", userData.id).get().then((qs) => {
          qs.forEach((doc) => {
            list.push(doc.data())
          })
          commit('setProjects', list)
        })
      }
    } catch (e) {
      const error = 'Please create a project'
      commit('setError', error)
    }
  },
  setMessage({ commit }) {
    axios.get('/test').then(res => res.data).then(items => {
      console.log(items);
      commit('setError', items)
    })
  }
}
export const getters = {
  user(state, getters, rootState, rootGetters) {
    return rootGetters['auth/getUser']
  },
}