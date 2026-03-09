<template>
  <div class="min-h-screen bg-background">
    <nav class="border-b">
      <div class="mx-auto flex h-14 max-w-7xl items-center gap-6 px-6">
        <NuxtLink to="/" class="text-lg font-bold">XStorage</NuxtLink>
  
      </div>
    </nav>
    <main class="mx-auto max-w-7xl px-6 py-8">
      <slot />
    </main>
  </div>
</template>
