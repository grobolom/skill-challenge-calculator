<template>
  <div class="main">
    <div class="container px-4 sm:px-8 lg:px-16 xl:px-20 mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8">
        <div class="hero-text col-span-12">
          <h1 class="font-bold text-4xl md:text-5xl text-gray-900 leading-tight mb-8">
            Skill Challenge Calculator
          </h1>

          <div class='blocks md:flex px-8 md:items-start'>
            <div class='block md:w-1/2 md:inline-block text-base text-left md:text-xs'>
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
              <NumberInput name="Skill Bonus" value="5" />
              <NumberInput name="Check DC" value="16" />
              <NumberInput name="Failures" value="3" />
            </div>
          </div>

          <div class='chart'>
            <Chart v-bind:info="info" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Chart from './Chart.vue'
import NumberInput from './NumberInput.vue'

export default defineComponent({
  components: {
    Chart,
    NumberInput,
  },
  setup () {
    const skillBonus = 4
    const checkDC = 15
    const failures = 3
    const info = [
      {
        successes: 3,
        probability: 0.500,
      },
      {
        successes: 4,
        probability: 0.343,
      },
      {
        successes: 5,
        probability: 0.226,
      },
      {
        successes: 6,
        probability: 0.114,
      },
      {
        successes: 7,
        probability: 0.089
      },
      {
        successes: 8,
        probability: 0.054,
      },
      {
        successes: 9,
        probability: 0.032,
      },
      {
        successes: 10,
        probability: 0.019,
      },
      {
        successes: 11,
        probability: 0.011,
      },
    ]

    return {
      skillBonus,
      checkDC,
      failures,
      info,
    }
  }
})
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
