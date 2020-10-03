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
import getSuccesses from '@/math/successes'

export default {
  props: ['skillBonus', 'checkDC', 'failures'],
  setup(props) {
    const info = computed(() => getSuccesses(props.skillBonus, props.checkDC, props.failures))

    return {
      info
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