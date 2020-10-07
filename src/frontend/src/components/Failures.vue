<template>
  <NumberInput
    name='Failures'
    id='failures'
    :value='failures'
    @on-increment='incrementFailures'
    @on-decrement='decrementFailures'
    @on-change='setFailures' />
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
      const failures = computed(() => store.state.failures)

      function incrementFailures() {
        store.dispatch('incrementFailures')
      }

      function decrementFailures() {
        store.dispatch('decrementFailures')
      }

      function setFailures (value: number) {
        store.dispatch('setFailures', value)
      }

      return {
        failures,
        incrementFailures,
        decrementFailures,
        setFailures
      }
    }
  }
</script>