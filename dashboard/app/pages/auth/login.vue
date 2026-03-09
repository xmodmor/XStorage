<script setup lang="ts">

definePageMeta({
  layout: 'empty'
})
</script>

<template>
  <div class="grid min-h-svh lg:grid-cols-2">
    <div class="flex flex-col gap-4 p-6 md:p-10">
      <div class="flex justify-center gap-2 md:justify-start">
        <NuxtLink to="/" class="flex items-center gap-2 font-medium group" aria-label="XStorage Home">
          <span class="text-xl font-bold tracking-tight transition group-hover:opacity-80">XStorage</span>
        </NuxtLink>
      </div>
      <div class="flex flex-1 items-center justify-center">
        <div class="w-full max-w-xs">
          <AuthLoginForm />
        </div>
      </div>
    </div>
    <div class="bg-muted relative hidden lg:block">
      <img
        src="/images/placeholder.png"
        alt="Image"
        class="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
      >
    </div>
  </div>
</template>
