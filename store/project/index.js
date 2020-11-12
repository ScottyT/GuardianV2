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
            const data = {
              client: doc.data().client,
              color: doc.data().color,
              description: doc.data().description,
              end: doc.data().end,
              name: doc.data().name,
              start: doc.data().start,
              type: doc.data().type,
              id: doc.id
            }
            list.push(data)
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
    await fireDb.collection("projects").doc(id).get()
      .then((doc) => {
        if (doc.exists) {
          commit('setProject', doc.data())
        } else {
          commit('setError', 'No such document exists!')
        }
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
  async favoritesAdded({ commit, getters }, event) {
    console.log(event)
    const favList = []
    let userData = getters['user'] ? getters['user'] : null;
    var userid = userData.id
    console.log(userData)
    favList.push(event)
    try {
      await fireDb.collection("users").doc(userid).set({
        favorites: favList
      }, { merge: true })
      .then(() => {          
        commit('setFavs', favList)
      })
      .catch((error) => {
        console.log(error)
      })
    } catch (e) {
      const error = 'Something wrong happened'
      commit('setError', error)
    }
  },
  
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