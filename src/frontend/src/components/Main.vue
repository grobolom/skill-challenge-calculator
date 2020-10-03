<template>
  <div class="main">
    <div class="container px-4 sm:px-8 lg:px-16 xl:px-20 mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8">
        <div class="hero-text col-span-12">
          <h1 class="font-bold text-4xl md:text-5xl text-gray-900 leading-tight mb-8">
            Skill Challenge Calculator
          </h1>

          <div class='blocks md:flex px-8 md:items-start'>
            <div class='block md:w-1/2 md:inline-block text-base text-left text-xs md:text-sm'>
              <p class="text-gray-800 leading-relaxed">
                Select your party's skill bonus, skill check DC, successes, and failures for the
                <a class='text-blue-400' href='https://www.youtube.com/watch?v=GvOeqDpkBm8'>Skill Challenge</a>
                you'd like to run. The calculator will tell you how likely your party
                is to succeed the challenge.
              </p>
              <p class="text-gray-800 leading-relaxed mt-8">
                You should approximate your party's average skill bonus - it may be hard to
                calculate the exact chance of success for parties with
                extremely varied skill bonuses.
              </p>
            </div>

            <div class='block inputs md:w-1/2 md:inline-block'>
              <NumberInput
                name="Skill Bonus"
                v-bind:value="skillBonus"
                @on-increment="incrementSkillBonus"
                @on-decrement="decrementSkillBonus"
                @on-change="setSkillBonus" />
              <NumberInput
                name="Check DC"
                v-bind:value="checkDC"
                @on-increment="incrementCheckDC"
                @on-decrement="decrementCheckDC"
                @on-change="setCheckDC" />
              <NumberInput
                name="Failures"
                v-bind:value="failures"
                @on-increment="incrementFailures"
                @on-decrement="decrementFailures"
                @on-change="setFailures" />
            </div>
          </div>

          <div class='chart'>
            <Chart :skillBonus="skillBonus" :checkDC="checkDC" :failures="failures" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Chart from './Chart.vue'
import NumberInput from './NumberInput.vue'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  components: {
    Chart,
    NumberInput,
  },
  setup () {
    const store = useStore()

    const skillBonus = computed(() => store.state.skillBonus)
    const checkDC = computed(() => store.state.checkDC)
    const failures = computed(() => store.state.failures)

    function incrementSkillBonus() {
      store.dispatch('incrementSkillBonus')
    }

    function decrementSkillBonus() {
      store.dispatch('decrementSkillBonus')
    }

    function incrementCheckDC() {
      store.dispatch('incrementCheckDC')
    }

    function decrementCheckDC() {
      store.dispatch('decrementCheckDC')
    }

    function incrementFailures() {
      store.dispatch('incrementFailures')
    }

    function decrementFailures() {
      store.dispatch('decrementFailures')
    }

    function setSkillBonus (value: number) {
      store.dispatch('setSkillBonus', value)
    }

    function setCheckDC (value: number) {
      store.dispatch('setCheckDC', value)
    }

    function setFailures (value: number) {
      store.dispatch('setFailures', value)
    }

    return {
      skillBonus,
      incrementSkillBonus,
      decrementSkillBonus,
      setSkillBonus,

      checkDC,
      incrementCheckDC,
      decrementCheckDC,
      setCheckDC,

      failures,
      incrementFailures,
      decrementFailures,
      setFailures,
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="postcss">
.block {
  @apply border-t pt-8 mt-8 border-yellow-500
}

.chart {
  @apply border-t mt-8 border-yellow-500 mx-8
}

.inputs {
  @apply px-8
}
</style>
