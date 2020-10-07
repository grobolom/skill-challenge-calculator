<template>
  <div class='wrapper'>
    <label class='label' :for='id'>{{ name }}</label>
    <button class='button left-button' tabindex="-1" v-on:click="$emit('on-decrement')">-</button>
    <input class='number' :id='id' type='number' :value="value" v-on:change="$emit('on-change', $event.target.value)" />
    <button class='button right-button' tabindex="-1" v-on:click="$emit('on-increment')">+</button>
  </div>
</template>

<script type='ts'>
  export default {
    props: {
      name: String,
      value: Number,
      id: String,
    },
  }
</script>

<style scoped lang="postcss">
.wrapper {
  @apply items-center mt-1
}

.button {
  @apply inline-block w-1/6 py-2 text-sm uppercase tracking-wider font-semibold;
  @apply bg-gray-400;
}

.number {
  @apply inline-block w-1/6 py-2 text-sm uppercase font-semibold text-center;
}

.number-input:focus {
  @apply outline-none shadow-outline;
}

.label {
  @apply inline-block w-1/2 py-2 text-sm font-semibold text-left float-none;
}
</style>