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
      console.log('changed skill bonus to ' + state.skillBonus)
    },
    decrementSkillBonus (state) {
      state.skillBonus = (state.skillBonus == 0) ? 0 : state.skillBonus - 1
      console.log('changed skill bonus to ' + state.skillBonus)
    },
    incrementCheckDC (state) {
      state.checkDC++
      console.log('changed check DC to ' + state.checkDC)
    },
    decrementCheckDC (state) {
      state.checkDC = (state.checkDC == 0) ? 0 : state.checkDC - 1
      console.log('changed check DC to ' + state.checkDC)
    },
    incrementFailures (state) {
      state.failures++
      console.log('changed failures to ' + state.failures)
    },
    decrementFailures (state) {
      state.failures = (state.failures == 0) ? 0 : state.failures - 1
      console.log('changed failures to ' + state.failures)
    },
  },
  actions: {
  },
  modules: {
  }
})