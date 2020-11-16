import {
  auth
} from '@/plugins/firebase';
import {
  fireDb
} from '@/plugins/firebase';
import { fieldValue } from '@/plugins/firebase';
import axios from 'axios';
import VueAxios from 'vue-axios';

export const state = () => ({
  projects: null,
  error: '',
  favorites: [],
  heartBtn: {
    liked: false
  },
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
  },
  updateLikeBtn: (state, payload) => {
    state.heartBtn.liked = payload.liked
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
          list.sort(function (a, b) {
            if (a.name > b.name) {
              return -1;
            }
            if (b.name > a.name) {
              return 1;
            }
            return 0;
          });
          console.log(list)
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
  async favoritesAdded({ commit, getters }, project) {
    console.log(project)
    const favList = []
    let userData = getters['user'] ? getters['user'] : null;
    var userid = userData.id
    console.log(userData)
    favList.push(project)
    try {
      await fireDb.collection("users").doc(userid).update({
        favorites: fieldValue.arrayUnion(project)
      })
      .then(() => {
        favList.sort(function (a, b) {
          if (a.name > b.name) {
            return -1;
          }
          if (b.name > a.name) {
            return 1;
          }
          return 0;
        })  
        commit('setFavs', favList)
      })
      .catch((error) => {
        console.log(error)
      })
    } catch (e) {
      const error = 'Something wrong happened'
      commit('setError', e)
    }
  },
  async removeFav({ commit, getters, state }, project) {
    let userData = getters['user'] ? getters['user'] : null;
    var userid = userData.id
    var favRef = fireDb.collection("users").doc(userid)
    
    let updatedFavList = state.favorites.filter(el => el.id !== project.id)
    console.log(updatedFavList)
    try {    
      await favRef.update({
        favorites: updatedFavList
      }).then(() => {
        
        commit("setFavs", updatedFavList)
      })
    } catch (e) {
      commit('setError', e)
    }
  },
  showLikedHeart({ commit }, btnState) {
    
    commit("updateLikedBtn", { liked: btnState });
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