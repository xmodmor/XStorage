<template>
  <div class="min-h-screen bg-background">
    <nav class="border-b">
      <div class="mx-auto flex h-14 max-w-7xl items-center gap-6 px-6">
        <NuxtLink to="/" class="text-lg font-bold">XStorage</NuxtLink>
        <NuxtLink to="/dashboard" class="text-sm text-muted-foreground hover:text-foreground">Dashboard</NuxtLink>
        <NuxtLink to="/apps" class="text-sm text-muted-foreground hover:text-foreground">Apps</NuxtLink>
        <NuxtLink to="/buckets" class="text-sm text-muted-foreground hover:text-foreground">Buckets</NuxtLink>
        <NuxtLink to="/objects" class="text-sm text-muted-foreground hover:text-foreground">Objects</NuxtLink>
        <NuxtLink to="/api-keys" class="text-sm text-muted-foreground hover:text-foreground">API Keys</NuxtLink>
      </div>
    </nav>
    <main class="mx-auto max-w-7xl px-6 py-8">
      <slot />
    </main>
  </div>
</template>
