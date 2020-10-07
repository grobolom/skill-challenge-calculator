<template>
  <div>
    <NumberInput
      name='Check DC'
      id='check-dc'
      :value='checkDC'
      @on-increment='incrementCheckDC'
      @on-decrement='decrementCheckDC'
      @on-change='setCheckDC' />
  </div>
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
      const checkDC = computed(() => store.state.checkDC)

      function incrementCheckDC() {
        store.dispatch('incrementCheckDC')
      }

      function decrementCheckDC() {
        store.dispatch('decrementCheckDC')
      }

      function setCheckDC (value: number) {
        store.dispatch('setCheckDC', value)
      }

      return {
        checkDC,
        incrementCheckDC,
        decrementCheckDC,
        setCheckDC,
      }
    }
  }
</script>