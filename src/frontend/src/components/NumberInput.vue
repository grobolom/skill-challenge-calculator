<template>
  <div class='wrapper'>
    <label class="label">{{ name }}</label>
    <button class='button left-button' v-on:click="decrement">-</button>
    <input class='number' type='number' :value="value" v-on:change="change" />
    <button class='button right-button' v-on:click="increment">+</button>
  </div>
</template>

<script type='ts'>
  export default {
    props: {
      name: String,
      value: Number,
    },
    methods: {
      increment() {
        this.$emit('on-increment')
      },
      decrement() {
        this.$emit('on-decrement')
      },
      change(event) {
        this.$emit('on-change', event.target.value)
      }
    }
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

/* remove the webkit up and down arrows for the number inputs */
input[type=number]::-webkit-outer-spin-button,
input[type=number]::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

input[type=number] {
    -moz-appearance:textfield;
}
</style>