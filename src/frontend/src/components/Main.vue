<template>
  <div class="main">
    <form @submit="formSubmit">
      <label for='avg_skill_bonus'>Average Skill Bonus</label>
      <input name='avg_skill_bonus' type='text' v-model="avg_skill_bonus" />

      <label for='skill_check_dc'>Skill Check DC</label>
      <input name='skill_check_dc' type='text' v-model="skill_check_dc" />

      <label for='successes'>Successes</label>
      <input name='successes' type='text' v-model="successes" />

      <label for='failures'>Failures</label>
      <input name='failures' type='text' v-model="failures" />

      <button>Submit</button>

      <label for='chance_of_success'>Chance of Success</label>
      <input name='chance_of_success' type='text' readonly :value="msg" />
    </form>
  </div>
</template>

<script>
export default {
  name: 'Main',
  props: {
    msg: String
  },
  data() {
    return {
      avg_skill_bonus: 5,
      skill_check_dc: 5,
      successes: 0,
      failures: 0,
    }
  },
  methods: {
    formSubmit(e) {
      e.preventDefault();
      let currentObj = this;

      this.axios.get("http://localhost:5051", {
        params: {
          skillBonus: this.avg_skill_bonus,
          dc: this.skill_check_dc,
          successes: this.successes,
          failures: this.failures
        }
      }).then(function(response) {
        currentObj.msg = response.data["ChanceOfSuccess"]
      }).catch(function(err) {
        currentObj.msg = err
      })

      return false;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
