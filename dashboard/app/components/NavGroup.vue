<script setup lang="ts">
import type { NavItem } from "@/types/nav"

const props = defineProps<{
  items: NavItem[]
  groupLabel?: string
  showDropdowns?: boolean
}>()

// Group items by their group property while preserving order
const groupedItems = computed(() => {
  const groups: Record<string, NavItem[]> = {}
  const groupOrder: string[] = []
  
  props.items.forEach(item => {
    const groupKey = item.group || 'default'
    if (!groups[groupKey]) {
      groups[groupKey] = []
      if (!groupOrder.includes(groupKey)) {
        groupOrder.push(groupKey)
      }
    }
    groups[groupKey].push(item)
  })
  
  return { groups, groupOrder }
})

// Get ordered groups (preserving original order)
const orderedGroups = computed(() => {
  return groupedItems.value.groupOrder.filter(key => 
    groupedItems.value.groups[key] && groupedItems.value.groups[key].length > 0
  )
})
</script>

<template>
  <div class="space-y-4">
    <!-- All groups in original order -->
    <SidebarGroup 
      v-for="groupKey in orderedGroups" 
      :key="groupKey"
    >
      <SidebarGroupLabel v-if="groupKey !== 'default'">{{ groupKey }}</SidebarGroupLabel>
      <SidebarGroupLabel v-else-if="groupLabel">{{ groupLabel }}</SidebarGroupLabel>
      <SidebarMenu>
        <NavItem 
          v-for="item in groupedItems.groups[groupKey]" 
          :key="item.title"
          :item="item"
          :show-dropdown="showDropdowns"
        />
      </SidebarMenu>
    </SidebarGroup>
  </div>
</template>
