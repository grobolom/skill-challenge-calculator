<template>
  <div>
    <svg id='viz' width='400' height='400' class='container-border' />
  </div>
</template>

<script>
import * as d3 from 'd3'

export default {
  props: ['info'],
  data () {
    return {}
  },
  updated () {
    d3.selectAll("#viz > *").remove();

    const svg = d3.select('#viz')
                  .attr('font-family', 'sans-serif')
                  .attr('font-size', 10);

    const margin = { top: 10, right: 10, bottom: 70, left: 80 }
    const width = +svg.attr('width') - margin.left - margin.right
    const height = +svg.attr('height') - margin.top - margin.bottom

    // we want a central piece to draw on that's translated properly
    const g = svg.append("g")
                 .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

    const x = d3.scalePoint()
                .domain([3, 4, 5, 6, 7, 8, 9, 10, 11])
                .range([0, width])

    const y = d3.scaleLinear()
                .domain([100, 0])
                .range([0, height])

    const line = d3.line()
                   .x(d => x(d.successes))
                   .y(d => y(d.probability * 100))

    const xAxisCall = d3.axisBottom(x)
    const xAxis = g.append("g")
                   .attr("class", "x-axis")
                   .attr("transform", "translate(" + 0 + "," + (height + 20) + ")")
                   .call(xAxisCall);

    const yAxisCall = d3.axisLeft(y)
    const yAxis = g.append("g")
                   .attr("class", "y-axis")
                   .attr("transform", "translate(" + -30 + "," + 0 + ")")
                   .call(yAxisCall)

    xAxis.append("text")
         .attr("transform", "translate(" + (width - 23) + ", 0)")
         .attr("y", -6)
         .text("Successes")
         .attr("fill", "currentColor")

    yAxis.append("text")
         .attr("transform", "rotate(-90)")
         .attr("y", 16)
         .text("Overall Success Chance")
         .attr("fill", "currentColor")

    g.append('path')
     .attr('d', line(this.info))
     .attr('fill', 'none')
     .attr('stroke', 'rgb(0 0 0)')
     .attr('stroke-width', 1)

    // these are blank spots on the chart to provide a background for
    // the text below to be added and readable
    g.append("g")
     .attr("fill", "#fff")
     .selectAll("circle")
     .data(this.info)
     .join("circle")
     .attr("cx", d => x(d.successes))
     .attr("cy", d => y(d.probability * 100))
     .attr("r", 12)

    // this is the percentages as text on the chart
    g.append('g')
     .attr('text-anchor', 'middle')
     .selectAll('text')
     .data(this.info)
     .join('text')
     .attr('x', d => x(d.successes))
     .attr('y', d => y(d.probability * 100))
     .attr('dy', '0.35em')
     .text(d => (d.probability * 100).toFixed(1))
  }
}
</script>
<style scoped>
</style>