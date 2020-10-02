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
import { ref } from 'vue'
import axios from 'axios'
import { onUpdated } from 'vue'

async function getData(skillBonus, checkDC, failures) {
  const response = await axios.get(
    process.env.VUE_APP_API_URL + "/v2",
    {
      params: {
        skillBonus: skillBonus,
        failures: failures,
        dc: checkDC,
      }
    }
  )

  return response.data
}

export default {
  props: ['skillBonus', 'checkDC', 'failures'],
  async setup(props) {
    const info = ref(null)

    onUpdated(async () => {
      info.value = await getData(props.skillBonus, props.checkDC, props.failures)
    })

    info.value = await getData(props.skillBonus, props.checkDC, props.failures)

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