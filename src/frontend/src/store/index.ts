import { createStore } from 'vuex'

export default createStore({
  state: {
    skillBonus: 4,
    checkDC: 15,
    failures: 3,
  },
  mutations: {
    incrementSkillBonus (state) {
      state.skillBonus++
    },
    decrementSkillBonus (state) {
      state.skillBonus = (state.skillBonus == 0) ? 0 : state.skillBonus - 1
    },
    incrementCheckDC (state) {
      state.checkDC++
    },
    decrementCheckDC (state) {
      state.checkDC = (state.checkDC == 0) ? 0 : state.checkDC - 1
    },
    incrementFailures (state) {
      state.failures++
    },
    decrementFailures (state) {
      state.failures = (state.failures == 0) ? 0 : state.failures - 1
    },
  },
  actions: {
    incrementSkillBonus({ commit }) {
      commit('incrementSkillBonus')
    },
    decrementSkillBonus({ commit }) {
      commit('decrementSkillBonus')
    },
    incrementCheckDC({ commit }) {
      commit('incrementCheckDC')
    },
    decrementCheckDC({ commit }) {
      commit('decrementCheckDC')
    },
    incrementFailures({ commit }) {
      commit('incrementFailures')
    },
    decrementFailures({ commit }) {
      commit('decrementFailures')
    }
  },
  modules: {
  }
})