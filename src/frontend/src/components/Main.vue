<template>
  <div class="main">
    <div class="container mx-auto max-w-xs">
      <form @submit="formSubmit">
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

        <div class="mb-4">
          <button class="btn-submit">Submit</button>
        </div>

        <div class="mb-4">
          <label class="label" for="chance_of_success">Chance of Success</label >
          <input class="number-input" id="chance_of_success" name="chance_of_success" type="text" readonly :value="msg" />
        </div>
      </form>
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
