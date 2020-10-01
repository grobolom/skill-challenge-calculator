<template>
  <div class="main">
    <div class="container px-4 sm:px-8 lg:px-16 xl:px-20 mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8">
        <div class="hero-text col-start-1 col-span-5 text-left">
          <h1 class="font-bold text-4xl md:text-5xl max-w-xl text-gray-900 leading-tight">
            Skill Challenge Calculator
          </h1>
          <hr class="w-12 h-1 bg-orange-500 rounded-full mt-8">
          <p class="text-gray-800 text-base leading-relaxed mt-8 font-medium">
            Select your party's skill bonus, skill check DC, successes, and failures for the
            <a class='text-blue-400' href='https://www.youtube.com/watch?v=GvOeqDpkBm8'>Skill Challenge</a>
            you'd like to run. The calculator will tell you how likely your party
            is to succeed the challenge.
          </p>
          <p class="text-gray-800 text-base leading-relaxed mt-8 font-medium">
            You should approximate your party's average skill bonus - it may be hard to
            calculate the exact chance of success for parties with
            extremely varied skill bonuses.
          </p>

          <hr class="w-12 h-1 bg-orange-500 rounded-full mt-8 mb-8">

          <NumberInput name="Skill Bonus" v-bind:startValue="skillBonus" v-bind:increment="incrementSkillBonus" />
          <NumberInput name="Check DC" v-bind:startValue="checkDC" />
          <NumberInput name="Failures" v-bind:startValue="failures" />
        </div>
        <div class="hero-image col-start-7 col-span-6 mt-16">
          <form @change="formSubmit">
            <Chart v-bind:info="info" />
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import { defineComponent } from 'vue'
import Chart from './Chart.vue'
import NumberInput from './NumberInput.vue'
import createStore from '../store/index'

export default defineComponent({
  components: {
    Chart,
    NumberInput,
  },
  setup () {
    const store = createStore

    const skillBonus = store.state.skillBonus
    const checkDC = store.state.checkDC
    const failures = store.state.failures

    return {
      skillBonus,
      checkDC,
      failures,
      incrementSkillBonus: store.mutations.incrementSkillBonus
      info: [
        {
          successes: 3,
          probability: .500,
        },
        {
          successes: 4,
          probability: .343,
        },
        {
          successes: 5,
          probability: .226
        },
        {
          successes: 6,
          probability: .114
        },
        {
          successes: 7,
          probability: .089
        },
        {
          successes: 8,
          probability: .054
        },
        {
          successes: 9,
          probability: .032
        },
        {
          successes: 10,
          probability: .019
        },
        {
          successes: 11,
          probability: .011
        },
      ]
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
