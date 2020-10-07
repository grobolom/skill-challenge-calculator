<template>
  <NumberInput
    name='Skill Bonus'
    id='skill-bonus'
    :value='skillBonus'
    @on-increment='incrementSkillBonus'
    @on-decrement='decrementSkillBonus'
    @on-change='setSkillBonus' />
</template>

<script lang='ts'>
  import { computed } from 'vue'
  import { useStore } from 'vuex'
  import NumberInput from './NumberInput.vue'

  export default {
    components: {
      NumberInput,
    },
    setup() {
      const store = useStore()
      const skillBonus = computed(() => store.state.skillBonus)

      function incrementSkillBonus() {
        store.dispatch('incrementSkillBonus')
      }

      function decrementSkillBonus() {
        store.dispatch('decrementSkillBonus')
      }

      function setSkillBonus (value: number) {
        store.dispatch('setSkillBonus', value)
      }

      return {
        skillBonus,
        incrementSkillBonus,
        decrementSkillBonus,
        setSkillBonus
      }
    }
  }
</script>