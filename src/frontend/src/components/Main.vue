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
            v-model="avg_skill_bonus"
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
            v-model="skill_check_dc"
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

<script>
export default {
  name: "Main",
  props: {
    msg: String
  },
  data() {
    return {
      avg_skill_bonus: 3,
      skill_check_dc: 14,
      successes: 6,
      failures: 3
    };
  },
  methods: {
    formSubmit(e) {
      e.preventDefault();
      const currentObj = this;

      this.axios
        .get("http://localhost:5051", {
          params: {
            skillBonus: this.avg_skill_bonus,
            dc: this.skill_check_dc,
            successes: this.successes,
            failures: this.failures
          }
        })
        .then(function(response) {
          currentObj.msg =
            +(response.data["ChanceOfSuccess"] * 100).toFixed(1) + "%";
        })
        .catch(function(err) {
          currentObj.msg = err;
        });

      return false;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
