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
  error: '',
  favorites: [],
  data: null
});
export const mutations = {
  setProjects: (state, payload) => {
    state.projects = payload
  },
  setError(state, error) {
    state.error = error
  },
  setFavs: (state, payload) => {
  state.favorites = payload
  },
  setProject: (state, payload) => {
    state.data = payload
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
  async fetchProject({ commit }, id) {
    var project = {}
    await fireDb.collection("projects").where("id", "==", id).get()
      .then((qs) => {
        qs.forEach((doc) => {
          project = doc.data()
        })
        commit('setProject', project)
      })
      .catch((e) => {
        commit('setError', e)
      })
    
    
  },
  setMessage({ commit }) {
    axios.get('/test').then(res => res.data).then(items => {
      console.log(items);
      commit('setError', items)
    })
  },
  async favoritesAdded({ commit }, event) {
    const favList = []
    favList.push(event)
    commit('setFavs', favList)
  }
}
export const getters = {
  user(state, getters, rootState, rootGetters) {
    return rootGetters['auth/getUser']
  },
  getProject: (state) => {
    return state.data
  },
  getProjects: (state) => {
    return state.projects
  }
}