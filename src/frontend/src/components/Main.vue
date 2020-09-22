<template>
  <div class="main">
    <div class="container px-4 sm:px-8 lg:px-16 xl:px-20 mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8 items-center">
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
        </div>
        <div class="hero-image col-start-7 col-span-6">
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
import { d3 } from 'd3'

const data = [
  {
    name: "Prostate",
    period: "5 Year",
    value: 99
  }
]

const height = 1000
const width = 800
const margin = ({top: 30, right: Math.max(width - 600, 20), bottom: 20, left: 120})
const formatValue = d3.format(",")

const x = d3.scalePoint()
    .domain(data.map((d: any) => d.period))
    .range([margin.left, width - margin.right])

const xAxis = (g: any) => g
    .attr("transform", `translate(0,${margin.top})`)
    .call(d3.axisTop(x))
    .call((g: any) => g.select(".domain").remove())

const y = d3.scaleLinear()
    .domain(d3.extent(data, (d: any) => d.value))
    .range([height - margin.bottom, margin.top])

const line = d3.line()
    .x((d: any) => x(d.period))
    .y((d: any) => y(d.value))

var chart = () => {
  const svg = d3.create("svg")
      .attr("viewBox", [0, 0, width, height])
      .attr("font-family", "sans-serif")
      .attr("font-size", 10);

  svg.append("g")
      .call(xAxis);

  svg.append("g")
      .attr("fill", "none")
      .attr("stroke", "#000")
      .selectAll("path")
      .data(d3.groups(data, (d: any) => d.name))
      .join("path")
      .attr("d", ([, group]: [any, any]) => line(group))
      .call((path: any) => path.clone(true))
      .attr("stroke", "#fff")
      .attr("stroke-width", 5);

  svg.append("g")
      .attr("fill", "#fff")
      .selectAll("circle")
      .data(data)
      .join("circle")
      .attr("cx", (d: any) => x(d.period))
      .attr("cy", (d: any) => y(d.value))
      .attr("r", 10);

  svg.append("g")
      .attr("text-anchor", "middle")
      .selectAll("text")
      .data(data)
      .join("text")
      .attr("x", (d: any) => x(d.period))
      .attr("y", (d: any) => y(d.value))
      .attr("dy", "0.35em")
      .text((d: any) => formatValue(d.value));

  svg.append("g")
      .attr("text-anchor", "end")
      .selectAll("text")
      .data(d3.groups(data, (d: any) => d.name))
      .join("text")
      .attr("x", margin.left - 12)
      .attr("y", ([key, [d]]: [any, [any]]) => y(d.value))
      .attr("dy", "0.35em")
      .text(([key]: [any]) => key);

  return svg.node();
}


export default class Main extends Vue {
  avgSkillBonus = 3
  skillCheckDc = 14
  successes = 3
  failures = 3
  msg = ""

  formSubmit(e: any) {
    e.preventDefault();

    axios
      .get(process.env.VUE_APP_API_URL, {
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
