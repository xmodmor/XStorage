<script setup lang="ts">
import { Database } from 'lucide-vue-next'

const props = withDefaults(defineProps<{
  mode?: 'full' | 'icon'
  size?: 'sm' | 'md' | 'lg'
  linkToHome?: boolean
}>(), {
  mode: 'full',
  size: 'md',
  linkToHome: true,
})

const iconSize = computed(() => {
  if (props.size === 'sm') return 18
  if (props.size === 'lg') return 28
  return 22
})

const textClass = computed(() => {
  if (props.size === 'sm') return 'text-sm'
  if (props.size === 'lg') return 'text-xl'
  return 'text-base'
})
</script>

<template>
  <component
    :is="linkToHome ? 'NuxtLink' : 'div'"
    :to="linkToHome ? '/' : undefined"
    class="flex items-center gap-1 select-none"
  >
    <div class="flex items-center justify-center rounded-md p-1.5 size-16">
      <img src="/images/logo.png" alt="XStorage Logo" class="w-full h-full object-contain" />
    </div>
    <span
      v-if="mode === 'full'"
      :class="[textClass, 'font-semibold tracking-tight text-primary']"
    >
      XStorage
    </span>
  </component>
</template>
