<template>
  <div>
    <h2>Chart</h2>
    <svg id='viz' width='400' height='400' class='container-border' />
  </div>
</template>

<script>
import * as d3 from 'd3'

export default {
  data () {
    return {}
  },
  mounted () {
    const data = [
      {
        successes: 4,
        percent: 50.5
      },
      {
        successes: 5,
        percent: 30.2
      },
      {
        successes: 6,
        percent: 25.1
      },
      {
        successes: 7,
        percent: 20.6
      }
    ]

    const width = 500
    const height = 500
    const svg = d3
      .select('#viz')
      .attr('viewBox', [-10, -10, width, height])
      .attr('font-family', 'sans-serif')
      .attr('font-size', 10);

    const x = d3.scalePoint()
                .domain(data.map(d => d.successes))
                .range([0, 400])

    const y = d3.scaleLinear()
                .domain([100, 0])
                .range([0, 400])

    const line = d3.line()
                   .x(d => x(d.successes))
                   .y(d => y(d.percent))

    svg
      .append('path')
      .attr('d', line(data))
      .attr('fill', 'none')
      .attr('stroke', 'rgb(0 0 0)')
      .attr('stroke-width', 1);

    // these are blank spots on the chart to provide a background for
    // the text below to be added and readable
    svg
      .append("g")
      .attr("fill", "#fff")
      .selectAll("circle")
      .data(data)
      .join("circle")
      .attr("cx", d => x(d.successes))
      .attr("cy", d => y(d.percent))
      .attr("r", 12);

    // this is the percentages as text on the chart
    svg
      .append('g')
      .attr('text-anchor', 'middle')
      .selectAll('text')
      .data(data)
      .join('text')
      .attr('x', d => x(d.successes))
      .attr('y', d => y(d.percent))
      .attr('dy', '0.35em')
      .text(d => d.percent);
  }
}
</script>
<style scoped>
</style>