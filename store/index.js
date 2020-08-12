export const state = () => ({
  header: {},
  footer: {},
  error: ''
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
  }
}
