<template>
  <div class='footer'>
    <p>
      made by <a href='https://twitter.com/grobolom'>@grobolom</a>
    </p>
    <span class='email'>
      <a href='mailto:vasja.volin@gmail.com'>
        <svg class='email-icon' xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
        </svg>
      </a>
    </span>
  </div>
</template>

<style scoped lang="postcss">
.footer {
  @apply flex
}

.email {
  @apply ml-1
}

.email-icon {
  width: 18px;
  height: 18px;
}
</style>