<template>
  <div class="main">
    <div class="container px-4 sm:px-8 lg:px-16 xl:px-20 mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8 items-center">
        <div class="hero-text col-span-6 text-left">
          <h1 class="font-bold text-4xl md:text-5xl max-w-xl text-gray-900 leading-tight">
            Skill Challenge Calculator
          </h1>
          <hr class="w-12 h-1 bg-orange-500 rounded-full mt-8">
          <p class="text-gray-800 text-base leading-relaxed mt-8 font-semibold">
            Put in your party's skill bonus, skill check DC, successes, and failures for the
            Skill Challenge you'd like to run. The calculator will tell you how likely your party
            is to succeed the challenge.
          </p>
          <p class="text-gray-800 text-base leading-relaxed mt-8 font-semibold">
            You should approximate your party's average skill bonus - it may be hard to
            calculate the exact chance of success for parties with
            extremely varied skill bonuses.
          </p>
        </div>
        <div class="hero-image col-span-6">
          <form @change="formSubmit">
            <div class="mb-4">
              <label class="label" for="avg_skill_bonus" >Average Skill Bonus</label >
              <input class="number-input" id="avg_skill_bonus" name="avg_skill_bonus" type="number" v-model="avgSkillBonus" />
            </div>

            <div class="mb-4">
              <label class="label" for="skill_check_dc" >Skill Check DC</label >
              <input class="number-input" id="skill_check_dc" name="skill_check_dc" type="number" v-model="skillCheckDc" />
            </div>

            <div class="mb-4">
              <label class="label" for="successes" >Successes</label >
              <input class="number-input" id="successes" name="successes" type="number" v-model="successes" />
            </div>

            <div class="mb-4">
              <label class="label" for="failures" >Failures</label >
              <input class="number-input" id="failures" name="failures" type="number" v-model="failures" />
            </div>

            <hr class="w-12 h-1 bg-orange-500 mt-8 mb-8">

            <div class="mb-4">
              <label class="label" for="chance_of_success">Chance of Success</label >
              <input class="number-input" id="chance_of_success" name="chance_of_success" type="text" readonly :value="msg" />
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-class-component';
import axios from 'axios'

export default class Main extends Vue {
  avgSkillBonus = 3
  skillCheckDc = 14
  successes = 3
  failures = 3
  msg = ""

  formSubmit(e: any) {
    e.preventDefault();

    axios
      .get("http://localhost:5051", {
        params: {
          skillBonus: this.avgSkillBonus,
          dc: this.skillCheckDc,
          successes: this.successes,
          failures: this.failures
        }
      })
      .then((response: any) => {
        this.msg = +(response.data["ChanceOfSuccess"] * 100).toFixed(1) + "%";
      })
      .catch((err: any) => {
        this.msg = err;
      });

    return false;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
