import {
  auth
} from '@/plugins/firebase';
import {
  fireDb
} from '@/plugins/firebase';
import axios from 'axios';
import VueAxios from 'vue-axios';

export const state = () => ({
  projects: [],
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
    let userData = getters['user'] ? getters['user'] : null;
    try {  
      if (userData.role == 'admin') {
        await fireDb.collection('projects').get().then((qs) => {
          qs.forEach((doc) => {
            list.push(doc.data())       
          })
          commit('setProjects', list)
        })
        
      } else if (userData.role == 'user') {
        await fireDb.collection('projects').where("client", "==", userData.email).get().then((qs) => {
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
  async updateProject({ commit, getters }, project) {
    let userData = getters['user'] ? getters['user'] : null;
    console.log("project to change:", project)
   
    
    var projectRef = fireDb.collection('projects').doc(project.id)
    await projectRef.update(project)
      .then(() => {
        console.log("project updated successfully")
      })
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