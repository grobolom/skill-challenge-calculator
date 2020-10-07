<template>
  <div class='wrapper'>
    <table class='table'>
      <thead>
        <tr>
          <th>
            Successes
          </th>
          <th>
            Chance to Succeed
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in info" :key="item.successes" class='row'>
          <td class="row-left">
            {{ item.successes }}
          </td>
          <td>
            {{ (item.probability * 100).toFixed(1) }}%
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script type='ts'>
import { computed } from 'vue'
import { useStore } from 'vuex'
import getSuccesses from '@/math/successes'

export default {
  setup() {
    const store = useStore()

    const skillBonus = computed(() => store.state.skillBonus)
    const checkDC = computed(() => store.state.checkDC)
    const failures = computed(() => store.state.failures)
    const info = computed(() => getSuccesses(skillBonus.value, checkDC.value, failures.value))

    return {
      info,
      skillBonus,
      checkDC,
      failures
    }
  },
}
</script>

<style scoped lang='postcss'>
.table {
  @apply w-full px-4 mt-4 table-auto
}

.row {
  @apply py-4
}

.row:nth-child(2n + 1) {
  @apply bg-gray-200
}

.row-left {
  @apply border-r border-gray-600
}
</style>