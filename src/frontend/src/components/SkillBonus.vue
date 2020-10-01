<template>
  <div>
    <NumberInput
      name="Skill Bonus"
      v-bind:value="skillBonus"
      />
  </div>
</template>

<script type='ts'>
  import NumberInput from './NumberInput.vue'
  import store from '../store/index'

  export default {
    props: {
      name: String,
      value: Number,
      increment: Function,
      decrement: Function,
    },
    setup() {
      console.log(this.store)
      const store = createStore

      console.log(createStore.mutations)
    }
  }
</script>

<style scoped lang="postcss">
@tailwind base;
@tailwind components;
@tailwind utilities;

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
  @apply inline-block w-1/2 py-2 text-sm uppercase font-semibold text-left float-none;
}

input[type=number]::-webkit-outer-spin-button,
input[type=number]::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

input[type=number] {
    -moz-appearance:textfield;
}
</style>