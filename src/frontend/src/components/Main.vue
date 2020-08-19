<template>
  <div class="main">
    <form>
      <label for='avg_skill_bonus'>Average Skill Bonus</label>
      <input name='avg_skill_bonus' type='text' />
      <label for='skill_check_dc'>Skill Check DC</label>
      <input name='skill_check_dc' type='text' />
      <label for='successes'>Successes</label>
      <input name='successes' type='text' />
      <label for='failures'>Failures</label>
      <input name='failures' type='text' />
      <label for='chance_of_success'>Chance of Success</label>
      <input name='chance_of_success' type='text' readonly :value="msg" />
    </form>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
