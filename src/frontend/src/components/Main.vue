<template>
  <div class="main">
    <div class="container mx-auto max-w-xs">
      <form @submit="formSubmit">
        <div class="mb-4">
          <label
            class="text-gray-700 text-sm font-bold mb-2 float-left"
            for="avg_skill_bonus"
            >Average Skill Bonus</label
          >
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline-1"
            id="avg_skill_bonus"
            name="avg_skill_bonus"
            type="number"
            v-model="avgSkillBonus"
          />
        </div>

        <div class="mb-4">
          <label
            class="text-gray-700 text-sm font-bold mb-2 float-left"
            for="skill_check_dc"
            >Skill Check DC</label
          >
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="skill_check_dc"
            name="skill_check_dc"
            type="number"
            v-model="skillCheckDc"
          />
        </div>

        <div class="mb-4">
          <label
            class="text-gray-700 text-sm font-bold mb-2 float-left"
            for="successes"
            >Successes</label
          >
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="successes"
            name="successes"
            type="number"
            v-model="successes"
          />
        </div>

        <div class="mb-4">
          <label
            class="text-gray-700 text-sm font-bold mb-2 float-left"
            for="failures"
            >Failures</label
          >
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="failures"
            name="failures"
            type="number"
            v-model="failures"
          />
        </div>

        <div class="mb-4">
          <button
            class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow"
          >
            Submit
          </button>
        </div>

        <div class="mb-4">
          <label
            class="text-gray-700 text-sm font-bold mb-2 float-left"
            for="chance_of_success"
            >Chance of Success</label
          >
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="chance_of_success"
            name="chance_of_success"
            type="text"
            readonly
            :value="msg"
          />
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import axios from 'axios'

@Options({})

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
<style scoped></style>
